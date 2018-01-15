package bounty

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Keep these in sync with the BountyRegistry contract
type Assertion struct {
	Author   common.Address `json:"author"`
	Verdict  uint8          `json:"verdict"`
	Bid      *big.Int       `json:"bid"`
	Metadata string         `json:"metadata"`
}

type NewAssertionEventLog struct {
	BountyGuid *big.Int
	Author     common.Address
	Verdict    uint8
	Index      *big.Int
	Bid        *big.Int
	Metadata   string
}
