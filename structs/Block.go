package structs

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce			int
}

// hash data of a block (timestamp, data, prevBlockHash) and assign it to Block.Hash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// create a new block with Data and Previous Block's Hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run() // finding nonce
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// initialize the Genesis Block (1st block in a blockchain)
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // a block without previous block's hash
}