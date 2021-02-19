package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/quarterblue/parsec/util"
)

type Blockchain struct {
	Blocks []*Block
}

func CreateBlockchain() *Blockchain {
	blockchain := &Blockchain{[]*Block{Genesis()}}
	return blockchain
}

func (chain *Blockchain) AddBlock(data string) {
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, lastBlock)
	chain.Blocks = append(chain.Blocks, newBlock)
	return
}

func (chain *Blockchain) ValidChain() bool {

	for i, block := range chain.Blocks {
		if i == 0 {
			continue
		}

		lastHash := chain.Blocks[i-1].Hash
		if bytes.Compare(lastHash, block.LastHash) != 0 {
			return false
		}

		info := bytes.Join([][]byte{util.Hexify(int64(block.Timestamp)),
			util.Hexify(int64(block.Nonce)),
			util.Hexify(int64(block.Difficulty)),
			block.Data, block.LastHash},
			[]byte{})

		trueHash := sha256.Sum256(info)
		if string(trueHash[:]) != string(block.Hash) {
			return false
		}
	}

	return true
}

func (chain *Blockchain) ReplaceChain(newChain *Blockchain) {
	if len(chain.Blocks) >= len(newChain.Blocks) {
		fmt.Println("Current blockchain is longer than the incoming chain.")
		return
	}

	if !newChain.ValidChain() {
		fmt.Println("incoming chain.")
		return
	}
	chain.Blocks = newChain.Blocks
}

func (chain *Blockchain) PrintChain() {
	for _, block := range chain.Blocks {
		block.PrintBlock()
	}
}
