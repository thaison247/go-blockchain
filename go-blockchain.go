package main

import (
	"github.com/thaison247/go-blockchain/structs"
)

func main() {
	// create a new blockchain
	//bc := structs.NewBlockchain()
	//defer bc.DB.Close()

	cli := structs.CLI{}
	cli.Run()
}
