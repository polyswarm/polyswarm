pragma solidity ^0.4.18;

import "zeppelin-solidity/contracts/math/SafeMath.sol";
import "zeppelin-solidity/contracts/ownership/Ownable.sol";
import "NectarToken.sol";


contract BountyRegistry is Ownable {
    using SafeMath for uint256;

    address internal owner;
    NectarToken token;

    uint256 public constant BOUNTY_LISTING_FEE = 10;
    uint256 public constant BOUNTY_ASSERTION_FEE = 1;
    uint256 public constant BOUNTY_BID_MINIMUM = 1;
    uint256 public constant BOUNTY_AMOUNT_MINIMUM = 1;

    struct Bounty {
        address author;
        uint128 guid;
        uint256 bountyAmount;
        string artifactHash;
        string artifactURI;
        uint256 blockDeadline;
    }

    struct Assertion {
        address author;
        bool malicious;
        uint256 blockTime;
        uint256 assertBid;
        string metadata;
    }
   
    mapping (uint128 => Bounty) public bountiesByGuid;
    mapping (address => Bounty[]) public bountiesByAddress;
    mapping (uint128 => Assertion[]) public assertionsByGuid;
    
    event NewBounty(
        address author,
        uint128 guid,
        uint256 amount,
        uint256 deadlineBlock
    );
    
    event NewAssertion(
        address author,
        uint128 bountyGuid,
        uint256 index
    );

    function BountyRegistry(address nectarTokenAddr) public {
        owner = msg.sender;
        token = NectarToken(nectarTokenAddr);
    }

    function registerAssertion(
        address bountyAuthor, 
        uint128 bountyGuid, 
        bool malicious, 
        uint256 assertBid, 
        string metadata
    )
        public
    {
        // Check if this bounty has been initialized
        require(bountyAuthor != address(0));
        require(bountiesByGuid[bountyGuid].author == bountyAuthor);
        // Check that our bid amount is sufficient
        require(assertBid >= BOUNTY_BID_MINIMUM);

        // Assess fees and transfer bid amount into escrow
        require(token.transferFrom(msg.sender, address(this), assertBid.add(BOUNTY_ASSERTION_FEE))); 
 
        // Instantiate an Assertion and populate it with function arguments.
        Assertion memory a = Assertion(
            msg.sender,
            malicious,
            block.number,
            assertBid,
            metadata
        );

        uint256 index = assertionsByGuid[bountyGuid].push(a) - 1;

        NewAssertion(msg.sender, bountyGuid, index);
    }
    
    function registerBounty(
        uint256 bountyAmt, 
        string aHash, 
        string aURI, 
        uint256 numBlocksDeadline, 
        uint128 guid
    )
        public
    {
        // Check if a bounty with this GUID has already been initialized
        require(bountiesByGuid[guid].author == address(0));
        // Check that our bounty amount is sufficient
        require(bountyAmt >= BOUNTY_AMOUNT_MINIMUM);

        // Assess fees and transfer bounty amount into escrow
        require(token.transferFrom(msg.sender, address(this), bountyAmt.add(BOUNTY_LISTING_FEE))); 
        
        // Instantiate a Bounty and populate with function arguments.
        Bounty memory b = Bounty(
            msg.sender,
            guid,
            bountyAmt,
            aHash,
            aURI,
            block.number + numBlocksDeadline
        );
        bountiesByGuid[guid] = b;
        bountiesByAddress[msg.sender].push(b);
        
        NewBounty(
            msg.sender,
            b.guid,
            b.bountyAmount,
            b.blockDeadline
        );
    }

    // TODO: The final verdict will be determined by arbiters, for now the only
    // vote that counts is the contract owner
    function settleBounty(
        uint128 guid,
        bool malicious
    )
        public
    {
        // Check if this bounty has been initialized
        require(bountiesByGuid[guid].author != address(0));
        // Check if the deadline has expired
        require(bountiesByGuid[guid].blockDeadline <= block.number);

        // Iterate all assertions and pay out accordingly from contract escrow
    }
}
