# Implementing a Blockchain in Go
## What a blockchain is made of
A blockchain is made up of a chain of different blocks. Each of the blocks contains
data that we want to add to the database. They also contain hashes of the past and next block.
Using structs, we can group associated data to one custom type:
```go
type Block struct {
	CurrentHash         []byte
	Data         []byte
	PreviousHash []byte
}
```
This structure shows that each block in a blockchain references the previous block before it. 
Any current block in the chain derives a hash based on the data and the previous hash (or the 
last block's hash).
A function is required to derive this hash for the current block. We will create a method to do
that:
```go
func (b *Block) DeriveHash() {
	
}
```
We used a pointer receiver method because we will be modifying the value of the dissident 
struct. The `Join` function in the `bytes` package is used to join the slices of byte together.
We will join the `Data` byte slice and the `PreviousHash` byte slice together as a two-dimensional 
slice of byte, and then an empty slice of byte.
The method becomes:
```go
func (b *Block) DeriveHash() {
    newHashData := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
}
```
Moving further, we will create a `sum256` hash for the current block, depending on the `Data` and the
`PreviousHash` i.e. `newHashData`. After that, we will push it into the block. That line of code is:
```go
hash := sha256.Sum256(newHashData)
b.Hash = hash[:]
```
