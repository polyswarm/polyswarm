package tests

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
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

type bountySuite struct {
	network *network.Network
}

var _ = Suite(&bountySuite{})

func (s *bountySuite) SetUpTest(c *C) {
	network, err := testing.SetUpTest()
	if err != nil {
		c.Fail()
	}

	s.network = network
}

func (s *bountySuite) TearDownTest(c *C) {
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

func postFakeVirusBounty(poster *bounty.BountyPoster) (*bounty.Bounty, error) {
	pth, _ := getFakeVirusPathAndHash()
	bnty, err := bounty.NewBounty(pth, 12, 20, 10)
	if err != nil {
		return nil, err
	}
	bnty.Upload()
	_, err = poster.PostBounty(context.Background(), bnty)
	if err != nil {
		return nil, err
	}

	return bnty, nil
}

func (s *bountySuite) TestBountyPoster(c *C) {
	registry_session, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	c.Assert(registry_session, NotNil)
	c.Assert(ok, Equals, true)

	receiver := bounty.NewBountyPoster(registry_session, s.network.Client())
	poster := bounty.NewBountyPoster(registry_session, s.network.Client())

	bountyWatchChan := make(chan *bounty.Bounty)
	err := receiver.WatchForBounties(bountyWatchChan)
	c.Assert(err, IsNil)

	// sychronous
	bnty, err := postFakeVirusBounty(poster)
	c.Assert(err, IsNil)

	ctractBnty, err := poster.GetActiveBounties()
	c.Assert(err, IsNil)

	foundBounty := false
	for _, b := range ctractBnty {
		if b.GuidEq(bnty) {
			foundBounty = true
		}
	}
	c.Assert(foundBounty, Equals, true)

	stopTimer := time.After(time.Second * 30)
	for {
		select {
		case newBountyStruct, chanOk := <-bountyWatchChan:
			c.Assert(chanOk, Equals, true)
			if newBountyStruct.Guid.Cmp(bnty.Guid) != 0 {
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

func (s *bountySuite) TestBountyPosterAssert(c *C) {
	registry_session, ok := contract.Session("BountyRegistry").(*bindings.BountyRegistrySession)
	c.Assert(registry_session, NotNil)
	c.Assert(ok, Equals, true)

	receiver := bounty.NewBountyPoster(registry_session, s.network.Client())
	poster := bounty.NewBountyPoster(registry_session, s.network.Client())

	bountyWatchChan := make(chan *bounty.Bounty)
	assertWatchChan := make(chan *bounty.Assertion)

	err := receiver.WatchForBounties(bountyWatchChan)
	c.Assert(err, IsNil)

	err = poster.WatchForAssertions(assertWatchChan)
	c.Assert(err, IsNil)

	time.Sleep(time.Second * 10)
	bnty, err := postFakeVirusBounty(poster)
	c.Assert(err, IsNil)

	stopChan := time.After(time.Second * 30)
	for {
		select {
		case newBountyStruct := <-bountyWatchChan:
			if newBountyStruct.Guid.Cmp(bnty.Guid) != 0 {
				break
			}

			asrt, _ := bounty.NewAssertion(true, 100, "")
			asrt.SetGuid(newBountyStruct.Guid)
			rcpt, err := receiver.PostAssertion(context.Background(), newBountyStruct, asrt)
			c.Assert(err, IsNil)

			c.Log("assert receipt", rcpt.String())
			assertTimeout := time.After(time.Second * 20)
			for {
				select {
				case newAssertStruct := <-assertWatchChan:
					if newAssertStruct.Guid.Cmp(asrt.Guid) != 0 || newAssertStruct.Guid.Cmp(bnty.Guid) != 0 {
						break
					}
					c.Log("Suceeded getting assertion back on contract side", newAssertStruct.Guid.String())
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
