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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func pollTransactionUntilTime(conn *ethclient.Client, txn *types.Transaction, d time.Duration) (*types.Receipt, error) {
	sliceSz := time.Millisecond * 200
	slices := d / (sliceSz)

	for i := 1; i < int(slices.Nanoseconds()); i++ {
		ctx := context.Background()
		receipt, err := conn.TransactionReceipt(ctx, txn.Hash())
		if err != nil && err != ethereum.NotFound {
			return nil, err

		}

		if err == nil {
			if txn.Gas().Cmp(receipt.GasUsed) == 0 {
				return nil, errors.New("out of gas")
			}

			return receipt, err
		}

		time.Sleep(sliceSz)
	}

	return nil, errors.New("Failed to get txn before timeout")
}

func (bp *BountyPoster) checkTxnUntilReceipt(txn *types.Transaction, logHeading string) (*types.Receipt, error) {
	receipt, err := pollTransactionUntilTime(bp.client, txn, 300*time.Second)
	if err != nil {
		return nil, err
	}

	h := ""
	for _, l := range receipt.Logs {
		for _, t := range l.Topics {
			h += t.String() + ", "
		}

	}

	log.Println(logHeading, "events will appear under following topic hashes: ", h)
	return receipt, nil
}

func (bp *BountyPoster) PostBounty(bnty *Bounty) (*types.Receipt, error) {
	txn, err := bp.session.RegisterBounty(bnty.BountyFee, bnty.BountyAmount, bnty.ArtifactHash, bnty.ArtifactURI, bnty.BlockDeadline, bnty.Guid)
	if err != nil {
		return nil, err
	}

	return bp.checkTxnUntilReceipt(txn, "bounty")
}

func (bp *BountyPoster) PostAssertion(bnty *Bounty, astn *Assertion) (*types.Receipt, error) {
	astn.Author = bp.session.TransactOpts.From
	astn.SetGuid(bnty.Guid)
	txn, err := bp.session.RegisterAssertion(bnty.Originator, bnty.Guid, astn.Malicious, astn.AssertBid, astn.Metadata)
	if err != nil {
		return nil, err

	}

	return bp.checkTxnUntilReceipt(txn, "assertion")
}

func (bp *BountyPoster) WatchForAssertions(aChan chan *Assertion) error {
	assertH := common.HexToHash("0x0b496f3ca4b1224bf741d2ce2e3657c9df30bdb6c2e02f598ad531e786f06f93")

	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics:    [][]common.Hash{{assertH}},
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
				type ex struct {
					Author common.Address
					Guid   *big.Int
					Num    *big.Int
				}

				var lm ex
				log.Println(logMsg.Topics[0].String())

				if err := abi.Unpack(&lm, "NewAssertion", logMsg.Data); err != nil {
					log.Println(err, "ignoring event")
				} else {
					log.Println("new assertion", lm.Author.String(), lm.Guid.String(), lm.Num.String())
					assertStruct, err := bp.session.Assertions(lm.Guid, lm.Num)
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
					na.Guid = lm.Guid

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
	// todo this is jacked from log message on bounty submit... kinda lame not sure how its generated ... if the signature si changed we'll ahve to deal
	bountyH := common.HexToHash("0x1ba76ebd15dae915e57648382e814958cb98e8f8a5d42e17b9df31d235c7a7b5")

	q := ethereum.FilterQuery{
		Addresses: []common.Address{contract.AddressOf("BountyRegistry")},
		Topics:    [][]common.Hash{{bountyH}},
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
				type ex struct {
					Originator common.Address
					Num        *big.Int
					Amount     *big.Int
					Deadline   *big.Int
				}

				var lm ex
				for _, t := range logMsg.Topics {
					log.Println("thash", t.String(), bountyH.String())
				}

				if err := abi.Unpack(&lm, "NewBounty", logMsg.Data); err != nil {
					log.Println(err, "ignoring event non-bounty post")

				} else {
					log.Println("new bounty amount addr", lm.Amount.String(), lm.Originator.String())
					bountyStruct, err := bp.session.Bounties(lm.Originator, lm.Num)

					// todo this is lame, done because we don't have type generated in registry solidity interface
					nb := new(Bounty)

					nb.Originator = bountyStruct.Originator
					nb.ArtifactHash = bountyStruct.ArtifactHash
					nb.ArtifactURI = bountyStruct.ArtifactURI
					nb.Guid = bountyStruct.Guid
					nb.BlockDeadline = bountyStruct.BlockDeadline
					nb.BountyAmount = bountyStruct.BountyAmount
					nb.BountyFee = bountyStruct.BountyFee

					log.Println("bounty struct: ", nb)

					bChan <- nb
					if err != nil {
						log.Fatalln("Failed to get bounty based on event: ", err)
					}
					log.Println("Bounty URI", bountyStruct.ArtifactURI, bountyStruct.ArtifactHash)
				}
				break

				log.Println(logMsg.String())
			case err := <-sub.Err():
				log.Fatalln(err)
				return

			}

		}
	}()
	return nil
}

func (bp *BountyPoster) GetActiveBounties() ([]Bounty, error) {
	bts := []Bounty{}
	for i := 0; ; i++ {
		bountyStruct, err := bp.session.Bounties(bp.session.TransactOpts.From, big.NewInt(int64(i)))
		if err != nil {
			break
		}

		b := Bounty{}
		b.Guid = bountyStruct.Guid
		b.ArtifactURI = bountyStruct.ArtifactURI
		b.ArtifactHash = bountyStruct.ArtifactHash
		// todo what else?
		bts = append(bts, b)
	}

	return bts, nil
}
