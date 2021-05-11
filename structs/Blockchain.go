package structs

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

type Blockchain struct {
	Tip []byte
	DB *bolt.DB
}

const (
	dbFile = "blockchain_%s.db" // db file name
	blocksBucket = "blocks" // collection of blocks in db
	lastBlock = "l" // last block in blockchain
	genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks" // genesis block data
)

// MineBlock mines a new block with the provided transactions
func (bc *Blockchain) MineBlock(transactions []*Transaction) *Block {
	var lastHash []byte
	var lastHeight int

	for _, tx := range transactions {
		if bc.VerifyTransaction(tx) != true {
			log.Panic("ERROR: Invalid transaction")
		}
	}

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		blockData := b.Get(lastHash)
		block := DeserializeBlock(blockData)

		lastHeight = block.Height

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(transactions, lastHash, lastHeight + 1)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return newBlock
}

// CreateBlockchain creates a new blockchain DB
func CreateBlockchain(address string, nodeID string) *Blockchain {
	dbFile := fmt.Sprintf(dbFile, nodeID)
	if dbExists(dbFile) {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}

	var tip []byte

	cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
	genesis := NewGenesisBlock(cbtx)

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}
		tip = genesis.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

// NewBlockchain creates a new Blockchain with genesis's coinbase
func NewBlockchain(nodeID string) *Blockchain {
	dbFile := fmt.Sprintf(dbFile, nodeID)
	if dbExists(dbFile) == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := Blockchain{tip, db}
	return &bc
}

// get blockchain iterator: this iterator point to the last block
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.DB}

	return bci
}

// FindUTXO finds all unspent transaction outputs and returns transactions with spent outputs removed
func (bc *Blockchain) FindUTXO() map[string]TXOutputs {
	UTXO := make(map[string]TXOutputs)
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	// duyệt từng block
	for {
		block := bci.Next()

		// duyệt từng Transaction trong block
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID) // Transaction ID

		Outputs:
			// duyệt từng ouput của Transaction
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				// Thêm Transaction vào mảng những Transactions chưa được gửi của người dùng có đ/c là address
				//if out.IsLockedWithKey(pubKeyHash) {
				//	unspentTXs = append(unspentTXs, *tx)
				//}

				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
			}

			// nếu Transaction này không phải là Coinbase Transaction (coinbase sẽ không có input)
			if tx.IsCoinbase() == false {
				// duyệt từng input của Transaction
				for _, in := range tx.Vin {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return UTXO
}

// FindUTXO finds all unspent outputs of an user
//func (bc *Blockchain) FindUTXO(pubKeyHash []byte) []TXOutput {
//	var UTXOs []TXOutput
//	unspentTransactions := bc.FindUnspentTransactions(pubKeyHash)
//
//	for _, tx := range unspentTransactions {
//		for _, out := range tx.Vout {
//			if out.IsLockedWithKey(pubKeyHash) {
//				UTXOs = append(UTXOs, out)
//			}
//		}
//	}
//
//	return UTXOs
//}

// GetBalance calculates account's balance of an user in blockchain system
func (bc *Blockchain) GetBalance(address string) int {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	UTXOSet := UTXOSet{bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}

//// FindSpendableOutputs finds and returns unspent outputs to reference in inputs
//func (bc *Blockchain) FindSpendableOutputs(pubKeyHash []byte, amount int) (int, map[string][]int) {
//	unspentOutputs := make(map[string][]int)
//	unspentTXs := bc.FindUnspentTransactions(pubKeyHash)
//	accumulated := 0
//
//Work:
//	for _, tx := range unspentTXs {
//		txID := hex.EncodeToString(tx.ID)
//
//		for outIdx, out := range tx.Vout {
//			if out.IsLockedWithKey(pubKeyHash) && accumulated < amount {
//				accumulated += out.Value
//				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)
//
//				if accumulated >= amount {
//					break Work
//				}
//			}
//		}
//	}
//
//	return accumulated, unspentOutputs
//}

// FindTransaction finds a transaction by its ID
func (bc *Blockchain) FindTransaction(ID []byte) (Transaction, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			if bytes.Compare(tx.ID, ID) == 0 {
				return *tx, nil
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return Transaction{}, errors.New("Transaction is not found")
}

// SignTransaction signs inputs of a Transaction
func (bc *Blockchain) SignTransaction(tx *Transaction, privKey ecdsa.PrivateKey) {
	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	tx.Sign(privKey, prevTXs)
}

// VerifyTransaction verifies transaction input signatures
func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	return tx.Verify(prevTXs)
}

func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

// AddBlock saves the block into the blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		blockInDb := b.Get(block.Hash)

		if blockInDb != nil {
			return nil
		}

		blockData := block.Serialize()
		err := b.Put(block.Hash, blockData)
		if err != nil {
			log.Panic(err)
		}

		lastHash := b.Get([]byte("l"))
		lastBlockData := b.Get(lastHash)
		lastBlock := DeserializeBlock(lastBlockData)

		if block.Height > lastBlock.Height {
			err = b.Put([]byte("l"), block.Hash)
			if err != nil {
				log.Panic(err)
			}
			bc.Tip = block.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// GetBestHeight returns the height of the latest block
func (bc *Blockchain) GetBestHeight() int {
	var lastBlock Block

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash := b.Get([]byte("l"))
		blockData := b.Get(lastHash)
		lastBlock = *DeserializeBlock(blockData)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return lastBlock.Height
}

// GetBlock finds a block by its hash and returns it
func (bc *Blockchain) GetBlock(blockHash []byte) (Block, error) {
	var block Block

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		blockData := b.Get(blockHash)

		if blockData == nil {
			return errors.New("Block is not found.")
		}

		block = *DeserializeBlock(blockData)

		return nil
	})
	if err != nil {
		return block, err
	}

	return block, nil
}

// GetBlockHashes returns a list of hashes of all the blocks in the chain
func (bc *Blockchain) GetBlockHashes() [][]byte {
	var blocks [][]byte
	bci := bc.Iterator()

	for {
		block := bci.Next()

		blocks = append(blocks, block.Hash)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return blocks
}


