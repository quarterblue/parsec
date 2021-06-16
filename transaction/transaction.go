package transaction

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Transaction has a combination of inputs and outputs.
// Outputs are where the coin is
// Inputs are outputs from previous transactions
type Transaction struct {
	ID   []byte
	Tin  []TXInput
	Tout []TXOutput
}

type TXOutput struct {
	Value  int
	Pubkey string
}

type TXInput struct {
	ID        []byte
	Tout      int
	Signature string
}

// Coinbase transaction is a special type of transaction which doesn't require
// Previous existing outputs, similar to a gensis block
func InitCoinbaseTX(out, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Giving Parsec Coin to %s", out)
	}
	tx := Transaction{nil, nil, nil}
	return &tx
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	// var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)

	if err != nil {
		log.Panic(err)
	}
}
