package bounty

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"os"
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

func (br *BountyRegistry) Upload(path string) ([32]byte, string, error) {
	hash := [32]byte{}

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return hash, "", err
	}

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return hash, "", err
	}
	copy(hash[:], h.Sum(nil))

	f.Seek(0, io.SeekStart)
	ipfs_hash, err := br.ipfssh.Add(f)
	if err != nil {
		return hash, "", err
	}

	return hash, "ipfs:" + ipfs_hash, nil
}

func (br *BountyRegistry) PostBounty(ctx context.Context, hash [32]byte, uri string, amount, blockDuration int) (*big.Int, error) {
	guidInt := new(big.Int)
	guidInt.SetBytes(uuid.Must(uuid.NewV4()).Bytes())

	amountInt := big.NewInt(int64(amount))
	blockDurationInt := big.NewInt(int64(amount))

	tx, err := br.session.PostBounty(guidInt, amountInt, hash, uri, blockDurationInt)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	_, err = perigord.WaitMined(ctx, br.client, tx)
	return guidInt, err
}

func (br *BountyRegistry) PostAssertion(ctx context.Context, bountyGuid *big.Int, malicious bool, bid int, metadata string) error {
	verdict := Malicious
	if !malicious {
		verdict = Benign
	}

	bidInt := big.NewInt(int64(bid))

	tx, err := br.session.PostAssertion(bountyGuid, verdict, bidInt, metadata)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	_, err = perigord.WaitMined(ctx, br.client, tx)
	return err
}

func (br *BountyRegistry) WatchForBounties(bChan chan BountyEvent) error {
	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics:    [][]common.Hash{{perigord.EventSignatureToTopicHash("NewBounty(address,uint128,uint256,bytes32,string,uint256)")}},
	}

	logChan := make(chan types.Log)
	sub, err := br.client.SubscribeFilterLogs(context.Background(), q, logChan)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(strings.NewReader(bindings.BountyRegistryABI))
	var abi abi.ABI
	if err := dec.Decode(&abi); err != nil {
		return err
	}

	go func() {
		log.Println("Started event monitor")

		for {
			select {
			case logMsg := <-logChan:
				var nbe NewBountyEventLog
				if err := abi.Unpack(&nbe, "NewBounty", logMsg.Data); err != nil {
					log.Println("Ignoring event non-bounty post: ", err)
				} else {
					log.Println("New bounty amount addr", nbe.Amount.String(), nbe.Author.String())

					bountyStruct, err := br.session.BountiesByGuid(nbe.Guid)
					if err != nil {
						log.Fatalln("Failed to get bounty based on event: ", err)
					}

					nb := Bounty(bountyStruct)
					bChan <- &nb

					log.Println("Bounty URI", nb.ArtifactURI, nb.ArtifactHash)
				}
				break
			case err := <-sub.Err():
				log.Fatalln(err)
				return
			}
		}
	}()

	return nil
}

func (br *BountyRegistry) WatchForAssertions(aChan chan AssertionEvent) error {
	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics:    [][]common.Hash{{perigord.EventSignatureToTopicHash("NewAssertion(address,uint8,uint128,uint256)")}},
	}

	logChan := make(chan types.Log)
	sub, err := br.client.SubscribeFilterLogs(context.Background(), q, logChan)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(strings.NewReader(bindings.BountyRegistryABI))
	var abi abi.ABI
	if err := dec.Decode(&abi); err != nil {
		return err
	}

	go func() {
		log.Println("Started assertion event monitor")

		for {
			select {
			case logMsg := <-logChan:
				var nae NewAssertionEventLog
				if err := abi.Unpack(&nae, "NewAssertion", logMsg.Data); err != nil {
					log.Println("Ignoring event: ", err)
				} else {
					log.Println("New assertion", nae.Author.String(), nae.BountyGuid.String())

					assertStruct, err := br.session.AssertionsByGuid(nae.BountyGuid, nae.Index)
					if err != nil {
						log.Fatalln("Failed to get assertion based on event: ", err)
						break
					}

					na := Assertion(assertStruct)
					aChan <- AssertionEvent{
						Assertion:  &na,
						BountyGuid: nae.BountyGuid,
					}

					log.Println("Assertion", na.Author.String(), na.Verdict, na.BlockNumber.String())
				}
				break
			case err := <-sub.Err():
				log.Fatalln(err)
				return
			}
		}
	}()

	return nil
}

func (br *BountyRegistry) GetActiveBounties() []*Bounty {
	ret := make([]*Bounty, 0)
	for i := 0; ; i++ {
		bountyGuid, err := br.session.BountyGuids(big.NewInt(int64(i)))
		if err != nil {
			break
		}

		bountyStruct, err := br.session.BountiesByGuid(bountyGuid)
		if err != nil {
			break
		}

		nb := Bounty(bountyStruct)
		ret = append(ret, &nb)
	}

	return ret
}

func (br *BountyRegistry) GetBountyByGuid(guid *big.Int) *Bounty {
	bountyStruct, err := br.session.BountiesByGuid(guid)
	if err != nil {
		return nil
	}

	nb := Bounty(bountyStruct)
	return &nb
}
