package tests

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"
	"os"
	"path"
	"time"

	. "gopkg.in/check.v1"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/network"
	"github.com/polyswarm/perigord/testing"

	"github.com/polyswarm/polyswarm/bindings"
	"github.com/polyswarm/polyswarm/bounty"
)

const IPFS_URL = "localhost:5001"

type BountyRegistrySuite struct {
	network *network.Network
}

var _ = Suite(&BountyRegistrySuite{})

func (s *BountyRegistrySuite) SetUpTest(c *C) {
	nw, err := testing.SetUpTest()
	if err != nil {
		c.Fatal(err)
	}

	s.network = nw
}

func (s *BountyRegistrySuite) TearDownTest(c *C) {
	testing.TearDownTest()
}

// USER TESTS GO HERE

func getFakeVirusPathAndHash() (string, string) {
	baseDir, _ := os.Getwd()
	p := path.Join(baseDir, "tests", "test_data", "fake_virus")
	f, e := os.Open(p)
	defer f.Close()
	if e != nil {
		panic(e)
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		panic(err)

	}
	return p, fmt.Sprintf("%x", h.Sum(nil))
}

func postFakeVirusBounty(poster *bounty.BountyRegistry) (*big.Int, error) {
	pth, _ := getFakeVirusPathAndHash()
	return poster.PostBounty(context.Background(), pth, 20, 300)
}

func (s *BountyRegistrySuite) TestBountyRegistry(c *C) {
	registry_session, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	c.Assert(registry_session, NotNil)
	c.Assert(ok, Equals, true)

	receiver := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_URL)
	poster := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_URL)

	bountyWatchChan := make(chan bounty.BountyEvent)
	err := receiver.WatchForBounties(bountyWatchChan)
	c.Assert(err, IsNil)

	// sychronous
	guid, err := postFakeVirusBounty(poster)
	c.Assert(err, IsNil)

	ctractBnty := poster.GetActiveBounties()

	foundBounty := false
	for _, b := range ctractBnty {
		if b.Guid.Cmp(guid) == 0 {
			foundBounty = true
		}
	}
	c.Assert(foundBounty, Equals, true)

	stopTimer := time.After(time.Second * 30)
	for {
		select {
		case newBountyStruct, chanOk := <-bountyWatchChan:
			c.Assert(chanOk, Equals, true)
			if newBountyStruct.Guid.Cmp(guid) != 0 {
				break
			}
			c.Log("got matching bounty", newBountyStruct.ArtifactURI)
			return
		case <-stopTimer:
			c.Fatal("Failed to get bounty event in alotted time")
			return
		}
	}

	c.Fatal("Failed to find posted bounty guid")
}

func (s *BountyRegistrySuite) TestBountyRegistryAssert(c *C) {
	registry_session, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	c.Assert(registry_session, NotNil)
	c.Assert(ok, Equals, true)

	receiver := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_URL)
	poster := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_URL)

	bountyWatchChan := make(chan bounty.BountyEvent)
	assertWatchChan := make(chan bounty.AssertionEvent)

	err := receiver.WatchForBounties(bountyWatchChan)
	c.Assert(err, IsNil)

	err = poster.WatchForAssertions(assertWatchChan)
	c.Assert(err, IsNil)

	time.Sleep(time.Second * 10)
	guid, err := postFakeVirusBounty(poster)
	c.Assert(err, IsNil)

	stopChan := time.After(time.Second * 30)
	for {
		select {
		case newBountyStruct := <-bountyWatchChan:
			if newBountyStruct.Guid.Cmp(guid) != 0 {
				break
			}

			err := receiver.PostAssertion(context.Background(), guid, true, 100, "")
			c.Assert(err, IsNil)

			assertTimeout := time.After(time.Second * 20)
			for {
				select {
				case assertEvent := <-assertWatchChan:
					if assertEvent.BountyGuid.Cmp(newBountyStruct.Guid) != 0 {
						break
					}
					c.Log("Suceeded getting assertion back on contract side", assertEvent.BountyGuid.String())
					return
				case <-assertTimeout:
					c.Fatal("Failed to get assertion event in allotted time")
					return
				}
			}

		case <-stopChan:
			c.Fatal("Failed to get bounty event in allotted time")
			return
		}
	}

	c.Fatal("Failed to find posted bounty guid")
}
