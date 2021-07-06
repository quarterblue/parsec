package blockchain

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

const (
	dbPath       = "./tmp/blocks"
	dbFile       = "blockchain.db"
	blocksBucket = "blocks"
)

type Blockchain struct {
	head     []byte
	Database *bolt.DB
}

type ChainIterator struct {
	currentHash []byte
	Database    *bolt.DB
}

func (chain *Blockchain) Iterator() *ChainIterator {
	chainIterator := &ChainIterator{chain.head, chain.Database}

	return chainIterator
}

func (i *ChainIterator) Next() *Block {
	var block *Block

	err := i.Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = Deserialize(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.LastHash

	return block
}

// Creates a new blockchain if there is no existing blockchain and store Gensis block inside.
// If there is an existing blockchain, set the head of the blockchain instance to the last block
// block hash stored in the database
func CreateBlockchain() *Blockchain {
	var head []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		// There is no blockchain existing in the database, make new and add Gensis Block
		if b == nil {
			genesis := Genesis()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = b.Put([]byte("lh"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			head = genesis.Hash
		} else {
			// There is a blockchain existing in the database, the the head pointer to the blockchain's last hash
			head = b.Get([]byte("lh"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	blockchain := Blockchain{head, db}
	return &blockchain
}

func (chain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := CreateBlock(data, lastHash, 10)

	err = chain.Database.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		chain.head = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	// lastBlock := chain.Blocks[len(chain.Blocks)-1]
	// newBlock := CreateBlock(data, lastBlock)
	// chain.Blocks = append(chain.Blocks, newBlock)
	// return
}

// func (chain *Blockchain) ValidChain() bool {
// 	for i, block := range chain.Blocks {
// 		if i == 0 {
// 			continue
// 		}

// 		lastHash := chain.Blocks[i-1].Hash
// 		if bytes.Compare(lastHash, block.LastHash) != 0 {
// 			return false
// 		}

// 		info := bytes.Join([][]byte{util.Hexify(int64(block.Timestamp)),
// 			util.Hexify(int64(block.Nonce)),
// 			util.Hexify(int64(block.Difficulty)),
// 			block.Data, block.LastHash},
// 			[]byte{})

// 		trueHash := sha256.Sum256(info)
// 		if string(trueHash[:]) != string(block.Hash) {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (chain *Blockchain) ReplaceChain(newChain *Blockchain) {
// 	if len(chain.Blocks) >= len(newChain.Blocks) {
// 		fmt.Println("Current blockchain is longer than the incoming chain.")
// 		return
// 	}

// 	if !newChain.ValidChain() {
// 		fmt.Println("incoming chain.")
// 		return
// 	}
// 	chain.Blocks = newChain.Blocks
// }

// func (chain *Blockchain) PrintChain() {
// for _, block := range chain.Blocks {
// 	block.PrintBlock()
// }
// }
