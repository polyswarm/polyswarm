package migrations

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/polyswarm/perigord"
	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"

	"github.com/polyswarm/polyswarm/bindings"
)

func BuyNectar(address common.Address, network *network.Network) error {
	nectarTokenSession, ok := contract.Session("NectarToken").(*bindings.NectarTokenSession)
	if nectarTokenSession == nil || !ok {
		return errors.New("Invalid NectarToken session")
	}

	nectarTokenSession.TransactOpts.GasLimit = 1000000
	tx, err := nectarTokenSession.Mint(address, big.NewInt(1000000000000000000))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if _, err := perigord.WaitMined(ctx, network.Client(), tx); err != nil {
		return err
	}

	return nil
}

func ApproveRegistry(address common.Address, network *network.Network) error {
	nectarTokenSession, ok := contract.Session("NectarToken").(*bindings.NectarTokenSession)
	if nectarTokenSession == nil || !ok {
		return errors.New("Invalid NectarToken session")
	}

	nectarTokenSession.TransactOpts.GasLimit = 1000000
	//tx, err := nectarTokenSession.Approve(address, big.NewInt(1000000000000000000), []byte{})
	tx, err := nectarTokenSession.Approve(address, big.NewInt(1000000000000000000))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if _, err := perigord.WaitMined(ctx, network.Client(), tx); err != nil {
		return err
	}

	return nil
}

type BountyRegistryDeployer struct{}

func (d *BountyRegistryDeployer) Deploy(ctx context.Context, network *network.Network) (common.Address, *types.Transaction, interface{}, error) {
	account := network.Accounts()[0]
	network.Unlock(account, "blah")

	auth := network.NewTransactor(account)

	address, transaction, contract, err := bindings.DeployBountyRegistry(auth, network.Client(), contract.AddressOf("NectarToken"))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	session := &bindings.BountyRegistrySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *BountyRegistryDeployer) Bind(ctx context.Context, network *network.Network, address common.Address) (interface{}, error) {
	account := network.Accounts()[0]
	network.Unlock(account, "blah")

	auth := network.NewTransactor(account)
	contract, err := bindings.NewBountyRegistry(address, network.Client())
	if err != nil {
		return nil, err
	}

	session := &bindings.BountyRegistrySession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("BountyRegistry", &BountyRegistryDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 3,
		F: func(ctx context.Context, network *network.Network) error {
			if err := contract.Deploy(ctx, "BountyRegistry", network); err != nil {
				return err
			}

			account := network.Accounts()[0]
			network.Unlock(account, "blah")

			auth := network.NewTransactor(account)

			if err := BuyNectar(auth.From, network); err != nil {
				return err
			}

			if err := ApproveRegistry(contract.AddressOf("BountyRegistry"), network); err != nil {
				return err
			}

			return nil
		},
	})
}
