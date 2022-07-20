package blockchain

// Blockchain struct creates a chain of other blocks using the Block scaffold
type Blockchain struct {
	Blocks []*Block
}

// InitBlockchain function starts the blockchain relying first on the Genesis block
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

// AddBlock method points to the Blockchain, adds a block to the chain, and gives a data string
func (chain *Blockchain) AddBlock(data string) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
