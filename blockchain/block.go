package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block struct derives a scaffold for a single  block in reversed linked-list format
type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int
}

//// GetHash method derives the hash
//func (b *Block) GetHash() {
//	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{}) // joins the byte slices together
//	hash := sha256.Sum256(info)                                    // sha256 is a simple hash function
//	b.Hash = hash[:]
//}

// CreateBlock creates a new block and derives a hash for it
func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	ProofOfWork := NewProof(block)
	nonce, hash := ProofOfWork.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis block function creates the first block, having only a data type
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize the blockchain data structure into a slice of bytes for the BadgerDB to persist
func (b *Block) Serialize() []byte {
	var resultByteBuffer bytes.Buffer
	encoder := gob.NewEncoder(&resultByteBuffer)

	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}
	return resultByteBuffer.Bytes()
}

// Deserialize a slice of byte data and output a pointer to a Block
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}
	return &block
}

// Handle function to handle all errors
func ()  {
	
}