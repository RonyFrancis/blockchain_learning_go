package main

import (
	"fmt"

	"github.com/RonyFrancis/blockchain_learning_go/blockchain"
)

func main() {
	fmt.Println("blockchain begining")
	// block := blockchain.NewBlock("first data ", []byte("prevBlockHash"))
	// fmt.Println(string(block.PrevBlockHash))

	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PoW: %t\n", block.IsValide())
	}
}
