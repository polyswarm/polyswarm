package bounty

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Keep these in sync with the BountyRegistry contract
type Assertion struct {
	Author      common.Address
	Verdict     uint8
	Bid         *big.Int
	Metadata    string
	BlockNumber *big.Int
}

type NewAssertionEventLog struct {
	Author     common.Address
	Verdict    uint8
	BountyGuid *big.Int
	Index      *big.Int
}

// Type generated from WatchForAssertions
type AssertionEvent struct {
	Assertion  *Assertion
	BountyGuid *big.Int
}

func NewAssertion(mal bool, bid int, metadata string) *Assertion {
	verdict := Malicious
	if !mal {
		verdict = Benign
	}

	return &Assertion{
		Verdict:  verdict,
		Bid:      big.NewInt(int64(bid)),
		Metadata: metadata,
	}
}
