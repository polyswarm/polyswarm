package bounty

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Keep these in sync with the BountyRegistry contract
type Assertion struct {
	Author   common.Address `json:"author"`
	Verdict  Verdict        `json:"verdict"`
	Bid      *big.Int       `json:"bid"`
	Metadata string         `json:"metadata"`
}

type RawAssertion struct {
	Author   common.Address
	Verdict  uint8
	Bid      *big.Int
	Metadata string
}

type NewAssertionEventLog struct {
	BountyGuid *big.Int
	Author     common.Address
	Verdict    uint8
	Index      *big.Int
	Bid        *big.Int
	Metadata   string
}

func NewAssertionFromRaw(ra RawAssertion) *Assertion {
	return &Assertion{
		Author:   ra.Author,
		Verdict:  Verdict(ra.Verdict),
		Bid:      ra.Bid,
		Metadata: ra.Metadata,
	}
}
