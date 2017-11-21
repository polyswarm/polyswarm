package migrations

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"
	"github.com/polyswarm/perigord/network"

	"github.com/polyswarm/polyswarm/bindings"
)

type BountyRegistryDeployer struct{}

func (d *BountyRegistryDeployer) Deploy(ctx context.Context, network *network.Network) (common.Address, *types.Transaction, interface{}, error) {
	account := network.Accounts()[0]
	network.Unlock(account, "blah")

	auth := network.NewTransactor(account)
	address, transaction, contract, err := bindings.DeployBountyRegistry(auth, network.Client())
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

			return nil
		},
	})
}
