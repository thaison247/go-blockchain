package main

import (
	"fmt"
	"github.com/thaison247/go-blockchain/structs"
)

func main() {
	bc := structs.NewBlockchain("Ivan Novic")
	defer bc.DB.Close()

	balance := bc.GetBalance("Ivan Novic")
	fmt.Printf("Balance of '%s': %d\n", "Ivan Novic", balance)
}
