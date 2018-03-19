pragma solidity ^0.4.18;

import "zeppelin-solidity/contracts/math/SafeMath.sol";
import "zeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "./NectarToken.sol";


contract BountyRegistry is Pausable {
    using SafeMath for uint256;

    struct Bounty {
        uint128 guid;
        address author;
        uint256 amount;
        string artifactURI;
        uint256 expirationBlock;
        bool resolved;
        uint256 verdicts;
    }

    struct Assertion {
        address author;
        uint256 bid;
        uint256 mask;
        uint256 verdicts;
        string metadata;
    }

    event NewBounty(
        uint128 guid,
        address author,
        uint256 amount,
        string artifactURI,
        uint256 expirationBlock
    );

    event NewAssertion(
        uint128 bountyGuid,
        address author,
        uint256 index,
        uint256 bid,
        uint256 mask,
        uint256 verdicts,
        string metdata
    );

    event NewVerdict(
        uint128 bountyGuid,
        uint256 verdicts
    );

    address internal owner;
    NectarToken internal token;

    // 0.0625NCT (1/16)
    uint256 public constant BOUNTY_FEE = 62500000000000000;
    uint256 public constant ASSERTION_FEE = 62500000000000000;
    uint256 public constant BOUNTY_AMOUNT_MINIMUM = 62500000000000000;
    uint256 public constant ASSERTION_BID_MINIMUM = 62500000000000000;

    uint128[] public bountyGuids;
    mapping (uint128 => Bounty) public bountiesByGuid;
    mapping (uint128 => Assertion[]) public assertionsByGuid;
    mapping (address => bool) public arbiters;

    function BountyRegistry(address nectarTokenAddr) public {
        owner = msg.sender;
        token = NectarToken(nectarTokenAddr);
    }

    modifier onlyArbiter() {
        require(arbiters[msg.sender]);
        _;
    }

    function addArbiter(address newArbiter) external whenNotPaused onlyOwner {
        require(newArbiter != address(0));
        require(!arbiters[newArbiter]);

        arbiters[newArbiter] = true;
    }

    function removeArbiter(address arbiter) external whenNotPaused onlyOwner {
        arbiters[arbiter] = false;
    }

    function postBounty(
        uint128 guid,
        uint256 amount,
        string artifactURI,
        uint256 durationBlocks
    )
        external
        whenNotPaused
    {
        // Check if a bounty with this GUID has already been initialized
        require(bountiesByGuid[guid].author == address(0));
        // Check that our bounty amount is sufficient
        require(amount >= BOUNTY_AMOUNT_MINIMUM);
        // Check that our URI is non-empty
        require(bytes(artifactURI).length > 0);
        // Check that our duration is non-zero
        require(durationBlocks > 0);

        // Assess fees and transfer bounty amount into escrow
        require(token.transferFrom(msg.sender, address(this), amount.add(BOUNTY_FEE)));

        Bounty memory b = Bounty(
            guid,
            msg.sender,
            amount,
            artifactURI,
            durationBlocks.add(block.number),
            false,
            0
        );
        bountiesByGuid[guid] = b;
        bountyGuids.push(guid);

        NewBounty(
            b.guid,
            b.author,
            b.amount,
            b.artifactURI,
            b.expirationBlock
        );
    }

    function postAssertion(
        uint128 bountyGuid,
        uint256 bid,
        uint256 mask,
        uint256 verdicts,
        string metadata
    )
        external
        whenNotPaused
    {
        // Check if this bounty has been initialized
        require(bountiesByGuid[bountyGuid].author != address(0));
        // Check that our bid amount is sufficient
        require(bid >= ASSERTION_BID_MINIMUM);
        // Check if this bounty is active
        require(bountiesByGuid[bountyGuid].expirationBlock > block.number);

        // Assess fees and transfer bid amount into escrow
        require(token.transferFrom(msg.sender, address(this), bid.add(ASSERTION_FEE)));

        Assertion memory a = Assertion(
            msg.sender,
            bid,
            mask,
            verdicts,
            metadata
        );

        uint256 index = assertionsByGuid[bountyGuid].push(a) - 1;

        NewAssertion(
            bountyGuid,
            a.author,
            index,
            a.bid,
            a.mask,
            a.verdicts,
            a.metadata
        );
    }

    function settleBounty(
        uint128 bountyGuid,
        uint256 verdicts
    )
        external
        onlyArbiter
        whenNotPaused
    {
        Bounty memory bounty = bountiesByGuid[bountyGuid];
        Assertion[] memory assertions = assertionsByGuid[bountyGuid];

        // Check if this bounty has been initialized
        require(bounty.author != address(0));
        // Check if the deadline has expired
        require(bounty.expirationBlock <= block.number);

        bounty.verdicts = verdicts;
        bounty.resolved = true;

        uint256 i = 0;

        uint256 numAssertions = assertions.length;
        uint256 pot = bounty.amount;
        uint256 fees = BOUNTY_FEE.add(ASSERTION_FEE.mul(numAssertions));

        uint256 numLosers = 0;
        for (i = 0; i < numAssertions; i++) {
            // TODO: For now, verdicts are all-or-nothing
            if (assertions[i].verdicts != verdicts) {
                pot = pot.add(assertions[i].bid);
                numLosers = numLosers.add(1);
            }
        }

        // Arbiter will get a split too
        uint256 numWinners = numAssertions.sub(numLosers).add(1);
        // Split is bounty amount + all bids divided by number of winners,
        // rounded down. Remainder goes to arbiter.
        uint256 split = pot.div(numWinners);
        uint256 remainder = pot % numWinners;

        uint256 reward = 0;
        for (i = 0; i < numAssertions; i++) {
            if (assertions[i].verdicts == verdicts) {
                reward = assertions[i].bid.add(split);
                // TODO: Don't revert if one transfer fails, what to do?
                // Transfers are not expected to ever fail though
                require(token.transfer(assertions[i].author, reward));
            }
        }

        // Transfer remainder of pot to arbiter, handles fractional NCT remainders
        require(token.transfer(msg.sender, split.add(fees).add(remainder)));

        NewVerdict(bountyGuid, verdicts);
    }

    function getNumberOfBounties() external view returns (uint) {
        return bountyGuids.length;
    }

    function getNumberOfAssertions(uint128 bountyGuid) external view returns (uint) {
        // Check if this bounty has been initialized
        require(bountiesByGuid[bountyGuid].author != address(0));

        return assertionsByGuid[bountyGuid].length;
    }
}
