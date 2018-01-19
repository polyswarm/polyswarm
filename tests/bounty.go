package tests

import (
	"context"
	"math/big"
	//	"os"
	//	"path/filepath"

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

//func postFakeVirusBounty(poster *bounty.BountyRegistry) (*big.Int, error) {
//	baseDir, _ := os.Getwd()
//	f, err := os.Open(filepath.Join(baseDir, "tests", "test_data", "fake_virus"))
//	if err != nil {
//		return nil, err
//	}
//
//	hash, uri, err := poster.UploadArtifact(f)
//	if err != nil {
//		return nil, err
//	}
//
//	return poster.PostBounty(context.Background(), hash, uri, 20, 300)
//}

func (s *BountyRegistrySuite) TestPostBounty(c *C) {
	registry_session, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	c.Assert(registry_session, NotNil)
	c.Assert(ok, Equals, true)

	bountyRegistry := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_HOST)

	eventChan := make(chan *bounty.Event)
	err := bountyRegistry.WatchForEvents(eventChan)
	c.Assert(err, IsNil)

	// Post fake bounty
	guid, err := bountyRegistry.PostBounty(context.Background(), [32]byte{}, "uri", 20, 300)
	c.Assert(err, IsNil)

	event := <-eventChan
	b, ok := event.Body.(bounty.NewBountyEventLog)
	c.Assert(b, NotNil)
	c.Assert(ok, Equals, true)
	c.Assert(b.Guid.Cmp(guid), Equals, 0)
}

func (s *BountyRegistrySuite) TestPostAssertion(c *C) {
	registry_session, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	c.Assert(registry_session, NotNil)
	c.Assert(ok, Equals, true)

	bountyRegistry := bounty.NewBountyRegistry(registry_session, s.network.Client(), IPFS_HOST)

	eventChan := make(chan *bounty.Event)
	err := bountyRegistry.WatchForEvents(eventChan)
	c.Assert(err, IsNil)

	// Post fake bounty
	guid, err := bountyRegistry.PostBounty(context.Background(), [32]byte{}, "uri", 20, 300)
	c.Assert(err, IsNil)

	event := <-eventChan
	b, ok := event.Body.(bounty.NewBountyEventLog)
	c.Assert(b, NotNil)
	c.Assert(ok, Equals, true)
	c.Assert(b.Guid.Cmp(guid), Equals, 0)

	err = bountyRegistry.PostAssertion(context.Background(), guid, []bool{true}, 100, "")
	c.Assert(err, IsNil)

	event = <-eventChan
	a, ok := event.Body.(bounty.NewAssertionEventLog)
	c.Assert(b, NotNil)
	c.Assert(ok, Equals, true)
	c.Assert(a.BountyGuid.Cmp(guid), Equals, 0)
	c.Assert(a.Verdicts.Cmp(big.NewInt(1)), Equals, 0)
}
