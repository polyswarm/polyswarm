package bounty

import (
	"bytes"
	"encoding/json"
	"math/big"
	"strings"
)

type Verdict uint8

const (
	Unknown Verdict = iota
	Malicious
	Benign
)

type NewVerdictEventLog struct {
	BountyGuid *big.Int
	Verdict    uint8
}

func (v *Verdict) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("\"")
	if *v == Malicious {
		buffer.WriteString("Malicious")
	} else if *v == Benign {
		buffer.WriteString("Benign")
	} else {
		buffer.WriteString("Unknown")
	}

	buffer.WriteString("\"")
	return buffer.Bytes(), nil
}

func (v *Verdict) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	s = strings.ToLower(s)
	if s == "malicious" {
		*v = Malicious
	} else if s == "benign" {
		*v = Benign
	} else {
		*v = Unknown
	}

	return nil
}
