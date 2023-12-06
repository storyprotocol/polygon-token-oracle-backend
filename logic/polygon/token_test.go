package polygon

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestGetTokenBalance(t *testing.T) {
	balance, err := GetTokenBalance(common.HexToAddress("0xEB449Df869d38e61AdDd3e60544a1B600ad93dfA"), common.HexToAddress("0xf398C12A45Bc409b6C652E25bb0a3e702492A4ab"))
	if err != nil {
		return
	}

	fmt.Println(balance)
}
