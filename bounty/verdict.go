package bounty

import (
	"math/big"
)

const (
	Unknown uint8 = iota
	Malicious
	Benign
)

type NewVerdictEventLog struct {
	BountyGuid *big.Int
	Verdict    uint8
}
