package blockchain

import (
	"fmt"
	"time"

	"github.com/quarterblue/parsec/util"
)

var mineRate int64 = 100000

type Block struct {
	Timestamp  int64  `json:"timestamp"`
	Data       []byte `json:"data"`
	Hash       []byte `json:"hash"`
	LastHash   []byte `json:"lasthash"`
	Nonce      int    `json:"nonce"`
	Difficulty int    `json:"difficulty"`
}

func (block *Block) PrintBlock() {
	unixTime := time.Unix(block.Timestamp, 0)
	fmt.Println("Timestamp: ", unixTime)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("Last Hash: %x\n", block.LastHash)
}

func CreateBlock(data string, lastBlock *Block) *Block {

	block := &Block{util.MakeTimeStamp(), []byte(data), []byte{}, lastBlock.Hash, 0, lastBlock.Difficulty}
	pow := CreatePow(block)
	nonce, hash := pow.ComputeHash()
	block.Nonce = nonce
	block.Hash = hash

	// Adjust difficulty
	// lastTime := lastBlock.Timestamp
	// block.AdjustDiff(lastTime)

	return block
}

func Genesis() *Block {
	data := "NYTimes Aug 7, 2020 Job Growth Slowed in July, Signaling a Loss of Economic Momentum"
	block := &Block{1596783600, []byte(data), []byte{}, []byte{}, 0, InitDifficulty}
	pow := CreatePow(block)
	nonce, hash := pow.ComputeHash()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

func (block *Block) AdjustDiff(timestamp int64) {
	difference := timestamp - block.Timestamp
	if difference > mineRate {
		block.Difficulty--
		return
	}
	block.Difficulty++
}
