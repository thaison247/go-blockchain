package structs

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type Blockchain struct {
	Tip []byte
	DB *bolt.DB
}

const (
	dbFile = "blockchain.db" // db file name
	blocksBucket = "blocks" // collection of blocks in db
	lastBlock = "l" // last block in blockchain
)

// add one block to Blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		// add {key: <new block's hash>, value: <new block in array of bytes>}
		if err := b.Put(newBlock.Hash, newBlock.Serialize()); err != nil {
			fmt.Println(err)
		}

		// change last block of the blockchain to new block just added
		if err = b.Put([]byte(lastBlock), newBlock.Hash); err != nil {
			fmt.Println(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})
}

// create a new Blockchain
func NewBlockchain() *Blockchain {
	var Tip []byte
	db, err := bolt.Open(dbFile, 0600, nil) // open database
	if err != nil {
		fmt.Println(err)
	}

	err = db.Update(func(tx *bolt.Tx) error { // open db read-write transaction
		b := tx.Bucket([]byte(blocksBucket)) // open "blocks" collection in db

		if b == nil { // if "blocks" collection does not exist
			genesis := NewGenesisBlock() // create genesis block
			b, err := tx.CreateBucket([]byte(blocksBucket)) // create "blocks" collection
			if err != nil {
				fmt.Println(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize()) // add {key: <genesis has>, value: <genesis block in array of bytes>} to collection
			err = b.Put([]byte(lastBlock), genesis.Hash) // the last block's in blockchain
			Tip = genesis.Hash
		} else { // if "blocks" collection does eixst
			Tip = b.Get([]byte("l")) // get the block's hash
		}

		return nil
	})

	bc := Blockchain{Tip, db}

	return &bc
}

// get blockchain iterator: this iterator point to the last block
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.DB}

	return bci
}

