package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
)

// Difficulty takes data from the block
const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int // requirement
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) // left shift

	PoW := &ProofOfWork{b, target}

	return PoW
}

func (PoW *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			PoW.Block.PreviousHash,
			PoW.Block.Data,
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
