package actions

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"polygon-token-oracle-backend/logic/blockchain/eth"
	"polygon-token-oracle-backend/logic/blockchain/eth/contracts/polygon_token_handle"
	"polygon-token-oracle-backend/logic/config"
)

func PushPolygonTokenBalance(target common.Address, requestId [32]byte, balance uint64) (string, error) {
	polygonTokenHandleContract, err := polygon_token_handle.NewPolygonTokenHandle(target, eth.DefaultClient.Client())
	if err != nil {
		return "", err
	}

	opts, err := eth.DefaultClient.GetDefaultTransactionOptions(context.Background(), config.Get().Wallet.PrivateKey)
	transaction, err := polygonTokenHandleContract.HandleCallback(opts, requestId, new(big.Int).SetUint64(balance))
	if err != nil {
		return "", err
	}

	return transaction.Hash().String(), nil
}
