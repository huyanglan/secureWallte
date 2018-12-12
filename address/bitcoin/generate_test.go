package bitcoin

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	wallet := NewBitcoinWalllet()
	address := wallet.GetBitcoinAddress()
	isVal := ValidateAddress(string(address))
	fmt.Println(isVal, "address", string(address))
}
