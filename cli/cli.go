package cli

import (
	"github.com/quarterblue/parsec/blockchain"
)

type CLI struct {
	chain *blockchain.Blockchain
}

func (cli *CLI) printHelp() {
}
