package bounty

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/polyswarm/bindings"
)

type BountyPoster struct {
	session *bindings.BountyRegistrySession
	client  *ethclient.Client
}

func NewBountyPoster(session *bindings.BountyRegistrySession, client *ethclient.Client) *BountyPoster {
	session.TransactOpts.GasLimit = big.NewInt(1000000)
	return &BountyPoster{
		session: session,
		client:  client,
	}
}

func checkOutOfGas(tx *types.Transaction, rcpt *types.Receipt) bool {
	return tx.Gas().Cmp(rcpt.GasUsed) == 0
}

func waitMined(ctx context.Context, backend bind.DeployBackend, tx *types.Transaction) (*types.Receipt, error) {
	ctx, _ = context.WithTimeout(ctx, 60*time.Second)

	rcpt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return nil, err
	}
	if checkOutOfGas(tx, rcpt) {
		return nil, errors.New("out of gas")
	}

	return rcpt, nil
}

func eventSignatureToTopicHash(signature string) common.Hash {
	return crypto.Keccak256Hash([]byte(signature))
}

func (bp *BountyPoster) PostBounty(ctx context.Context, bnty *Bounty) (*types.Receipt, error) {
	tx, err := bp.session.RegisterBounty(bnty.BountyFee, bnty.BountyAmount, bnty.ArtifactHash, bnty.ArtifactURI, bnty.BlockDeadline, bnty.Guid)
	if err != nil {
		return nil, err
	}

	return waitMined(ctx, bp.client, tx)
}

func (bp *BountyPoster) PostAssertion(ctx context.Context, bnty *Bounty, astn *Assertion) (*types.Receipt, error) {
	astn.Author = bp.session.TransactOpts.From
	astn.SetGuid(bnty.Guid)
	tx, err := bp.session.RegisterAssertion(bnty.Originator, bnty.Guid, astn.Malicious, astn.AssertBid, astn.Metadata)
	if err != nil {
		return nil, err
	}

	return waitMined(ctx, bp.client, tx)
}

func (bp *BountyPoster) WatchForAssertions(aChan chan *Assertion) error {
	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics:    [][]common.Hash{{eventSignatureToTopicHash("NewAssertion(address,uint128,uint256)")}},
	}

	logChan := make(chan types.Log)
	sub, err := bp.client.SubscribeFilterLogs(context.Background(), q, logChan)
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
					assertStruct, err := bp.session.Assertions(nae.Guid, nae.Num)
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

func (bp *BountyPoster) WatchForBounties(bChan chan *Bounty) error {
	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics:    [][]common.Hash{{eventSignatureToTopicHash("NewBounty(address,uint256,uint256,uint256)")}},
	}

	logChan := make(chan types.Log)
	sub, err := bp.client.SubscribeFilterLogs(context.Background(), q, logChan)
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
					bountyStruct, err := bp.session.BountiesByAddress(nbe.Originator, nbe.Num)
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

func (bp *BountyPoster) GetActiveBounties() []*Bounty {
	bts := []*Bounty{}
	for i := 0; ; i++ {
		bountyStruct, err := bp.session.BountiesByAddress(bp.session.TransactOpts.From, big.NewInt(int64(i)))
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

func (bp *BountyPoster) GetBountyByGuid(guid *big.Int) *Bounty {
	activeBounties := bp.GetActiveBounties()

	for _, bnty := range activeBounties {
		if bnty.Guid.Cmp(guid) == 0 {
			return bnty
		}
	}

	return nil
}
