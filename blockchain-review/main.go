package blockchain_review

import (
	"bytes"
	"crypto/sha256"
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
	b.Hash = hash[:]
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

func main() {

}
