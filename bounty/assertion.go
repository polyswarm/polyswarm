package bounty

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Keep these in sync with the BountyRegistry contract
type Assertion struct {
	Author    common.Address
	Malicious bool
	BlockTime *big.Int
	AssertBid *big.Int
	Metadata  string

	BountyGuid *big.Int
}

type NewAssertionEvent struct {
	Author     common.Address
	BountyGuid *big.Int
	Index      *big.Int
}

func NewAssertion(mal bool, bid int, metadata string) *Assertion {
	return &Assertion{
		Malicious: mal,
		AssertBid: big.NewInt(int64(bid)),
		Metadata:  metadata,
	}
}
