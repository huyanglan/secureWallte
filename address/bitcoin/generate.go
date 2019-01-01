package bitcoin

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)
const version = byte(0x00) // use to generate address version
const addressChecksumLen = 4   // use to generate address Checksum bit
type BitcoinWallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// batch generate address
func BatchGenerateAddress() {

}

func NewBitcoinWalllet() *BitcoinWallet {
	privateKey, publicKey := newKeyPair()
	log.Println("-------Key----", privateKey, "-------",publicKey)
	wallet := BitcoinWallet{privateKey, publicKey}
	return &wallet
}

func (w BitcoinWallet) GetBitcoinAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)
	// joint version
	versionPayload := append([]byte{version}, pubKeyHash...)
	checkSum := checksum(versionPayload)
	fullPayload := append(versionPayload, checkSum...)
	// base58 code
	return Base58Encode(fullPayload)
}

func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1:len(pubKeyHash) - addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))
	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

// twice sha256 hash to generate checkSum
func checksum(payload []byte) []byte{
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressChecksumLen]
}

// deal public key with RIPEMD160
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	if _, err := RIPEMD160Hasher.Write(publicSHA256[:]); err != nil {
		log.Panic(err)
	}

	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

// generate public key and private key
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	// ECC generate private key
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	log.Println("--------", private)
	if err != nil {
		log.Panic(err)
	}
	// private key generate public key
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pubKey
}

