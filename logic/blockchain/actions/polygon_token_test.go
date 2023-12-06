package actions

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestPushPolygonTokenBalance(t *testing.T) {
	requestId := [32]byte{144, 125, 81, 106, 158, 94, 205, 160, 141, 151, 184, 141, 122, 142, 161, 33, 194, 25, 139, 118, 113, 221, 147, 82, 187, 0, 198, 149, 194, 59, 184, 9}
	tx, err := PushPolygonTokenBalance(common.HexToAddress("0x7D3F01b896505350920deFaD8cb4162ecB93770D"), requestId, 100)
	if err != nil {
		panic(err)
	}

	fmt.Printf("tx: %s\n", tx)
}
