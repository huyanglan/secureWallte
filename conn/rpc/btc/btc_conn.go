package btc
import (
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/rpcclient"
	"log"
	"time"
)

const (
	defaultServerUrl string = "http://40.81.25.93:8333"
	defaultUserName string = "bitcoinrpc"
	defalutPassword string = "123456"
)

// Connect to full peer
func ConnPeer() {
	// most of these handlers will only be called if you register for notifications
	ntfnHandlers := rpcclient.NotificationHandlers{
		OnFilteredBlockConnected: func(height int32, header *wire.BlockHeader, txns []*btcutil.Tx) {
			log.Printf("Block connected: %v (%d) %v",
				header.BlockHash(), height, header.Timestamp)
		},
		OnFilteredBlockDisconnected: func(height int32, header *wire.BlockHeader) {
			log.Printf("Block disconnected: %v (%d) %v",
				header.BlockHash(), height, header.Timestamp)
		},
	}
	// Connect to local btcd RPC server using websockets.
	//btcdHomeDir := btcutil.AppDataDir("btcd", false)
	//certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	connCfg := &rpcclient.ConnConfig{
		Host:         defaultServerUrl,
		Endpoint:     "ws",
		User:         defaultUserName,
		Pass:         defalutPassword,
	}
	client, err := rpcclient.New(connCfg, &ntfnHandlers)
	if err != nil {
		log.Fatal(err)
	}

	// Register for block connect and disconnect notifications.
	if err := client.NotifyBlocks(); err != nil {
		log.Fatal(err)
	}
	log.Println("NotifyBlocks: Registration Complete")

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)

	// For this example gracefully shutdown the client after 10 seconds.
	// Ordinarily when to shutdown the client is highly application
	// specific.
	log.Println("Client shutdown in 10 seconds...")
	time.AfterFunc(time.Second*10, func() {
		log.Println("Client shutting down...")
		client.Shutdown()
		log.Println("Client shutdown complete.")
	})

	// Wait until the client either shuts down gracefully (or the user
	// terminates the process with Ctrl+C).
	client.WaitForShutdown()
}

func RpcCall() {

}

