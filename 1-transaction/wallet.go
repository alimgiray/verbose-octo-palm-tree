package main

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewWallet(key string) *Wallet {
	wallet := &Wallet{}
	wallet.setPrivateKey(key)
	wallet.setPublicKey()
	return wallet
}

func (w *Wallet) setPrivateKey(key string) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal("Unable to parse private key: ", err)
	}

	w.privateKey = privateKey
}

func (w *Wallet) setPublicKey() {
	publicKey := w.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("Unable to cast public key to ECDSA")
	}

	w.publicKey = publicKeyECDSA
}

func (w *Wallet) GetAddress() common.Address {
	return crypto.PubkeyToAddress(*w.publicKey)
}
