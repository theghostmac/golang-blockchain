package main

import (
	"fmt"
	"github.com/theghostmac/golang-blockchain/blockchain"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockchain() //start the blockchain
	chain.AddBlock("Busta's Node")       //adds new block to the blockchain
	chain.AddBlock("Goodness' Node")     //adds new block joining the first block
	chain.AddBlock("Perelyn's Node")     //adds new block to the second block
	// everybody gets a block
	chain.AddBlock("Brymes' Node")
	chain.AddBlock("Noble's Node")   //adds new block to the second block
	chain.AddBlock("Hendrix's Node") //adds new block to the second block
	chain.AddBlock("Oyin's Node")    //adds new block to the second block

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash is: %s\n", block.PreviousHash)
		fmt.Printf("The data in the block is: %s\n", block.Data)
		fmt.Printf("The current block's hash is: %s\n", block.Hash)
	}

	// added proof of work algorithm to the main function
	PoW := blockchain.NewProof()
	fmt.Printf("PoW: %s\n", strconv.FormatBool(PoW.Validate()))
	fmt.Println()
}
