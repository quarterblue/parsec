package main

import (
	"github.com/quarterblue/parsec/blockchain"
)

func main() {

	bchain := blockchain.CreateBlockchain()
	bchain.AddBlock("Hello World")
	bchain.AddBlock("Hello Again World!")
	bchain.AddBlock("Hello Again & Again World")
	bchain.PrintChain()
}
