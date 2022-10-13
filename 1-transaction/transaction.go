package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

type Transaction struct {
	wallet    *Wallet
	toAddress common.Address
	amount    *big.Int
	tx        *types.Transaction
}

func NewTransactionFrom(wallet *Wallet) *Transaction {
	return &Transaction{
		wallet: wallet,
	}
}

func (t *Transaction) To(to string) *Transaction {
	t.toAddress = common.HexToAddress(to)
	return t
}

func (t *Transaction) WithAmount(amount int64) *Transaction {
	t.amount = big.NewInt(amount)
	return t
}

func (t *Transaction) Prepare(nonce uint64, gasFeeCap, gasTipCap *big.Int) *Transaction {
	var data []byte

	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       uint64(100000),
		To:        &t.toAddress,
		Value:     t.amount,
		Data:      data,
	})

	t.tx = tx
	return t
}

func (t *Transaction) Sign() *Transaction {
	signer := types.LatestSignerForChainID(params.GoerliChainConfig.ChainID)
	signedTx, err := types.SignTx(t.tx, signer, t.wallet.privateKey)
	if err != nil {
		log.Fatal("Unable to sign tx: ", err)
	}
	t.tx = signedTx
	return t
}

func (t *Transaction) Get() *types.Transaction {
	return t.tx
}
