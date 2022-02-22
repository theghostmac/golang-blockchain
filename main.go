package main

import (
	"bytes"
	"crypto/sha256"
)

// Block struct implements blockchain reversed linkedlist format
type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

// GetHash method
func (b *Block) GetHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{}) // joins the byte slices together
	hash := sha256.Sum256(info)                                    // sha256 is a simple hash function
	b.Hash = hash[:]
}
func main() {

}
