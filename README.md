# verbose-octo-palm-tree

## Transaction problem

- This app uses `Görli` testnet. You'll need a wallet in that network.

- You can use this faucet to obtain some GoerliETH: https://goerli-faucet.pk910.de/

- Navigate to directory:

  `cd 1-transaction`

- Run app using this command:

  `go run . -from={your_private_key} -to={receiver_address} -amount={amount_to_send}`

- `from` flag is required
- Receiver address (`to` flag) by default is public Görli Faucet, you can omit if you want
- Amount by default is 1000 GöWei, you can omit if you want
- If everything goes well, it'll print the transaction hash.
- You can use etherscan.com to validate the transaction.
- If there is an error, respective error message will be printed on console.

## Teacher Problem

- Navigate to directory:

  `cd 2-teacher`

- Run program:

  `go run .`

- For testing:

  `go test ./...`

- After running the program, press `h` for help.
