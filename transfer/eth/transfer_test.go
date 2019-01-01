package eth

import "testing"

func TestTransaction(t *testing.T) {

}

func TestTranferWithToken(t *testing.T){
	fromAddress := ""
	senderPrivateKey := ""
	toAddress := "0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d"
	tokenAddress := "0x28b149020d2152179873ec60bed6bf7cd705775d"
	tokenAmount := "1000000000000000000000"
	var amount int64 = 5
	var gasLimit int64 = 210000
	var gasPrice int64 = 20
	TranferWithToken(fromAddress, senderPrivateKey, toAddress, tokenAddress, tokenAmount, amount ,gasLimit ,gasPrice)
}


