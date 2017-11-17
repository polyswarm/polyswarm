// TODO: come up with a better name than bounty "amount"
// TODO: add log messages
// TODO: always be explicit between 'storage' and 'memory'

pragma solidity ^0.4.15;

contract BountyRegistry {
	
	
	//// 
	// VARIABLES
	////
	// TODO: make scope for all of these explicit
	address owner;
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
	
	// bounties is a map from addresses to a list of Bountys
	mapping (address => Bounty[]) public bounties;

	// assertions is a map from GUID to a list of Assertions
	// it maps to a list because Go abigen doesn't do nested types well.
	mapping (uint128 => Assertion[]) public assertions;
	
	// warning! if you change event signatures you will have to extract log topics from unit tests and reimplement in bounty data...
	// TODO fix this...
	
	
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
        	
		Bounty[] memory authorBounties = bounties[bountyAuthor];

		// In mappings, returned value is initialized with 0 if entry doesn't exist.
		// Seems the current best approach is to check a field that *must* be non-zero

		// Ensure bounty author has submitted at least one bounty
		require(authorBounties.length != 0x0);
        
		// TODO: we have a loop of indeterminate length in a function that requires Gas
		// What happens if we run out of Gas during this loop?
		// TODO: check for overflow of i
		bool found = false; // TODO: scope
		for (uint i = 0; i < authorBounties.length; i++) {
			
			Bounty memory candidateBounty = authorBounties[i];
			if (candidateBounty.guid == bountyGuid) {
				
				// Instantiate an Assertion and populate it with function arguments.
				Assertion memory a;
				a.author = msg.sender;
				a.malicious = malicious;
				a.blockTime = block.number;
				a.assertBid = assertBid;
				a.metadata = metadata;

                		// TODO ensure we remove these when bounty is complete
				
				assertions[candidateBounty.guid].push(a);
				// TODO: make an explicit size
				var cLen = assertions[candidateBounty.guid].length-1;
				
				// TODO: check return value
				NewAssertion(msg.sender, bountyGuid, cLen);
				
				found = true;
				break;
			}
		}

		// If the bounty GUID wasn't found, refund Gas and return something useful.
		require(found);
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
		// TODO: ensure no duplicate bounty GUIDs

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
		
		bounties[msg.sender].push(b);
		
		uint cLen = bounties[msg.sender].length-1;
		
		// TODO: check return value
		NewBounty(msg.sender, cLen, b.bountyAmount, b.blockDeadline);
	}
}

