package bounty

import (
	"context"
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

	"github.com/polyswarm/go-ipfs-api"
	"github.com/polyswarm/perigord/contract"

	"github.com/polyswarm/perigord"
	"github.com/polyswarm/polyswarm/bindings"

	"github.com/satori/go.uuid"
)

type BountyRegistry struct {
	session *bindings.BountyRegistrySession
	client  *ethclient.Client
	ipfssh  *shell.Shell
}

func NewBountyRegistry(session *bindings.BountyRegistrySession, client *ethclient.Client, ipfs string) *BountyRegistry {
	session.TransactOpts.GasLimit = big.NewInt(1000000)

	ipfssh := shell.NewShell(ipfs)
	ipfssh.SetTimeout(15 * time.Second)

	return &BountyRegistry{
		session: session,
		client:  client,
		ipfssh:  ipfssh,
	}
}

func (br *BountyRegistry) UploadArtifacts(rs map[string]io.Reader) (string, error) {
	uris, err := br.ipfssh.AddMultipleWithOpts(rs, true, false, true)
	if err != nil || len(uris) == 0 {
		return "", err
	}

	return uris[len(uris)-1], nil
}

func (br *BountyRegistry) DownloadArtifact(ipfshash string) (io.ReadCloser, error) {
	return br.ipfssh.Cat(ipfshash)
}

func (br *BountyRegistry) ListArtifacts(ipfshash string) ([]string, error) {
	ls, err := br.ipfssh.List(ipfshash)
	if err != nil {
		return []string{}, err
	}

	if len(ls) == 0 {
		return []string{ipfshash}, nil
	}

	var hashes []string
	for _, link := range ls {
		hashes = append(hashes, link.Hash)
	}

	return hashes, nil
}

type ArtifactStats struct {
	Hash           string `json:"hash"`
	BlockSize      int    `json:"block_size"`
	CumulativeSize int    `json:"cumulative_size"`
	DataSize       int    `json:"data_size"`
	NumLinks       int    `json:"num_links"`
}

func (br *BountyRegistry) StatArtifact(ipfshash string) (*ArtifactStats, error) {
	stat, err := br.ipfssh.ObjectStat(ipfshash)
	if err != nil {
		return nil, err
	}

	return &ArtifactStats{
		Hash:           stat.Hash,
		BlockSize:      stat.BlockSize,
		CumulativeSize: stat.CumulativeSize,
		DataSize:       stat.DataSize,
		NumLinks:       stat.NumLinks,
	}, nil
}

func (br *BountyRegistry) PostBounty(ctx context.Context, uri string, amount, blockDuration int) (*big.Int, error) {
	guidInt := new(big.Int)
	guidInt.SetBytes(uuid.Must(uuid.NewV4()).Bytes())

	amountInt := big.NewInt(int64(amount))
	blockDurationInt := big.NewInt(int64(blockDuration))

	_, err := br.session.PostBounty(guidInt, amountInt, uri, blockDurationInt)
	if err != nil {
		return nil, err
	}

	//ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	//defer cancel()
	//_, err = perigord.WaitMined(ctx, br.client, tx)
	return guidInt, nil
}

func (br *BountyRegistry) PostAssertion(ctx context.Context, bountyGuid *big.Int, bid int, mask []bool, verdicts []bool, metadata string) error {
	bidInt := big.NewInt(int64(bid))
	maskInt := boolArrayToBigInt(mask)
	verdictsInt := boolArrayToBigInt(verdicts)

	_, err := br.session.PostAssertion(bountyGuid, bidInt, maskInt, verdictsInt, metadata)
	if err != nil {
		return err
	}

	//ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	//defer cancel()
	//_, err = perigord.WaitMined(ctx, br.client, tx)
	return nil
}

type Event struct {
	Type string
	Body interface{}
}

func (br *BountyRegistry) WatchForEvents(eventChan chan *Event) error {
	topics := map[string]common.Hash{
		"Bounty":    perigord.EventSignatureToTopicHash("NewBounty(uint128,address,uint256,string,uint256)"),
		"Assertion": perigord.EventSignatureToTopicHash("NewAssertion(uint128,address,uint256,uint256,uint256,uint256,string)"),
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
						Body: NewBountyEventFromLog(nbe),
					}

					eventChan <- event
				} else if logMsg.Topics[0] == topics["Assertion"] {
					var nae NewAssertionEventLog
					if err := abi.Unpack(&nae, "NewAssertion", logMsg.Data); err != nil {
						log.Println("error unpacking log:", err)
						break
					}

					// FIXME: Work around bug in decoding string data from event log
					assertion, err := br.session.AssertionsByGuid(nae.BountyGuid, nae.Index)
					if err == nil {
						nae.Metadata = assertion.Metadata
					}

					event := &Event{
						Type: "Assertion",
						Body: NewAssertionEventFromLog(nae),
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
						Body: NewVerdictEventFromLog(nve),
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
