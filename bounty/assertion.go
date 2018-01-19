package bounty

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Keep these in sync with the BountyRegistry contract
type Assertion struct {
	Author   common.Address `json:"author"`
	Verdicts *big.Int       `json:"verdicts"`
	Bid      *big.Int       `json:"bid"`
	Metadata string         `json:"metadata"`
}

type NewAssertionEventLog struct {
	BountyGuid *big.Int
	Author     common.Address
	Verdicts   *big.Int
	Index      *big.Int
	Bid        *big.Int
	Metadata   string
}
