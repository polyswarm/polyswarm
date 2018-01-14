package bounty

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Keep these in sync with the BountyRegistry contract
type Bounty struct {
	Author          common.Address
	Guid            *big.Int
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
	Verdict         uint8
}

type NewBountyEventLog struct {
	Author          common.Address
	Guid            *big.Int
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
}

// Type generated from WatchForBounties
type BountyEvent *Bounty
