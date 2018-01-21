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
	Resolved        bool           `json:"resolved"`
	Verdicts        []bool         `json:"verdicts"`
}

type RawBounty struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
	Resolved        bool
	Verdicts        *big.Int
}

func NewBountyFromRaw(rb RawBounty) *Bounty {
	return &Bounty{
		Guid:            uuid.FromBytesOrNil(rb.Guid.Bytes()),
		Author:          rb.Author,
		Amount:          rb.Amount,
		ArtifactHash:    rb.ArtifactHash,
		ArtifactURI:     rb.ArtifactURI,
		ExpirationBlock: rb.ExpirationBlock,
		Resolved:        rb.Resolved,
		Verdicts:        bigIntToBoolArray(rb.Verdicts),
	}
}

type NewBountyEventLog struct {
	Guid            *big.Int
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    common.Hash
	ArtifactURI     string
	ExpirationBlock *big.Int
}

type NewBountyEvent struct {
	Guid            string
	Author          common.Address
	Amount          *big.Int
	ArtifactHash    common.Hash
	ArtifactURI     string
	ExpirationBlock *big.Int
}

func NewBountyEventFromLog(nbe NewBountyEventLog) *NewBountyEvent {
	return &NewBountyEvent{
		Guid:            uuid.FromBytesOrNil(nbe.Guid.Bytes()).String(),
		Author:          nbe.Author,
		Amount:          nbe.Amount,
		ArtifactHash:    nbe.ArtifactHash,
		ArtifactURI:     nbe.ArtifactURI,
		ExpirationBlock: nbe.ExpirationBlock,
	}
}
