package migrations

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/polyswarm/perigord/contract"
	"github.com/polyswarm/perigord/migration"

	"github.com/polyswarm/polyswarm/bindings"
)

type NectarTokenDeployer struct{}

func (d *NectarTokenDeployer) Deploy(ctx context.Context, network *migration.Network) (common.Address, *types.Transaction, interface{}, error) {
	auth := network.NewTransactor(0)
	address, transaction, contract, err := bindings.DeployNectarToken(auth, network.Backend(), big.NewInt(1000000), "Nectar Token", 18, "NCT")
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	session := &bindings.NectarTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *NectarTokenDeployer) Bind(ctx context.Context, network *migration.Network, address common.Address) (interface{}, error) {
	auth := network.NewTransactor(0)
	contract, err := bindings.NewNectarToken(address, network.Backend())
	if err != nil {
		return nil, err
	}

	session := &bindings.NectarTokenSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("NectarToken", &NectarTokenDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 3,
		F: func(ctx context.Context, network *migration.Network) error {
			if err := contract.Deploy(ctx, "NectarToken", network); err != nil {
				return err
			}

			return nil
		},
	})
}
