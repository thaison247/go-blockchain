package structs

import (
	"encoding/hex"
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
	genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks" // genesis block data
)

// MineBlock mines a new block with the provided transactions
func (bc *Blockchain) MineBlock(transactions []*Transaction) {
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	newBlock := NewBlock(transactions, lastHash)

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

// NewBlockchain creates a new Blockchain with genesis's coinbase sent for an specific address
func NewBlockchain(address string) *Blockchain {
	var Tip []byte
	db, err := bolt.Open(dbFile, 0600, nil) // open database
	if err != nil {
		fmt.Println(err)
	}

	err = db.Update(func(tx *bolt.Tx) error { // open db read-write transaction
		b := tx.Bucket([]byte(blocksBucket)) // open "blocks" collection in db

		if b == nil { // if "blocks" collection does not exist
			genesis := NewGenesisBlock(NewCoinbaseTX(address, genesisCoinbaseData)) // create genesis block
			b, err := tx.CreateBucket([]byte(blocksBucket)) // create "blocks" collection
			if err != nil {
				fmt.Println(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize()) // add {key: <genesis has>, value: <genesis block in array of bytes>} to collection
			err = b.Put([]byte(lastBlock), genesis.Hash) // the last block's in blockchain
			Tip = genesis.Hash
		} else { // if "blocks" collection does exist
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

// FindUnspentTransactions : tìm các Transactions có chứa unspent outputs của người dùng address
func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
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
				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			// nếu Transaction này không phải là Coinbase Transaction (coinbase sẽ không có input)
			if tx.IsCoinbase() == false {
				// duyệt từng input của Transaction
				for _, in := range tx.Vin {
					// nếu input có thể mở khóa bằng address
					// (~ đây là một giao dịch mà người dùng (address) gửi tiền đi từ một ouput của transaction khác)
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid) // transtion id của 1 transaction trước đó mà chứa output của giao dịch gửi tiền cho người dùng có đ/c address
						// thêm bộ {key: value} = {tx's ID: vị trí của spent output trong transaction đó}
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
						// ---> spentTXOs: [key: là id của transaction chứa spent output của ngươi dùng address
						//                  value:  danh sách những vị trí của spent output trong transaction đó]
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTXs
}

// FindUTXO finds all unspent outputs of an user
func (bc *Blockchain) FindUTXO(address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

// GetBalance calculates account's balance of an user in blockchain system
func (bc *Blockchain) GetBalance(address string) int {
	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}


