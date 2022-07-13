package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	CurrentHash  []byte
	Data         []byte
	PreviousHash []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	newHashData := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash := sha256.Sum256(newHashData)
	b.CurrentHash = hash[:]
}

func CreateBlock(data string, previousBlockHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousBlockHash}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	previousBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, previousBlock.CurrentHash)
	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitialBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {
	chain := InitialBlockchain()

	chain.AddBlock("This is the First Block after Genesis")
	chain.AddBlock("This is the Second Block after Genesis")
	chain.AddBlock("This is the Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Data contained in Block: %s\n", block.Data)
		fmt.Printf("Current Hash: %x\n", block.CurrentHash)
	}
}
