package eth

import (
	"log"
	"testing"
)

func TestGeneateAddress(t *testing.T) {
	privateKeyHex, publicKeyHex, address  := GenerateAddress()
	log.Println("---privateKeyHex, publicKeyHex, address----", privateKeyHex, publicKeyHex, address)
}
