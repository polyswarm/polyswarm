package tests

import (
	"context"
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
	"time"

	. "gopkg.in/check.v1"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/network"
	"github.com/polyswarm/perigord/testing"

	"github.com/polyswarm/polyswarm/bindings"
	"github.com/polyswarm/polyswarm/bounty"
)

const IPFS_HOST = "localhost:5001"

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

func postFakeVirusBounty(poster *bounty.BountyRegistry) (*big.Int, error) {
	baseDir, _ := os.Getwd()
	p := filepath.Join(baseDir, "tests", "test_data", "fake_virus")
	hash, uri, err := poster.Upload(p)
	if err != nil {
		return nil, err
	}

	return poster.PostBounty(context.Background(), hash, uri, 20, 300)
}

func (s *BountyRegistrySuite) TestBountyRegistry(c *C) {
	registry_session, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	c.Assert(registry_session, NotNil)
	c.Assert(ok, Equals, true)

	receiver := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_HOST)
	poster := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_HOST)

	bountyWatchChan := make(chan bounty.BountyEvent)
	err := receiver.WatchForBounties(bountyWatchChan)
	c.Assert(err, IsNil)

	// sychronous
	guid, err := postFakeVirusBounty(poster)
	c.Assert(err, IsNil)

	activeBounties := poster.GetActiveBounties()

	foundBounty := false
	for _, b := range activeBounties {
		if b.Guid.Cmp(guid) == 0 {
			foundBounty = true
			c.Log(json.Marshal(b))
		}
	}
	c.Assert(foundBounty, Equals, true)

	stopTimer := time.After(time.Second * 30)
	for {
		select {
		case bounty, chanOk := <-bountyWatchChan:
			c.Assert(chanOk, Equals, true)
			if bounty.Guid.Cmp(guid) != 0 {
				break
			}

			c.Log("got matching bounty", bounty.ArtifactURI)
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

	receiver := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_HOST)
	poster := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_HOST)

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
		case bounty := <-bountyWatchChan:
			if bounty.Guid.Cmp(guid) != 0 {
				break
			}

			err := receiver.PostAssertion(context.Background(), guid, true, 100, "")
			c.Assert(err, IsNil)

			assertTimeout := time.After(time.Second * 20)
			for {
				select {
				case assertion := <-assertWatchChan:
					if assertion.BountyGuid.Cmp(bounty.Guid) != 0 {
						break
					}
					c.Log("Suceeded getting assertion back on contract side", assertion.BountyGuid.String())
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
