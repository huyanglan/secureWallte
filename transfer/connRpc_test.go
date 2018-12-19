package transfer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math"
	"math/big"
	"strings"
	"testing"
)

func TestConnRpc(t *testing.T) {
	// Create test account
	ks := keystore.NewKeyStore("/", keystore.StandardScryptN, keystore.StandardScryptP)
	address, _ := ks.NewAccount("passwd")
	// get private key
	privateKey, err := ks.Export(address, "passwd", "passwd")
	if err != nil {
		panic(err)
	}
	ConnRpc()
	fmt.Println("Address", address.Address.Hex())
	fmt.Println("account", privateKey)
}

func TestTransfer(t *testing.T) {
	rpcDial, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(rpcDial)
	// import account and key
	auth, err := bind.NewTransactor(strings.NewReader("json"), "password")
    // address is contract address
	token, err := NewToken(common.HexToAddress("0xcedd33c310a0422df87e6cb4817da3f30a793565"), client)
	if err != nil {
		panic(err)
	}
	// token bit
	decimal, err := token.Decimals(nil)
	if err != nil {
		panic(err)
	}

	// deal code length
	tenDecimal := big.NewFloat(math.Pow(10,float64(decimal)))
	fmt.Println("tenDecimal", tenDecimal, auth)
	//var amount float
	//convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})
	//toAddress := ""
	//txs, err := token.Transfer(auth, common.HexToAddress(toAddress), convertAmount)
	//fmt.Println("Transfer", txs)
}


