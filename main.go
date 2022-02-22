package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Block struct derives a scaffold for a single  block in reversed linked-list format
type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

// Blockchain struct creates a chain of other blocks using the Block scaffold
type Blockchain struct {
	blocks []*Block
}

// GetHash method derives the hash
func (b *Block) GetHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{}) // joins the byte slices together
	hash := sha256.Sum256(info)                                    // sha256 is a simple hash function
	b.Hash = hash[:]
}

// CreateBlock creates a new block and derives a hash for it
func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.GetHash()
	return block
}

// Genesis block function creates the first block, having only a data type
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// AddBlock method points to the Blockchain, adds a block to the chain, and gives a data string
func (chain *Blockchain) AddBlock(data string) {
	previousBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// InitBlockchain function starts the blockchain relying first on the Genesis block
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockchain()         //start the blockchain
	chain.AddBlock("Busta's Block")   //adds new block to the blockchain
	chain.AddBlock("Goodness' Block") //adds new block joining the first block
	chain.AddBlock("Perelyn's Block") //adds new block to the second block
	// everybody gets a block
	chain.AddBlock("Brymes' Block")
	chain.AddBlock("Noble's Block")   //adds new block to the second block
	chain.AddBlock("Hendrix's Block") //adds new block to the second block
	chain.AddBlock("Oyin's Block")    //adds new block to the second block

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash is: %s\n", block.PreviousHash)
		fmt.Printf("The data in the block is: %s\n", block.Data)
		fmt.Printf("The current block's hash is: %s\n", block.Hash)
	}
}
