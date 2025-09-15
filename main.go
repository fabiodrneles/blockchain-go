package main

import (
	"fmt"

	"github.com/fabiodrneles/blockchain-go/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
