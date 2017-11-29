// TODO: come up with a better name than bounty "amount"
// TODO: add log messages
// TODO: always be explicit between 'storage' and 'memory'

pragma solidity ^0.4.18;

contract BountyRegistry {
	//// 
	// VARIABLES
	////
	address internal owner;
	uint256 public constant BOUNTY_LISTING_FEE = 10;
	uint256 public constant BOUNTY_ASSERTION_FEE = 1;
	uint public constant BOUNTY_BID_MINIMUM = 1;
	uint public constant BOUNTY_AMOUNT_MINIMUM = 1;

	struct Assertion {
		address author;
		bool malicious;
		uint256 blockTime;
		uint256 assertBid;
		string metadata;
	}
	
	struct Bounty {
		address originator;
		uint256 bountyFee;
		uint256 bountyAmount;
		string artifactHash;
		string artifactURI;
		uint256 blockDeadline; // block number
		uint128 guid;
	}
	
	// bountiesByAddress is a map from addresses to a list of Bountys
	mapping (address => Bounty[]) public bountiesByAddress;

	// bountyByGuid is a map from GUID to a Bounty
	mapping (uint128 => Bounty) public bountyByGuid;

	// assertions is a map from GUID to a list of Assertions
	mapping (uint128 => Assertion[]) public assertions;
	
	////
	// EVENTS
	////
	event NewBounty (
		address originator,  
		uint num, 
		uint256 amount, 
		uint256 deadlineBlock
	);
	
	event NewAssertion (
		address author, 
		uint128 guid, 
		uint num
	);
		
	//// 
	// CONSTRUCTOR
	////
	function BountyRegistry () 
	// returns nil
	{
		owner = msg.sender;
	}

	////
	// PUBLIC FUNCTIONS
	////
	function registerAssertion (
		address bountyAuthor, 
		uint128 bountyGuid, 
		bool malicious, 
		uint256 assertBid, 
		string metadata)
	public
	// Transactional (requires Gas)
	// returns nil
	{
		// TODO check this bid and make transfer into contract escrow
		require(assertBid >= BOUNTY_BID_MINIMUM);
        require(bountyAuthor != address(0));
        require(bountyByGuid[bountyGuid].originator == bountyAuthor);
        	
        // Instantiate an Assertion and populate it with function arguments.
        Assertion memory a;
        a.author = msg.sender;
        a.malicious = malicious;
        a.blockTime = block.number;
        a.assertBid = assertBid;
        a.metadata = metadata;

        assertions[bountyGuid].push(a);
        uint cLen = assertions[bountyGuid].length-1;

        // TODO: check return value
        NewAssertion(msg.sender, bountyGuid, cLen);
	}
	
	function registerBounty (
		uint256 bFee, 
		uint256 bountyAmt, 
		string aHash, 
		string aURI, 
		uint256 numBlocksDeadline, 
		uint128 guid) 
	public
	// Transactional (requires Gas)
	// returns nil	
	{
        // Check if a bounty with this GUID has already been initialized
        require(bountyByGuid[guid].originator == address(0));

		// Check whether posted Fee is sufficient.
		// TODO: check that ambassador has enough token as second part here.
		require(bFee >= BOUNTY_LISTING_FEE);
		require(bountyAmt >= BOUNTY_AMOUNT_MINIMUM);
		
		// Instantiate a Bounty and populate with function arguments.
		Bounty memory b;
		b.originator = msg.sender;
		// TODO: check for overflow?
		b.blockDeadline = block.number + numBlocksDeadline;
		b.artifactHash = aHash;
		b.artifactURI = aURI;
		b.bountyAmount = bountyAmt;
		b.bountyFee = bFee;
		b.guid = guid;
		
        bountyByGuid[guid] = b;
		bountiesByAddress[msg.sender].push(b);
		uint cLen = bountiesByAddress[msg.sender].length-1;
		
		// TODO: check return value
		NewBounty(msg.sender, cLen, b.bountyAmount, b.blockDeadline);
	}
}

