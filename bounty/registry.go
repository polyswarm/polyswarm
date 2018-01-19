package bounty

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ipfs/go-ipfs-api"

	"github.com/polyswarm/perigord"
	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/polyswarm/bindings"

	"github.com/satori/go.uuid"
)

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
		ret = append(ret, v.Bit(i) == 1)
	}
	return ret
}

type BountyRegistry struct {
	session *bindings.BountyRegistrySession
	client  *ethclient.Client
	ipfssh  *shell.Shell
}

func NewBountyRegistry(session *bindings.BountyRegistrySession, client *ethclient.Client, ipfs string) *BountyRegistry {
	session.TransactOpts.GasLimit = 1000000
	return &BountyRegistry{
		session: session,
		client:  client,
		ipfssh:  shell.NewShell(ipfs),
	}
}

func (br *BountyRegistry) UploadArtifact(r io.Reader) (common.Hash, string, error) {
	hashR, hashW := io.Pipe()
	ipfsR, ipfsW := io.Pipe()

	hashChan := make(chan common.Hash, 1)
	uriChan := make(chan string, 1)
	doneChan := make(chan bool, 2)
	errorChan := make(chan error, 2)

	go func() {
		hash := common.Hash{}
		h := sha256.New()
		if _, err := io.Copy(h, hashR); err != nil {
			errorChan <- err
		}
		copy(hash[:], h.Sum(nil))
		hashChan <- hash
		doneChan <- true
	}()

	go func() {
		ipfs_hash, err := br.ipfssh.Add(ipfsR)
		if err != nil {
			errorChan <- err
		}
		uriChan <- ipfs_hash
		doneChan <- true
	}()

	go func() {
		defer hashW.Close()
		defer ipfsW.Close()

		mw := io.MultiWriter(hashW, ipfsW)
		io.Copy(mw, r)
	}()

	for c := 0; c < 2; c++ {
		select {
		case _ = <-doneChan:
			break
		case err := <-errorChan:
			return common.Hash{}, "", err
		}
	}

	return <-hashChan, <-uriChan, nil
}

func (br *BountyRegistry) DownloadArtifact(ipfshash string) (io.ReadCloser, error) {
	return br.ipfssh.Cat(ipfshash)
}

func (br *BountyRegistry) PostBounty(ctx context.Context, hash common.Hash, uri string, amount, blockDuration int) (*big.Int, error) {
	guidInt := new(big.Int)
	guidInt.SetBytes(uuid.Must(uuid.NewV4()).Bytes())

	amountInt := big.NewInt(int64(amount))
	blockDurationInt := big.NewInt(int64(blockDuration))

	tx, err := br.session.PostBounty(guidInt, amountInt, hash, uri, blockDurationInt)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	_, err = perigord.WaitMined(ctx, br.client, tx)
	return guidInt, err
}

func (br *BountyRegistry) PostAssertion(ctx context.Context, bountyGuid *big.Int, verdicts []bool, bid int, metadata string) error {
	bidInt := big.NewInt(int64(bid))
	verdictsInt := boolArrayToBigInt(verdicts)
	log.Println(verdictsInt)

	tx, err := br.session.PostAssertion(bountyGuid, verdictsInt, bidInt, metadata)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	_, err = perigord.WaitMined(ctx, br.client, tx)
	return err
}

type Event struct {
	Type string
	Body interface{}
}

func (br *BountyRegistry) WatchForEvents(eventChan chan *Event) error {
	topics := map[string]common.Hash{
		"Bounty":    perigord.EventSignatureToTopicHash("NewBounty(uint128,address,uint256,bytes32,string,uint256)"),
		"Assertion": perigord.EventSignatureToTopicHash("NewAssertion(uint128,address,uint256,uint256,uint256,string)"),
		"Verdict":   perigord.EventSignatureToTopicHash("NewVerdict(uint128,uint256)"),
	}

	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics: [][]common.Hash{{
			topics["Bounty"],
			topics["Assertion"],
			topics["Verdict"],
		}},
	}

	dec := json.NewDecoder(strings.NewReader(bindings.BountyRegistryABI))
	var abi abi.ABI
	if err := dec.Decode(&abi); err != nil {
		return err
	}

	logChan := make(chan types.Log)
	sub, err := br.client.SubscribeFilterLogs(context.Background(), q, logChan)
	if err != nil {
		return err
	}

	go func() {
		log.Println("Starting event monitor")
		for {
			select {
			case logMsg := <-logChan:
				if len(logMsg.Topics) != 1 {
					log.Println("incorrect number of topics")
					break
				}

				if logMsg.Topics[0] == topics["Bounty"] {
					var nbe NewBountyEventLog
					if err := abi.Unpack(&nbe, "NewBounty", logMsg.Data); err != nil {
						log.Println("error unpacking log:", err)
						break
					}

					event := &Event{
						Type: "Bounty",
						Body: nbe,
					}

					eventChan <- event
				} else if logMsg.Topics[0] == topics["Assertion"] {
					var nae NewAssertionEventLog
					if err := abi.Unpack(&nae, "NewAssertion", logMsg.Data); err != nil {
						log.Println("error unpacking log:", err)
						break
					}

					event := &Event{
						Type: "Assertion",
						Body: nae,
					}

					eventChan <- event
				} else if logMsg.Topics[0] == topics["Verdict"] {
					var nve NewVerdictEventLog
					if err := abi.Unpack(&nve, "NewVerdict", logMsg.Data); err != nil {
						log.Println("error unpacking log:", err)
						break
					}

					event := &Event{
						Type: "Verdict",
						Body: nve,
					}

					eventChan <- event
				}
				break
			case err := <-sub.Err():
				log.Println(err)
				break
			}
		}
	}()

	return nil
}

func (br *BountyRegistry) GetBounties() []*Bounty {
	ret := make([]*Bounty, 0)
	for i := 0; ; i++ {
		bountyGuid, err := br.session.BountyGuids(big.NewInt(int64(i)))
		if err != nil {
			break
		}

		rawBounty, err := br.session.BountiesByGuid(bountyGuid)
		if err != nil {
			break
		}

		ret = append(ret, NewBountyFromRaw(RawBounty(rawBounty)))
	}

	return ret
}

func (br *BountyRegistry) GetActiveBounties() []*Bounty {
	ret := make([]*Bounty, 0)

	block, err := br.client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return ret
	}
	latest := block.Number()

	bounties := br.GetBounties()
	for _, b := range bounties {
		if b.ExpirationBlock.Cmp(latest) >= 0 {
			ret = append(ret, b)
		}
	}

	return ret
}

func (br *BountyRegistry) GetBountyByGuid(guid *big.Int) *Bounty {
	rawBounty, err := br.session.BountiesByGuid(guid)
	if err != nil {
		return nil
	}

	return NewBountyFromRaw(RawBounty(rawBounty))
}

func (br *BountyRegistry) GetAssertionsByGuid(guid *big.Int) []*Assertion {
	ret := make([]*Assertion, 0)
	for i := 0; ; i++ {
		rawAssertion, err := br.session.AssertionsByGuid(guid, big.NewInt(int64(i)))
		if err != nil {
			break
		}

		assertion := Assertion(rawAssertion)
		ret = append(ret, &assertion)
	}

	return ret
}
