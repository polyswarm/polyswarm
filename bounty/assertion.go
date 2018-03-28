package bounty

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	uuid "github.com/satori/go.uuid"
)

// Keep these in sync with the BountyRegistry contract
type Assertion struct {
	Author   common.Address `json:"author"`
	Bid      *big.Int       `json:"bid"`
	Mask     *big.Int       `json:"mask"`
	Verdicts *big.Int       `json:"verdicts"`
	Metadata string         `json:"metadata"`
}

type NewAssertionEventLog struct {
	BountyGuid *big.Int
	Author     common.Address
	Index      *big.Int
	Bid        *big.Int
	Mask       *big.Int
	Verdicts   *big.Int
	Metadata   string
}

type NewAssertionEvent struct {
	BountyGuid string
	Author     common.Address
	Index      *big.Int
	Bid        *big.Int
	Mask       []bool
	Verdicts   []bool
	Metadata   string
}

func NewAssertionEventFromLog(nae NewAssertionEventLog) *NewAssertionEvent {
	return &NewAssertionEvent{
		BountyGuid: uuid.FromBytesOrNil(nae.BountyGuid.Bytes()).String(),
		Author:     nae.Author,
		Index:      nae.Index,
		Bid:        nae.Bid,
		Mask:       bigIntToBoolArray(nae.Mask),
		Verdicts:   bigIntToBoolArray(nae.Verdicts),
		Metadata:   nae.Metadata,
	}
}
