package bitcoin

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	wallet := NewBitcoinWalllet()
	address := wallet.GetBitcoinAddress()
	fmt.Println("address", address)
}
