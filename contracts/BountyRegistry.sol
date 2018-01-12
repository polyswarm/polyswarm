pragma solidity ^0.4.18;

contract BountyRegistry {
    //// 
    // VARIABLES
    ////
    address internal owner;
    uint256 public constant BOUNTY_LISTING_FEE = 10;
    uint256 public constant BOUNTY_ASSERTION_FEE = 1;
    uint256 public constant BOUNTY_BID_MINIMUM = 1;
    uint256 public constant BOUNTY_AMOUNT_MINIMUM = 1;

    struct Bounty {
        address author;
        uint128 guid;
        uint256 bountyFee;
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
    
    ////
    // EVENTS
    ////
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

    //// 
    // CONSTRUCTOR
    ////
    function BountyRegistry() public payable {
        owner = msg.sender;
    }

    ////
    // PUBLIC FUNCTIONS
    ////
    function registerAssertion(
        address bountyAuthor, 
        uint128 bountyGuid, 
        bool malicious, 
        uint256 assertBid, 
        string metadata
    )
        public
        payable
    {
        // TODO check this bid and make transfer into contract escrow
        require(assertBid >= BOUNTY_BID_MINIMUM);
        require(bountyAuthor != address(0));
        require(bountiesByGuid[bountyGuid].author == bountyAuthor);
            
        // Instantiate an Assertion and populate it with function arguments.
        Assertion memory a;
        a.author = msg.sender;
        a.malicious = malicious;
        a.blockTime = block.number;
        a.assertBid = assertBid;
        a.metadata = metadata;

        uint256 index = assertionsByGuid[bountyGuid].push(a) - 1;

        NewAssertion(msg.sender, bountyGuid, index);
    }
    
    function registerBounty(
        uint256 bFee, 
        uint256 bountyAmt, 
        string aHash, 
        string aURI, 
        uint256 numBlocksDeadline, 
        uint128 guid
    )
        public
        payable
    {
        // Check if a bounty with this GUID has already been initialized
        require(bountiesByGuid[guid].author == address(0));

        // Check whether posted Fee is sufficient.
        // TODO: check that ambassador has enough token as second part here.
        require(bFee >= BOUNTY_LISTING_FEE);
        require(bountyAmt >= BOUNTY_AMOUNT_MINIMUM);
        
        // Instantiate a Bounty and populate with function arguments.
        Bounty memory b;
        b.author = msg.sender;
        b.guid = guid;
        b.blockDeadline = block.number + numBlocksDeadline;
        b.artifactHash = aHash;
        b.artifactURI = aURI;
        b.bountyAmount = bountyAmt;
        b.bountyFee = bFee;
        
        bountiesByGuid[guid] = b;
        bountiesByAddress[msg.sender].push(b);
        
        NewBounty(msg.sender, b.guid, b.bountyAmount, b.blockDeadline);
    }
}

