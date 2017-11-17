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
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/polyswarm/polyswarm/bindings"
)

const DEFAULT_REG_ADDR = "deadbeef" // todo

type BountyPoster struct {
	keyJson       string
	keyPassphrase string

	bountyContractAddr common.Address
	cli                *ethclient.Client

	bountyReg *bindings.BountyRegistry
}

func NewBountyPoster(cli *ethclient.Client, kJson string, pass string) (*BountyPoster, error) {
	b := BountyPoster{keyJson: kJson, keyPassphrase: pass}
	if err := b.SetRegistry(common.HexToAddress(DEFAULT_REG_ADDR)); err != nil {
		return nil, err
	}
	b.cli = cli
	return &b, nil
}

func (bp *BountyPoster) SetRegistry(addr common.Address) error {
	bp.bountyContractAddr = addr
	br, err := bindings.NewBountyRegistry(bp.bountyContractAddr, bp.cli)

	if err != nil {
		return err
	}

	bp.bountyReg = br

	return nil
}

func (bp *BountyPoster) checkTxnUntilReceipt(ctx context.Context, txn *types.Transaction, logHeading string) (*types.Receipt, error) {
	tmout, _ := context.WithTimeout(ctx, time.Second*30)
	rcpt, txnE := bp.cli.TransactionReceipt(tmout, txn.Hash())
	if txnE != nil {
		return nil, txnE
	}

	if txn.Gas().Cmp(rcpt.GasUsed) == 0 {
		return nil, errors.New("out of gas")
	}
	h := ""

	for _, l := range rcpt.Logs {
		for _, t := range l.Topics {
			h += t.String() + ", "
		}

	}
	log.Println(logHeading, "events will appear under following topic hashes: ", h)
	return rcpt, nil
}

func (bp *BountyPoster) PostBounty(ctx context.Context, bnty *Bounty) (*types.Receipt, error) {
	auth, err := bind.NewTransactor(strings.NewReader(bp.keyJson), bp.keyPassphrase)
	if err != nil {
		return nil, err
	}

	txn, exnE := bp.bountyReg.RegisterBounty(auth, bnty.BountyFee, bnty.BountyAmount, bnty.ArtifactHash, bnty.ArtifactURI, bnty.BlockDeadline, bnty.Guid)

	if exnE != nil {
		return nil, exnE
	}

	return bp.checkTxnUntilReceipt(ctx, txn, "bounty")

}

func (bp *BountyPoster) PostAssertion(ctx context.Context, bnty *Bounty, astn *Assertion) (*types.Receipt, error) {
	auth, err := bind.NewTransactor(strings.NewReader(bp.keyJson), bp.keyPassphrase)
	if err != nil {
		return nil, err
	}

	astn.Author = getOriginAccountJson(bp.keyJson)
	astn.SetGuid(bnty.Guid)
	txn, exnE := bp.bountyReg.RegisterAssertion(auth, bnty.Originator, bnty.Guid, astn.Malicious, astn.AssertBid, astn.Metadata)

	if exnE != nil {
		return nil, exnE

	}

	return bp.checkTxnUntilReceipt(ctx, txn, "assertion")

}

func getOriginAccountJson(js string) common.Address {
	dec := json.NewDecoder(strings.NewReader(js))

	type tmp struct {
		Address string `json:address`
	}
	t := tmp{}

	e := dec.Decode(&t)
	if e != nil {
		panic(e)
	}

	return common.HexToAddress(t.Address)
}

func (bp *BountyPoster) WatchForAssertions(ctx context.Context, aChan chan *Assertion) error {
	assertH := common.HexToHash("0x0b496f3ca4b1224bf741d2ce2e3657c9df30bdb6c2e02f598ad531e786f06f93")

	q := ethereum.FilterQuery{

		Addresses: []common.Address{bp.bountyContractAddr},
		Topics:    [][]common.Hash{{assertH}},
	}

	logChan := make(chan types.Log)
	sub, err := bp.cli.SubscribeFilterLogs(ctx, q, logChan)
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

				if decE := abi.Unpack(&lm, "NewAssertion", logMsg.Data); decE != nil {
					log.Println(decE, "ignoring event")
				} else {
					log.Println("new assertion", lm.Author.String(), lm.Guid.String(), lm.Num.String())
					assertStruct, err := bp.bountyReg.Assertions(nil, lm.Guid, lm.Num)

					if err != nil {
						log.Fatalln("Failed to get assertion based on event")
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

func (bp *BountyPoster) WatchForBounties(ctx context.Context, bChan chan *Bounty) error {
	// todo this is jacked from log message on bounty submit... kinda lame not sure how its generated ... if the signature si changed we'll ahve to deal
	bountyH := common.HexToHash("0x1ba76ebd15dae915e57648382e814958cb98e8f8a5d42e17b9df31d235c7a7b5")

	q := ethereum.FilterQuery{

		Addresses: []common.Address{bp.bountyContractAddr},
		Topics:    [][]common.Hash{{bountyH}},
	}

	logChan := make(chan types.Log)
	sub, err := bp.cli.SubscribeFilterLogs(ctx, q, logChan)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(strings.NewReader(bindings.BountyRegistryABI))
	var abi abi.ABI
	if err := dec.Decode(&abi); err != nil {
		return err
	}
	/*
		for _, arg := range abi.Events["NewBounty"].Inputs{
			log.Println(arg.Name, arg.Type)
		}*/

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

				if decE := abi.Unpack(&lm, "NewBounty", logMsg.Data); decE != nil {
					log.Println(decE, "ignoring event non-bounty post")

				} else {
					log.Println("new bounty amount addr", lm.Amount.String(), lm.Originator.String())
					bountyStruct, err := bp.bountyReg.Bounties(nil, lm.Originator, lm.Num)

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
					if err != nil {
						log.Fatalln("Failed to get bounty based on event")
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
		bountyStruct, err := getBounty(bp, i)
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
func getBounty(bp *BountyPoster, i int) (struct {
	Originator    common.Address
	BountyFee     *big.Int
	BountyAmount  *big.Int
	ArtifactHash  string
	ArtifactURI   string
	BlockDeadline *big.Int
	Guid          *big.Int
}, error) {
	bountyStruct, err := bp.bountyReg.Bounties(nil, getOriginAccountJson(bp.keyJson), big.NewInt(int64(i)))
	return bountyStruct, err
}
