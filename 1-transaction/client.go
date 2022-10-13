package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	address string
	client  *ethclient.Client
}

// Address to set network endpoint
func (c *Client) Address(address string) *Client {
	c.address = address
	return c
}

// Connect to the client
func (c *Client) Connect() {
	client, err := ethclient.Dial(c.address)
	if err != nil {
		log.Fatal("Unable to connect to network!", err)
	}
	c.client = client

}

// Send funds from wallet to target with given amount
func (c *Client) Send(w *Wallet, to string, amount int64) string {
	nonce := c.generateNonce(w)
	gasFeeCap, _ := c.client.SuggestGasPrice(context.Background())
	gasTipCap, _ := c.client.SuggestGasTipCap(context.Background())

	tx := NewTransactionFrom(w).
		To(to).
		WithAmount(amount).
		Prepare(nonce, gasFeeCap, gasTipCap).
		Sign()

	return c.commit(tx)
}

func (c *Client) generateNonce(w *Wallet) uint64 {
	nonce, err := c.client.PendingNonceAt(context.Background(), w.GetAddress())
	if err != nil {
		log.Fatal("Unable to get nonce: ", err)
	}
	return nonce
}

// commit sends signed transaction and prints result
func (c *Client) commit(tx *Transaction) string {
	err := c.client.SendTransaction(context.Background(), tx.Get())
	if err != nil {
		log.Fatal("Unable to submit transaction: ", err)
	}

	return tx.Get().Hash().Hex()
}
