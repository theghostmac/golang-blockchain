package blockchain

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
	Blocks []*Block
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
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// InitBlockchain function starts the blockchain relying first on the Genesis block
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}
