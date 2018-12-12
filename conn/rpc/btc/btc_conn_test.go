package btc

import (
	"fmt"
	"github.com/btcsuite/rpcclient"

	"log"
	"testing"
)

func TestBtcConn(t *testing.T) {
	ConnPeer()
}

func  TestHttpConn(t *testing.T) {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "40.81.25.93:8332",
		User:         "bitcoinrpc",
		Pass:         "123456",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	count, err := client.GetDifficulty()
	fmt.Println(count)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Block count: %d", blockCount)
}


