package bountyregistry

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/polyswarm/bindings"
	"github.com/polyswarm/polyswarm/ethutil"
)

type BountyRegistry struct {
	session *bindings.BountyRegistrySession
	client  *ethclient.Client
}

func NewBountyRegistry(session *bindings.BountyRegistrySession, client *ethclient.Client) *BountyRegistry {
	session.TransactOpts.GasLimit = big.NewInt(1000000)
	return &BountyRegistry{
		session: session,
		client:  client,
	}
}

func (br *BountyRegistry) PostBounty(ctx context.Context, bnty *Bounty) (*types.Receipt, error) {
	tx, err := br.session.RegisterBounty(bnty.BountyFee, bnty.BountyAmount, bnty.ArtifactHash, bnty.ArtifactURI, bnty.BlockDeadline, bnty.Guid)
	if err != nil {
		return nil, err
	}

	return ethutil.WaitMined(ctx, br.client, tx)
}

func (br *BountyRegistry) PostAssertion(ctx context.Context, bnty *Bounty, astn *Assertion) (*types.Receipt, error) {
	astn.Author = br.session.TransactOpts.From
	astn.SetGuid(bnty.Guid)
	tx, err := br.session.RegisterAssertion(bnty.Originator, bnty.Guid, astn.Malicious, astn.AssertBid, astn.Metadata)
	if err != nil {
		return nil, err
	}

	return ethutil.WaitMined(ctx, br.client, tx)
}

func (br *BountyRegistry) WatchForBounties(bChan chan *Bounty) error {
	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics:    [][]common.Hash{{ethutil.EventSignatureToTopicHash("NewBounty(address,uint256,uint256,uint256)")}},
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
				type NewBountyEvent struct {
					Originator common.Address
					Num        *big.Int
					Amount     *big.Int
					Deadline   *big.Int
				}

				var nbe NewBountyEvent
				if err := abi.Unpack(&nbe, "NewBounty", logMsg.Data); err != nil {
					log.Println("ignoring event non-bounty post: ", err)
				} else {
					log.Println("new bounty amount addr", nbe.Amount.String(), nbe.Originator.String())
					bountyStruct, err := br.session.BountiesByAddress(nbe.Originator, nbe.Num)
					if err != nil {
						log.Fatalln("Failed to get bounty based on event: ", err)
					}

					// todo this is lame, done because we don't have type generated in registry solidity interface
					nb := new(Bounty)

					nb.Originator = bountyStruct.Originator
					nb.ArtifactHash = bountyStruct.ArtifactHash
					nb.ArtifactURI = bountyStruct.ArtifactURI
					nb.Guid = bountyStruct.Guid
					nb.BlockDeadline = bountyStruct.BlockDeadline
					nb.BountyAmount = bountyStruct.BountyAmount
					nb.BountyFee = bountyStruct.BountyFee

					bChan <- nb
					log.Println("Bounty URI", bountyStruct.ArtifactURI, bountyStruct.ArtifactHash)
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

func (br *BountyRegistry) WatchForAssertions(aChan chan *Assertion) error {
	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics:    [][]common.Hash{{ethutil.EventSignatureToTopicHash("NewAssertion(address,uint128,uint256)")}},
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
				type NewAssertionEvent struct {
					Author common.Address
					Guid   *big.Int
					Num    *big.Int
				}

				var nae NewAssertionEvent
				if err := abi.Unpack(&nae, "NewAssertion", logMsg.Data); err != nil {
					log.Println("ignoring event: ", err)
				} else {
					log.Println("new assertion", nae.Author.String(), nae.Guid.String(), nae.Num.String())
					assertStruct, err := br.session.Assertions(nae.Guid, nae.Num)
					if err != nil {
						log.Fatalln("Failed to get assertion based on event: ", err)
						break

					}

					// todo this is lame, done because we don't have type generated in registry solidity interface
					na := new(Assertion)

					na.Author = assertStruct.Author
					na.Malicious = assertStruct.Malicious
					na.BlockTime = assertStruct.BlockTime
					na.AssertBid = assertStruct.AssertBid
					na.Metadata = assertStruct.Metadata

					// todo this is not part of solidity dstruct, but must be added...
					na.Guid = nae.Guid

					aChan <- na
					log.Println("Assertion", na.Author.String(), na.Malicious, na.BlockTime.String())
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
	bts := []*Bounty{}
	for i := 0; ; i++ {
		bountyStruct, err := br.session.BountiesByAddress(br.session.TransactOpts.From, big.NewInt(int64(i)))
		if err != nil {
			break
		}

		b := Bounty{}
		b.Guid = bountyStruct.Guid
		b.ArtifactURI = bountyStruct.ArtifactURI
		b.ArtifactHash = bountyStruct.ArtifactHash
		// TODO what else?
		bts = append(bts, &b)
	}

	return bts
}

func (br *BountyRegistry) GetBountyByGuid(guid *big.Int) *Bounty {
	activeBounties := br.GetActiveBounties()

	for _, bnty := range activeBounties {
		if bnty.Guid.Cmp(guid) == 0 {
			return bnty
		}
	}

	return nil
}
