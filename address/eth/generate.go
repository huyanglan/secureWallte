package eth

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"regexp"
)

func GenerateAddress() (string, string, string){
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//  get private key of string   formate
	privateKeyHex := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	publicKeyHex := hexutil.Encode(publicKeyBytes)[4:]

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	if !VaildateAddress(address) {
		log.Panic("---address is invalidate----", address)
	}
	return privateKeyHex, publicKeyHex, address
}

func VaildateAddress(address string) bool{
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}