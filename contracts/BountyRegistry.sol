pragma solidity ^0.4.18;

import "zeppelin-solidity/contracts/math/SafeMath.sol";
import "zeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "NectarToken.sol";


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

    function BountyRegistry(address nectarTokenAddr) public {
        owner = msg.sender;
        token = NectarToken(nectarTokenAddr);
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

    // TODO: The final verdict will be determined by arbiters, for now the only
    // vote that counts is the contract owner
    function settleBounty(
        uint128 guid,
        uint256 verdicts
    )
        external
        whenNotPaused
    {
        // Check if this bounty has been initialized
        require(bountiesByGuid[guid].author != address(0));
        // Check if the deadline has expired
        require(bountiesByGuid[guid].expirationBlock <= block.number);

        bountiesByGuid[guid].verdicts = verdicts;
        bountiesByGuid[guid].resolved = true;

        // TODO: Iterate all assertions and pay out accordingly from contract escrow

        NewVerdict(guid, verdicts);
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
