package migrations

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"

	"github.com/polyswarm/polyswarm/bindings"
)

type MigrationsDeployer struct{}

func (d *MigrationsDeployer) Deploy(ctx context.Context, network *migration.Network) (common.Address, *types.Transaction, interface{}, error) {
    auth := network.NewTransactor(0)
	address, transaction, contract, err := bindings.DeployMigrations(auth, network.Backend())
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	session := &bindings.MigrationsSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *MigrationsDeployer) Bind(ctx context.Context, network *migration.Network, address common.Address) (interface{}, error) {
    auth := network.NewTransactor(0)
	contract, err := bindings.NewMigrations(address, network.Backend())
	if err != nil {
		return nil, err
	}

	session := &bindings.MigrationsSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("Migrations", &MigrationsDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 1,
		F: func(ctx context.Context, network *migration.Network) error {
			if err := contract.Deploy(ctx, "Migrations", network); err != nil {
				return err
			}

			return nil
		},
	})
}
