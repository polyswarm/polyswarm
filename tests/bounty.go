package tests

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"

	. "gopkg.in/check.v1"

	_ "github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/testing"

	_ "github.com/polyswarm/polyswarm/bindings"
)

type bountySuite struct {
	network *migration.Network
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

func (s *bountySuite) TestBounty(c *C) {

}
