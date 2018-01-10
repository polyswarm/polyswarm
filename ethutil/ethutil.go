package ethutil

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func CheckOutOfGas(tx *types.Transaction, rcpt *types.Receipt) bool {
	return tx.Gas().Cmp(rcpt.GasUsed) == 0
}

func WaitMined(ctx context.Context, backend bind.DeployBackend, tx *types.Transaction) (*types.Receipt, error) {
	ctx, _ = context.WithTimeout(ctx, 60*time.Second)

	rcpt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return nil, err
	}
	if CheckOutOfGas(tx, rcpt) {
		return nil, errors.New("out of gas")
	}

	return rcpt, nil
}

func EventSignatureToTopicHash(signature string) common.Hash {
	return crypto.Keccak256Hash([]byte(signature))
}
