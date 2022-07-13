# Implementing a Blockchain in Go
## What a blockchain is made of
A blockchain is made up of a chain of different blocks. Each of the blocks contains
data that we want to add to the database. They also contain hashes of the past and next block.
Using structs, we can group associated data to one custom type:
```go
type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}
```
This structure shows that each block in a blockchain references the previous block before it. 
Any current block in the chain derives a hash based on the data and the previous hash (or the 
last block's hash). 