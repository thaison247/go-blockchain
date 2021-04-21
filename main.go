package main

import (
	"fmt"
	"github.com/thaison247/go-blockchain/structs"
	"strconv"
)

func main() {
	// create a new blockchain
	bc := structs.NewBlockchain()

	// add some blocks
	bc.AddBlock("Son sent 1 dollar to LA")
	bc.AddBlock("Son sent 1 dollar to Tai")

	//// show blockchain
	//for _, block := range bc.Blocks {
	//	fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Timestamp: %x\n", strconv.FormatInt(block.Timestamp, 10))
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	fmt.Printf("Validation: %v\n\n", structs.NewProofOfWork(block).Validate())
	//}

	it := bc.Iterator()
	for {
		block := it.Next()
		if block == nil {
			break
		}

		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Timestamp: %x\n", strconv.FormatInt(block.Timestamp, 10))
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Validation: %v\n\n", structs.NewProofOfWork(block).Validate())
	}
}
