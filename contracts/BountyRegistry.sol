pragma solidity ^0.4.18;

import "zeppelin-solidity/contracts/math/SafeMath.sol";
import "zeppelin-solidity/contracts/lifecycle/Pausable.sol";
import "NectarToken.sol";


contract BountyRegistry is Pausable {
    using SafeMath for uint256;

    enum Verdict {
        Unknown,
        Malicious,
        Benign
    }

    struct Bounty {
        address author;
        uint128 guid;
        uint256 amount;
        bytes32 artifactHash;
        string artifactURI;
        uint256 expirationBlock;
        Verdict verdict;
    }

    struct Assertion {
        address author;
        Verdict verdict;
        uint256 bid;
        string metadata;
        uint256 blockNumber;
    }

    event NewVerdict(
        uint128 bountyGuid,
        Verdict verdict
    );

    event NewBounty(
        address author,
        uint128 guid,
        uint256 amount,
        bytes32 artifactHash,
        string artifactURI,
        uint256 expirationBlock
    );

    event NewAssertion(
        address author,
        Verdict verdict,
        uint128 bountyGuid,
        uint256 index
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
        bytes32 artifactHash,
        string artifactURI,
        uint256 durationBlocks
    )
        external
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
            msg.sender,
            guid,
            amount,
            artifactHash,
            artifactURI,
            durationBlocks.add(block.number),
            Verdict.Unknown
        );
        bountiesByGuid[guid] = b;
        bountyGuids.push(guid);

        NewBounty(
            b.author,
            b.guid,
            b.amount,
            b.artifactHash,
            b.artifactURI,
            b.expirationBlock
        );
    }

    function postAssertion(
        uint128 bountyGuid,
        Verdict verdict,
        uint256 bid,
        string metadata
    )
        external
    {
        // Check if this bounty has been initialized
        require(bountiesByGuid[bountyGuid].author != address(0));
        // Check if this bounty is active
        require(bountiesByGuid[bountyGuid].expirationBlock > block.number);
        // Require our judgement to be one of Malicious or Benign
        require(verdict != Verdict.Unknown);
        // Check that our bid amount is sufficient
        require(bid >= ASSERTION_BID_MINIMUM);

        // Assess fees and transfer bid amount into escrow
        require(token.transferFrom(msg.sender, address(this), bid.add(ASSERTION_FEE)));

        Assertion memory a = Assertion(
            msg.sender,
            verdict,
            bid,
            metadata,
            block.number
        );

        uint256 index = assertionsByGuid[bountyGuid].push(a) - 1;

        NewAssertion(msg.sender, verdict, bountyGuid, index);
    }

    // TODO: The final verdict will be determined by arbiters, for now the only
    // vote that counts is the contract owner
    function settleBounty(
        uint128 guid,
        Verdict verdict
    )
        external
    {
        // Check if this bounty has been initialized
        require(bountiesByGuid[guid].author != address(0));
        // Check if the deadline has expired
        require(bountiesByGuid[guid].expirationBlock <= block.number);

        bountiesByGuid[guid].verdict = verdict;

        // TODO: Iterate all assertions and pay out accordingly from contract escrow

        NewVerdict(guid, verdict);
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
