package polygon

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"polygon-token-oracle-backend/logic/blockchain/eth/contracts/erc20_721"
)

func GetTokenBalance(tokenAddress, tokenOwnerAddress common.Address) (uint64, error) {
	polygonTokenContract, err := erc20_721.NewErc20721(tokenAddress, DefaultClient.Client())
	if err != nil {
		return 0, err
	}

	res, err := polygonTokenContract.BalanceOf(&bind.CallOpts{}, tokenOwnerAddress)
	if err != nil {
		return 0, err
	}

	return res.Uint64(), nil
}
