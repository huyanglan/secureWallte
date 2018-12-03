package transfer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)


func ConnRpc()  {
	rpcDial, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	client := ethclient.NewClient(rpcDial)
	fmt.Println("client", client)
}

func GetBalance(address string) {
	client, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println("rpc Dial err", err)
	}
	var account[] string
	err = client.Call(&account, "eth_accounts")
	var result string
	if err = client.Call(&result, "eth_getBalance", account[0], "latest"); err != nil {
		panic(err)
	}
	fmt.Println("account[0]:  balance[0]: ", account[0], result)
}