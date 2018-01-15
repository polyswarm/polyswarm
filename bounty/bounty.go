package bounty

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Keep these in sync with the BountyRegistry contract
type Bounty struct {
	Guid            *big.Int       `json:"guid"`
	Author          common.Address `json:"author"`
	Amount          *big.Int       `json:"amount"`
	ArtifactHash    [32]byte       `json:"hash"`
	ArtifactURI     string         `json:"uri"`
	ExpirationBlock *big.Int       `json:"expiration"`
	Verdict         uint8          `json:"verdict"`
}

type NewBountyEventLog struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
}
