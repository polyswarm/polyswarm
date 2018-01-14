package bounty

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/satori/go.uuid"
	"io"
	"math/big"
	"os"
	"path/filepath"
)

// Keep these in sync with the BountyRegistry contract
type Bounty struct {
	Author          common.Address
	Guid            *big.Int
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
	Verdict         uint8
}

type NewBountyEventLog struct {
	Author          common.Address
	Guid            *big.Int
	Amount          *big.Int
	ArtifactHash    [32]byte
	ArtifactURI     string
	ExpirationBlock *big.Int
}

// Type generated from WatchForBounties
type BountyEvent *Bounty

func NewBounty(pth string, bountyAmount, blocksFromNow int) (*Bounty, error) {
	b := Bounty{
		Amount:          big.NewInt(int64(bountyAmount)),
		ExpirationBlock: big.NewInt(int64(blocksFromNow)),
		Guid:            new(big.Int),
	}

	b.Guid.SetBytes(uuid.NewV4().Bytes())
	if err := b.SetArtifact(pth); err != nil {
		return nil, err
	}

	return &b, nil
}

func (b *Bounty) SetArtifact(pth string) error {
	pth, err := filepath.Abs(pth)
	if err != nil {
		return err
	}

	f, err := os.Open(pth)
	defer f.Close()
	if err != nil {
		return err
	}

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return err
	}

	// TODO: ipfs
	b.ArtifactURI = "file://" + pth
	copy(b.ArtifactHash[:], h.Sum(nil))

	return nil
}
