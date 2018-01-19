package bounty

import (
	"math/big"
)

type NewVerdictEventLog struct {
	BountyGuid *big.Int
	Verdicts   *big.Int
}
