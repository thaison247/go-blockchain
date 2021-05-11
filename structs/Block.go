package structs

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce			int
	Height			int
}

//  hashes data of a block (timestamp, data, prevBlockHash) and assign it to Block.Hash
//func (b *Block) SetHash() {
//	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
//	hash := sha256.Sum256(headers)
//
//	b.Hash = hash[:]
//}

// NewBlock creates a new block with Transactions and Previous Block's Hash
func NewBlock(transactions []*Transaction, prevBlockHash []byte, height int) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0, height}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run() // finding nonce
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock initializes the Genesis Block (1st block in a blockchain)
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0) // a block without previous block's hash
}

// Serialize serializes block to array of bytes
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	if err := encoder.Encode(b); err != nil {
		fmt.Println(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes array of bytes to block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	if err := decoder.Decode(&block); err != nil {
		fmt.Println(err)
	}

	return &block
}

// HashTransactions hashes all transactions in block. Return array of bytes.
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

