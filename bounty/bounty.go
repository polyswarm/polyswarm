package bounty

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/satori/go.uuid"
)

// Keep these in sync with the BountyRegistry contract
type Bounty struct {
	Guid            uuid.UUID      `json:"guid"`
	Author          common.Address `json:"author"`
	Amount          *big.Int       `json:"amount"`
	ArtifactHash    common.Hash    `json:"hash"`
	ArtifactURI     string         `json:"uri"`
	ExpirationBlock *big.Int       `json:"expiration"`
	Verdict         Verdict        `json:"verdict"`
}

type RawBounty struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
	Verdict         uint8
}

type NewBountyEventLog struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    common.Hash
	ArtifactURI     string
	ExpirationBlock *big.Int
}

func NewBountyFromRaw(rb RawBounty) *Bounty {
	return &Bounty{
		Guid:            uuid.Must(uuid.FromBytes(rb.Guid.Bytes())),
		Author:          rb.Author,
		Amount:          rb.Amount,
		ArtifactHash:    rb.ArtifactHash,
		ArtifactURI:     rb.ArtifactURI,
		ExpirationBlock: rb.ExpirationBlock,
		Verdict:         Verdict(rb.Verdict),
	}
}
