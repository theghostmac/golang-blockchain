package blockchain

//
//import (
//	"bytes"
//	"crypto/sha256"
//	"encoding/binary"
//	"fmt"
//	"log"
//	"math"
//	"math/big"
//)
//
//// Difficulty takes data from the block
//const Difficulty = 12
//
//type ProofOfWork struct {
//	Block  *Block
//	Target *big.Int // requirement
//}
//
//func NewProof(b *Block) *ProofOfWork {
//	target := big.NewInt(1)
//	target.Lsh(target, uint(256-Difficulty)) // left shift
//
//	PoW := &ProofOfWork{b, target}
//
//	return PoW
//}
//
//func (PoW *ProofOfWork) InitData(nonce int) []byte {
//	data := bytes.Join(
//		[][]byte{
//			PoW.Block.PreviousHash,
//			PoW.Block.Data,
//		},
//		[]byte{},
//	)
//	return data
//}
//
//func ToHex(num int64) []byte {
//	buffer := new(bytes.Buffer)
//	err := binary.Write(buffer, binary.BigEndian, num)
//	if err != nil {
//		log.Panic(err)
//	}
//	return buffer.Bytes()
//}
//
//func (PoW *ProofOfWork) Run() (int, []byte) {
//	var intHash big.Int
//	var hash [32]byte
//
//	nonce := 0
//
//	// nonce is an infinite loop that hashes
//	for nonce < math.MaxInt64 {
//		data := PoW.InitData(nonce)
//		hash = sha256.Sum256(data)
//
//		fmt.Printf("\r%x", hash)
//		intHash.SetBytes(hash[:])
//
//		if intHash.Cmp(PoW.Target) == -1 {
//			break // hash is less than target
//		} else {
//			nonce++
//		}
//	}
//	fmt.Println()
//	return nonce, hash[:]
//}
