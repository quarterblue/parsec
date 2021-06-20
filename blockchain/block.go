package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
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

// Prints block information for debugging purposes
func (block *Block) PrintBlock() {
	unixTime := time.Unix(block.Timestamp, 0)
	fmt.Println("Timestamp: ", unixTime)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("Last Hash: %x\n", block.LastHash)
}

func CreateBlock(data string, lastBlock []byte, difficulty int) *Block {
	block := &Block{util.MakeTimeStamp(), []byte(data), []byte{}, lastBlock, 0, difficulty}
	pow := CreatePow(block)
	nonce, hash := pow.ComputeHash()
	block.Nonce = nonce
	block.Hash = hash[:]

	// Adjust difficulty
	// lastTime := lastBlock.Timestamp
	// block.AdjustDiff(lastTime)

	return block
}

// Generates the first hardcoded Gensis block.
// The data used for the hash is NYTimes headline title from August 7, 2020.
func Genesis() *Block {
	data := "NYTimes Aug 7, 2020 Job Growth Slowed in July, Signaling a Loss of Economic Momentum"
	block := &Block{1596783600, []byte(data), []byte{}, []byte{}, 0, InitDifficulty}
	pow := CreatePow(block)
	nonce, hash := pow.ComputeHash()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

// Utility function to adjust the difficulty of the proof of work algorithm.
// Difficulty is adjusted based on mineRate set at top line
func (block *Block) AdjustDiff(timestamp int64) {
	difference := timestamp - block.Timestamp
	if difference > mineRate {
		block.Difficulty--
		return
	}
	block.Difficulty++
}

// Serializes a block into bytestring for persistance storage
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// Deserializes bytestring into block when loading from persistance storage
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}
	return &block
}
