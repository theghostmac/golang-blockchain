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
We may now proceed to write a function that creates the new block. It will receive a string of datq
and the `PreviousHash` byte slice, and returns a pointer to a `Block`:
```go
func CreateBlock(data string, PreviousHash []byte) *Block {
	
}
```
This method will create a block. For the CurrentHash, we will put in an empty byte slice, the
data will be a string of data converted into a byte slice, and the `PreviousHash` will be passed
normally since it is already a byte slice. Moving on, we will call the `DeriveHash` method to 
derive a hash for this new block, before it is returned. The method becomes:
```go
func CreateBlock(data string, previousBlockHash []byte) *Block {
    block := &Block{[]byte{}, []byte(data), previousBlockHash}
    block.DeriveHash()
    return block
}
```
What we have now is a program that collects a data, hashes it, then creates a block to store it.
What makes it a blockchain is if there exist more than one block linked together as a chain. 
This means we have to create a `Blockchain` type of blocks, having an array of pointers to blocks:
```go
type Blockchain struct {
	blocks []*Block
}
```
Next, we will create a method that will merge a block to the chain. It will point to the `Blockchain` and 
take in a string of data. We will take the previous block, create a new block, then merge the new block to 
the chain of blocks. Given, it doesn't make any sense for the first block in the blockchain to have a
`previousBlock`.