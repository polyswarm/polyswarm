package bounty

import (
	"math/big"

	uuid "github.com/satori/go.uuid"
)

type NewVerdictEventLog struct {
	BountyGuid *big.Int
	Verdicts   *big.Int
}

type NewVerdictEvent struct {
	BountyGuid string
	Verdicts   []bool
}

func NewVerdictEventFromLog(nve NewVerdictEventLog) *NewVerdictEvent {
	return &NewVerdictEvent{
		BountyGuid: uuid.FromBytesOrNil(nve.BountyGuid.Bytes()).String(),
		Verdicts:   bigIntToBoolArray(nve.Verdicts),
	}
}

func boolArrayToBigInt(b []bool) *big.Int {
	bytes := make([]byte, len(b)/8+1)
	var cur byte

	for i, v := range b {
		cur = cur << 1
		if v {
			cur = cur | 1
		}

		if i%8 == 7 {
			bytes = append(bytes, cur)
			cur = 0
		}
	}

	if len(b)%8 != 0 {
		bytes = append(bytes, cur)
	}

	ret := new(big.Int)
	ret.SetBytes(bytes)
	return ret
}

func bigIntToBoolArray(v *big.Int) []bool {
	len := v.BitLen()
	ret := make([]bool, len)
	for i := 0; i < len; i++ {
		ret[i] = v.Bit(i) == 1
	}
	return ret
}
