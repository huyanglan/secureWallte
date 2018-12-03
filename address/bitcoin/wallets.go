package bitcoin

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
const walletFile = "wallet_%s.dat"

type BitcoinWallets struct {
	BitcoinWallets map[string]*BitcoinWallet
}

func NewBitcoinWallets(nodeID string) (*BitcoinWallets, error) {
	wallets := BitcoinWallets{}
	wallets.BitcoinWallets = make(map[string]*BitcoinWallet)
	err := wallets.LoadFromFile(nodeID)
	return &wallets, err
}

// Load wallets from the file
func (ws *BitcoinWallets) LoadFromFile(nodeID string) error{
	walletFile := fmt.Sprintf(walletFile, nodeID)
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}
	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}
	var wallets BitcoinWallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	if err = decoder.Decode(&wallets); err != nil {
		log.Panic(err)
	}
	ws.BitcoinWallets = wallets.BitcoinWallets
	return nil
}

func (ws *BitcoinWallets) CreateWallet() string {
	wallet := NewBitcoinWalllet()
	address := fmt.Sprintf("%s", wallet.GetBitcoinAddress())
	return address
}

func (ws *BitcoinWallets) GetAddresses() []string{
	var addresses []string
	for address := range ws.BitcoinWallets{
		addresses = append(addresses, address)
	}
	return addresses
}

func NewWallets(nodeID string) (*BitcoinWallets, error) {
	wallets := BitcoinWallets{}
	wallets.BitcoinWallets = make(map[string]*BitcoinWallet)

	err := wallets.LoadFromFile(nodeID)
	return &wallets, err
}

func (ws *BitcoinWallets) SaveToFile(nodeID string) {
	var content bytes.Buffer
	walletFile := fmt.Sprintf(walletFile, nodeID)
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)
	if err := encoder.Encode(ws); err != nil {
		log.Panic(err)
	}
	if err := ioutil.WriteFile(walletFile, content.Bytes(), 0644); err != nil {
		log.Panic(err)
	}
}



