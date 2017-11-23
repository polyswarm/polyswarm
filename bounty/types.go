package bounty

import (
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/satori/go.uuid"
	"io"
	"math/big"
	"os"
)

// todo this constantly needs to be kept up to date with the contract abi since it only returns anon structure
type Bounty struct {
	Originator    common.Address
	BountyFee     *big.Int
	BountyAmount  *big.Int
	ArtifactHash  string
	ArtifactURI   string
	BlockDeadline *big.Int
	Guid          *big.Int

	localFilePath string
}

type Assertion struct {
	Author    common.Address
	Malicious bool
	BlockTime *big.Int
	AssertBid *big.Int
	Metadata  string

	// TODO this is not part of dstruct in solidity contract.
	Guid *big.Int
}

func NewAssertion(mal bool, bid int, metadata string) (*Assertion, error) {
	a := Assertion{
		Malicious: mal,
		AssertBid: big.NewInt(int64(bid)),
		Metadata:  metadata,
	}
	return &a, nil

}

func (a *Assertion) SetAuthor(oa common.Address) {
	a.Author = oa
}

func (a *Assertion) SetGuid(g *big.Int) {
	a.Guid = g
}

func NewBounty(pth string, fee, bountyAmount, blocksFromNow int) (*Bounty, error) {
	b := Bounty{}
	b.setGuid()
	b.localFilePath = pth

	b.BountyFee = big.NewInt(int64(fee))
	b.BountyAmount = big.NewInt(int64(bountyAmount))
	b.BlockDeadline = big.NewInt(int64(blocksFromNow))
	if err := b.hashPath(); err != nil {
		return nil, err
	}

	return &b, nil
}

func (b *Bounty) hashPath() error {
	f, e := os.Open(b.localFilePath)
	defer f.Close()
	if e != nil {
		return e
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return e

	}

	b.ArtifactHash = fmt.Sprintf("%x", h.Sum(nil))
	return nil
}

func (b *Bounty) setGuid() {
	bountyGuid := new(big.Int)
	bountyGuid.SetBytes(uuid.NewV4().Bytes())
	b.Guid = bountyGuid
}

func (b *Bounty) Upload() {
	// todo upload this to ipfs etc depending on scheme, otherwise just file me.
	b.ArtifactURI = "file://" + b.localFilePath

}

func (b *Bounty) GuidEq(other *Bounty) bool {
	return other.Guid.Cmp(b.Guid) == 0
}
