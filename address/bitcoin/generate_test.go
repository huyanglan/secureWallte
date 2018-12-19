package bitcoin

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
)

func TestGenerate(t *testing.T) {
	wallet := NewBitcoinWalllet()
	address := wallet.GetBitcoinAddress()
	isVal := ValidateAddress(string(address))
	fmt.Println(isVal, "address", string(address))
}
// 测试发送交易
func TestTransaction(t *testing.T) {
	chainId := big.NewInt(3) // ropsten
	client, err := ethclient.Dial("https://ropsten.infura.io")
	fromaddress :="0xc548c20F02A7fd669c86C881A0c243fdbD51ACd5"
	senderPrivKey, _ := crypto.HexToECDSA("3D53A0EE623CE16DEB9A7BD8337255E7EC21AEE300F7CD6A11E7FF92AFC3B2E6")
	recipientAddr := common.HexToAddress("0xFe56293Cd703E8207e89e68f91fD8aC7dE7Dff4b")

	//获取nonce值
	fromAccDef := accounts.Account{
		Address: common.HexToAddress(fromaddress),
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAccDef.Address)
	fmt.Print("nonce---- :",nonce)

	amount := big.NewInt(100000000000000000) // 0.1 ether
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(20000000000) // 20 gwei

	tx := types.NewTransaction(nonce, recipientAddr, amount, gasLimit, gasPrice, nil)

	signer := types.NewEIP155Signer(chainId)
	signedTx, _ := types.SignTx(tx, signer, senderPrivKey)


	var buff bytes.Buffer
	signedTx.EncodeRLP(&buff)
	fmt.Printf("0x%x\n", buff.Bytes())


	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Transaction error :",err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
