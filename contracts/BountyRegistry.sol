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
        uint256 verdicts;
        uint256 bid;
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
        uint256 verdicts,
        uint256 index,
        uint256 bid,
        string metdata
    );

    event NewVerdict(
        uint128 bountyGuid,
        uint256 verdicts
    );

    address internal owner;
    NectarToken internal token;

    uint256 public constant BOUNTY_FEE = 10;
    uint256 public constant ASSERTION_FEE = 1;
    uint256 public constant BOUNTY_AMOUNT_MINIMUM = 1;
    uint256 public constant ASSERTION_BID_MINIMUM = 1;

    uint128[] public bountyGuids;
    mapping (uint128 => Bounty) public bountiesByGuid;
    mapping (uint128 => Assertion[]) public assertionsByGuid;
    mapping (address => bool) public arbiters;

    function BountyRegistry(address nectarTokenAddr) public {
        owner = msg.sender;
        token = NectarToken(nectarTokenAddr);
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
        uint256 verdicts,
        uint256 bid,
        string metadata
    )
        external
        whenNotPaused
    {
        // Check if this bounty has been initialized
        require(bountiesByGuid[bountyGuid].author != address(0));
        // Check if this bounty is active
        require(bountiesByGuid[bountyGuid].expirationBlock > block.number);
        // Check that our bid amount is sufficient
        require(bid >= ASSERTION_BID_MINIMUM);

        // Assess fees and transfer bid amount into escrow
        require(token.transferFrom(msg.sender, address(this), bid.add(ASSERTION_FEE)));

        Assertion memory a = Assertion(
            msg.sender,
            verdicts,
            bid,
            metadata
        );

        uint256 index = assertionsByGuid[bountyGuid].push(a) - 1;

        NewAssertion(
            bountyGuid,
            a.author,
            a.verdicts,
            index,
            a.bid,
            a.metadata
        );
    }

    function settleBounty(
        uint128 bountyGuid,
        uint256 verdicts
    )
        external
        whenNotPaused
    {
        // Check if this bounty has been initialized
        require(bountiesByGuid[bountyGuid].author != address(0));
        // Check if the deadline has expired
        require(bountiesByGuid[bountyGuid].expirationBlock <= block.number);
        // Check if we are an arbiter allowed to settle
        require(arbiters[msg.sender]);

        bountiesByGuid[bountyGuid].verdicts = verdicts;
        bountiesByGuid[bountyGuid].resolved = true;

        uint256 pot = bountiesByGuid[bountyGuid].amount;
        uint256 numAssertions = assertionsByGuid[bountyGuid].length;
        uint256 numLosers = 0;
        for (uint256 i = 0; i < numAssertions; i++) {
            Assertion memory a = assertionsByGuid[bountyGuid][i];

            // For now, verdicts are all-or-nothing
            if (a.verdicts != verdicts) {
                pot = pot.add(a.bid);
                numLosers = numLosers.add(1);
            }
        }

        // Arbiter will get a split too
        uint256 numWinners = numAssertions.sub(numLosers).add(1);
        uint256 split = pot.div(numWinners);

        for (i = 0; i < numAssertions; i++) {
            a = assertionsByGuid[bountyGuid][i];

            if (a.verdicts == verdicts) {
                uint256 reward = a.bid.add(split);
                // TODO: Don't revert if one transfer fails, what to do?
                require(token.transfer(a.author, reward));
                pot = pot.sub(reward);
            }
        }

        // Transfer remainder of pot to arbiter, handles fractional NCT remainders
        require(token.transfer(msg.sender, pot));

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
