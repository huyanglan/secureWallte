package eth

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func Transaction(fromAddress ,senderPrivateKey  ,toAddress  string, tranAmount ,_gasLimit ,_gasPrice int64) string  {
	chainId := big.NewInt(1) // main, 3 ropsten
	client, err := ethclient.Dial("http://40.81.25.93:8445")
	
	senderPrivKey, _ := crypto.HexToECDSA(senderPrivateKey)
	recipientAddr := common.HexToAddress(toAddress)

	fromAccDef := accounts.Account{
		Address: common.HexToAddress(fromAddress),
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAccDef.Address)

	// transfer amount
	amount := big.NewInt(tranAmount) // 0.1 ether
	// set gas
	gasLimit := uint64(_gasLimit)
	gasPrice := big.NewInt(_gasPrice) // 20 gwei
	//gasPrice, err := client.SuggestGasPrice(context.Background())

	tx := types.NewTransaction(nonce, recipientAddr, amount, gasLimit, gasPrice, nil)
	// sign
	signer := types.NewEIP155Signer(chainId)
	signedTx, _ := types.SignTx(tx, signer, senderPrivKey)

	var buff bytes.Buffer
	signedTx.EncodeRLP(&buff)
	// send transaction
	err = client.SendTransaction(context.Background(),signedTx)
	if err != nil {
		log.Println("EthNetWork-- SendTransactiono --err:",err)
	}
	// Get transaction hash
	hash :=  fmt.Sprintf("%s", signedTx.Hash().Hex())
	return hash
}


func TranferWithToken(_fromAddress, _senderPrivateKey, _toAddress, _tokenAddress, _tokenAmount  string, _amount ,_gasLimit ,_gasPrice int64) {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("FFE962244D80F95197089FE5FF87BE0163D485E7986A7070A498136012FD7B61")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	// fromAddress
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Panic(err)
	}

	value := big.NewInt(_amount)      // in wei (0 eth)
	gasLimit := uint64(_gasLimit) // in units
	gasPrice := big.NewInt(_gasPrice) // 20 gwei
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(_toAddress)
	tokenAddress := common.HexToAddress(_tokenAddress)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString(_tokenAmount, 10) // 1000000000000000000000 为1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)


	chainID, err := client.NetworkID(context.Background())
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	if err != nil {
		log.Fatal(err)
	}
	var buff bytes.Buffer
	signedTx.EncodeRLP(&buff)
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	// Get transaction hash
	fmt.Sprintf("%s", signedTx.Hash().Hex())
	//return hash
}


