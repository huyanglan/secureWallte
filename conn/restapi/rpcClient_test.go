package restapi

import (
	"fmt"
	"testing"

)

func TestRPC(t *testing.T) {
	client, err := newClient("40.81.25.93", 8332, "bitcoinrpc", "123456", false)
	if err != nil {
		panic(err)
	}
	param := make([]int, 1)
	rsp, err := client.call("getblockcount", param)
	fmt.Println("sss", rsp)
}
