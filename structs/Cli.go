package structs

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {}

// Run command-line interface
func (cli *CLI) Run() {
	//addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	//printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	//addBlockData := addBlockCmd.String("data", "", "Block data")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")

	switch os.Args[1] {
	//case "addblock":
	//	addBlockCmd.Parse(os.Args[2:])
	//case "printchain":
	//	printChainCmd.Parse(os.Args[2:])
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	//case "printchain":
	//	err := printChainCmd.Parse(os.Args[2:])
	//	if err != nil {
	//		log.Panic(err)
	//	}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	//if addBlockCmd.Parsed() {
	//	fmt.Println(*addBlockData)
	//	if *addBlockData == "" {
	//		addBlockCmd.Usage()
	//		os.Exit(1)
	//	}
	//	cli.addBlock(*addBlockData)
	//}
	//
	//if printChainCmd.Parsed() {
	//	cli.printChain()
	//}

	if createBlockchainCmd.Parsed() {
		fmt.Println(*createBlockchainAddress)
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}
}

// createBlockchain creates a new Blockchain
func (cli *CLI) createBlockchain(address string) {
	bc := NewBlockchain(address)
	bc.DB.Close()
	fmt.Println("Done!")
}


// add new block to blockchain
//func (cli *CLI) addBlock(data string) {
//	cli.BC.MineBlock(data)
//	fmt.Println("Success!")
//}

// print chain of blocks
//func (cli *CLI) printChain() {
//	bci := cli.BC.Iterator()
//
//	for {
//		block := bci.Next()
//
//		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
//		fmt.Printf("Data: %s\n", block.Data)
//		fmt.Printf("Hash: %x\n", block.Hash)
//		pow := NewProofOfWork(block)
//		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
//		fmt.Println()
//
//		if len(block.PrevBlockHash) == 0 {
//			break
//		}
//	}
//}

func (cli *CLI) printUsage() {
	fmt.Println("flags:")
	fmt.Printf("/t\"addblock\": addblock \"[data]\" ---> Add new block to blockchain with some data\n")
	fmt.Printf("/t\"printchain\": printchain ---> Print all blocks in blockchain from latest to oldest\n")
}

func (cli *CLI) getBalance(address string) {
	bc := NewBlockchain(address)
	defer bc.DB.Close()

	balance := bc.GetBalance(address)
	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
