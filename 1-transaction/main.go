package main

import (
	"flag"
	"fmt"
)

var from, to string
var amount int64

func init() {
	flag.StringVar(&from, "from", "", "Your private key")
	flag.StringVar(&to, "to", "0xA5058fbcD09425e922E3E9e78D569aB84EdB88Eb", "Receiver public key")
	flag.Int64Var(&amount, "amount", 1000, "Amount to send")
	flag.Parse()
}

func main() {
	client := Client{}
	client.Address("https://goerli.infura.io/v3/a26b774954414d04a82d6329c583d6f2").Connect()

	wallet := NewWallet(from)

	hash := client.Send(wallet, to, amount)
	fmt.Println("Transaction:", hash)
}
