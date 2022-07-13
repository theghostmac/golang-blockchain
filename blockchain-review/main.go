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

func (b *Block) DeriveHash() {
	newHashData := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash := sha256.Sum256(newHashData)
	b.Hash = hash[:]
}

func main() {

}
