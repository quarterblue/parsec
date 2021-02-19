package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	"github.com/quarterblue/parsec/util"
)

const InitDifficulty = 12

const maxNonce = math.MaxInt64

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func CreatePow(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-b.Difficulty))
	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) hashData(nonce int) []byte {
	fmt.Println(nonce)
	hashedData := bytes.Join(
		[][]byte{
			util.Hexify(int64(pow.Block.Timestamp)),
			util.Hexify(int64(nonce)),
			util.Hexify(int64(pow.Block.Difficulty)),
			pow.Block.Data,
			pow.Block.LastHash,
		},
		[]byte{})

	return hashedData
}

func (pow *ProofOfWork) ComputeHash() (int, []byte) {

	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	// // Adjust Time
	// lastTime := lastBlock.Timestamp
	// block.AdjustDiff(lastTime)

	for nonce < maxNonce {
		data := pow.hashData(nonce)
		hash := sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		fmt.Printf("%x\n", hash)

		if hashInt.Cmp(pow.Target) == -1 {
			return nonce, hash[:]
		} else {
			pow.Block.Timestamp = util.MakeTimeStamp()
			nonce++
		}
	}
	return nonce, hash[:]
}
