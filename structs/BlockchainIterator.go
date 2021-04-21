package structs

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

// get the next block in blockchain (actually... previous block)
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	if i.currentHash == nil {
		return nil
	}

	err := i.db.View(func(tx *bolt.Tx) error { // open read db transaction
		b := tx.Bucket([]byte(blocksBucket)) // open collection "blocks"
		encodedBlock := b.Get(i.currentHash) // get block by key
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	i.currentHash = block.PrevBlockHash

	return block
}