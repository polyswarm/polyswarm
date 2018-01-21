package bounty

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	uuid "github.com/satori/go.uuid"
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

type NewAssertionEvent struct {
	BountyGuid string
	Author     common.Address
	Verdicts   []bool
	Index      *big.Int
	Bid        *big.Int
	Metadata   string
}

func NewAssertionEventFromLog(nae NewAssertionEventLog) *NewAssertionEvent {
	return &NewAssertionEvent{
		BountyGuid: uuid.FromBytesOrNil(nae.BountyGuid.Bytes()).String(),
		Author:     nae.Author,
		Verdicts:   bigIntToBoolArray(nae.Verdicts),
		Index:      nae.Index,
		Bid:        nae.Bid,
		Metadata:   nae.Metadata,
	}
}
