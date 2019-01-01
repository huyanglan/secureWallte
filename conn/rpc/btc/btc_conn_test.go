package btc

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/rpcclient"
	"math/rand"
	"strings"
	"time"

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
	count, err := client.GetBestBlockHash()
	fmt.Println(count)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Block count: %d", blockCount)
}

func TestInterface(t *testing.T) {

		fmt.Println("------------",RandStringRunes(20))
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return strings.ToLower(string(b))
}

func getInterface() *interface{}{

	return nil
}
type T int64
func TestTypeToOri(t *testing.T) {
	//var n T = 1
	//n = 5
	s1 := GetMD5Hash("12345678")
	log.Println("-----", s1)
}

func GetMD5Hash(_text string) string {
	hasher := md5.New()
	hasher.Write([]byte(_text))
	return hex.EncodeToString(hasher.Sum(nil))
}

