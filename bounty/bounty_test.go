package bounty

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/polyswarm/perigord/contract"
)

// todo load testkey text from dir
const TEST_KEY = `{"address":"5b35a8dbe81e368a991afbf26eff6c607ca1bf2c","crypto":{"cipher":"aes-128-ctr","ciphertext":"29bda81400fe510778c5e35e740e55641c9f4205fa8ddd692836bea24d46e8ee","cipherparams":{"iv":"6a740ef61907470518ea9582d60936c5"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"4fef0536bbaa8eeaba42b6d7e4aceaed40ff94d436b4c5b175b2f0a9ff18208d"},"mac":"1e1e990b499d9a32178676e8f09c580bc98deb7c4d80ae412a08a0aa4cc27fd3"},"id":"e6287b13-4ce2-48c3-8169-5417c2b8e2fe","version":3}`
const TEST_PASS = "blah"

func getFakeVirusPathAndHash() (string, string) {
	baseDir, _ := os.Getwd()
	p := path.Join(baseDir, "test_data", "fake_virus")
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

func TestBountyPosterClass(t *testing.T) {
	conn, err := ethclient.Dial("/home/user/etherdev/geth.ipc")

	err, t, poster, e, receiver := getBountyPosterAndReceiver(err, t, conn)
	if e != nil {
		t.Fatal(e)
		t.FailNow()
	}

	bountyWatchChan := make(chan *Bounty)

	bwe := receiver.WatchForBounties(context.Background(), bountyWatchChan)
	if bwe != nil {
		t.Fatal(bwe)
		t.FailNow()
	}
	// end recipient setup

	bnty, err := postFakeVirusBounty(t, poster)
	// above is sychronous
	ctractBnty, err := poster.GetActiveBounties()
	if err != nil {
		t.Fatalf("failed to get active bounties %v", err)
		t.FailNow()
	}
	foundBounty := false
	for _, b := range ctractBnty {
		if b.GuidEq(bnty) {
			t.Log("found bounty", b)
			foundBounty = true

		}
	}
	if !foundBounty {
		t.Fatal("Couldn't find bounty in poll")
		t.FailNow()
	}

	stopTimer := time.After(time.Second * 20)
	for {
		select {
		case newBountyStruct, chanOk := <-bountyWatchChan:
			if !chanOk {
				t.Fatal("chan bad")
				t.FailNow()
			}
			if newBountyStruct.Guid.Cmp(bnty.Guid) != 0 {
				break
			}
			t.Log("got matching bounty", newBountyStruct.ArtifactURI)
			return
		case <-stopTimer:
			t.Fatal("Failed to get bounty event in alotted time")
			t.FailNow()
			return

		}

	}

	t.Fatal("Failed to find posted bounty guid")
	t.FailNow()

}

func TestBountyPosterAssertTooClass(t *testing.T) {
	conn, err := ethclient.Dial("/home/user/etherdev/geth.ipc")

	err, t, poster, _, receiver := getBountyPosterAndReceiver(err, t, conn)

	bountyWatchChan := make(chan *Bounty)

	assertWatchChan := make(chan *Assertion)

	bwe := receiver.WatchForBounties(context.Background(), bountyWatchChan)
	if bwe != nil {
		t.Fatal(bwe)
		t.FailNow()
	}
	// end recipient setup

	bwa := poster.WatchForAssertions(context.Background(), assertWatchChan)
	if bwa != nil {
		t.Fatal(bwa)
		t.FailNow()
	}
	time.Sleep(time.Millisecond * 100)
	bnty, err := postFakeVirusBounty(t, poster)

	stopChan := time.After(time.Second * 30)
	for {
		select {
		case newBountyStruct := <-bountyWatchChan:
			if newBountyStruct.Guid.Cmp(bnty.Guid) != 0 {
				break
			}
			// todo post assertion

			asrt, _ := NewAssertion(true, 100, "")

			asrt.SetGuid(newBountyStruct.Guid)

			rcpt, err := receiver.PostAssertion(context.Background(), newBountyStruct, asrt)

			if err != nil {
				t.Fatal(err)
				t.FailNow()
				return
			}
			t.Log("assert receipt", rcpt.String())
			assertTimeout := time.After(time.Second * 20)
			for {
				select {
				case newAssertStruct := <-assertWatchChan:
					if newAssertStruct.Guid.Cmp(asrt.Guid) != 0 || newAssertStruct.Guid.Cmp(bnty.Guid) != 0 {
						break
					}
					t.Log("Suceeded getting assertion back on contract side", newAssertStruct.Guid.String())
					return
				case <-assertTimeout:
					t.Fatal("Failed to get assertion event in allotted time")
					t.FailNow()
					return
				}
			}

		case <-stopChan:
			t.Fatal("Failed to get bounty event in allotted time")
			t.FailNow()
			return

		}
	}

	t.Fatal("Failed to find posted bounty guid")
	t.FailNow()
}

func postFakeVirusBounty(t *testing.T, poster *BountyPoster) (*Bounty, error) {
	pth, _ := getFakeVirusPathAndHash()
	bnty, err := NewBounty(pth, 12, 20, 10)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	bnty.Upload()
	rcpt, pE := poster.PostBounty(context.Background(), bnty)

	if pE != nil {
		t.Fatal(pE)
		t.FailNow()
	}
	t.Log("fake bounty recpt", rcpt.String())
	return bnty, err
}
func getBountyPosterAndReceiver(err error, t *testing.T, conn *ethclient.Client) (error, *testing.T, *BountyPoster, error, *BountyPoster) {
	addr := contract.AddressOf("BountyRegistry")
	if addr == nil {
		t.Fatal("contract not deployed")
		t.FailNow()
	}
	bp, e := NewBountyPoster(conn, TEST_KEY, TEST_PASS)
	if e != nil {
		t.Fatal(e)
		t.FailNow()
	}

	bp.SetRegistry(*addr)
	// start recipient setup
	br, e := NewBountyPoster(conn, TEST_KEY, TEST_PASS)
	if e != nil {
		t.Fatal(e)
		t.FailNow()
	}

	return err, t, bp, e, br
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)
