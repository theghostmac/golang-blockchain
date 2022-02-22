package main

import (
	"bytes"
	"crypto/sha256"
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

// CreateBlock creates a new block and derives a hash
func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.GetHash()
	return block
}

// AddBlock method points to the Blockchain and gives a data string
func (chain *Blockchain) AddBlock(data string) {

}

func main() {

}
