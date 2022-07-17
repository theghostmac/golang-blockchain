# Implementing a Blockchain in Go
## Introduction
Building a blockchain is different from developing on a blockchain network. While the latter is called 
blockchain development, the former is called blockchain engineering. To implement blockchain technology, 
you must learn a plethora of concepts that serve as the bedrock of the blockchain technology. This 
article does not offer the theoretical knowledge required to build a blockchain network; instead it provides
a practical and more hands-on approach to building a blockchain.

## What is a blockchain
Blockchain is a subset of distributed ledger technology. It is a kind of decentralized and distributed
database ledger system used for making transactions open to everyone on the chain. One of the most critical feature 
of blockchain technology that makes it different from several other database systems is timestamping.

## What a blockchain is made of
A blockchain is made up of a chain of different blocks. Each of the blocks contains
data (or transactions) that we want to add to the database. The transactions are grouped and appended
to the blockchain as a unit. They also contain hashes of the past and next block.
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
the chain of blocks. 
```go
func (chain *Blockchain) AddBlock(data string) {
	previousBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, previousBlock.CurrentHash)
	chain.blocks = append(chain.blocks, newBlock)
}
```
Given, it doesn't make any sense for the first block in the blockchain to have a `previousBlock`. Thus,
we will create a function that uses the first a.k.a Genesis block to create an initial version of the 
blockchain. It will return a reference to the Blockchain type having an array of pointers to `Block` 
while calling the Genesis function.
```go
func Genesis() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitialBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
	
}
```

Finally, let's link the entire blockchain system together in the `main` function. This is done by
creating a chain and initializing the blockchain. You can then go ahead to add data to as many 
blocks as you want with the `AddBlock` function.
To see the blocks as a blockchain, you can run a for loop to range through the contents of each block. 
The main function looks thus:
```go
func main() {
	chain := InitialBlockchain()

	chain.AddBlock("This is the First Block after Genesis")
	chain.AddBlock("This is the Second Block after Genesis")
	chain.AddBlock("This is the Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PreviousHash)
		fmt.Printf("Data contained in Block: %s\n", block.Data)
		fmt.Printf("Current Hash: %x\n", block.CurrentHash)
	}
}

```

The output after running the source file is similar to the following:
```bash
Previous Hash: 
Data contained in Block: Genesis Block
Current Hash: 89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3
Previous Hash: 89eb0ac031a63d2421cd05a2fbe41f3ea35f5c3712ca839cbf6b85c4ee07b7a3
Data contained in Block: This is the First Block after Genesis
Current Hash: f89f193b0b7dfd647bf362fa4a8700aaf8434aedd1be079bb605166571aaefd8
Previous Hash: f89f193b0b7dfd647bf362fa4a8700aaf8434aedd1be079bb605166571aaefd8
Data contained in Block: This is the Second Block after Genesis
Current Hash: 9e2e5b35b53748d12b3d39d9e32b970b05c3c8643b42bfdcbd5ee61a4088120c
Previous Hash: 9e2e5b35b53748d12b3d39d9e32b970b05c3c8643b42bfdcbd5ee61a4088120c
Data contained in Block: This is the Third Block after Genesis
Current Hash: 4a6e691c26c0436c5fa748825ef355b72f7c057302387cec300d29e64aaa01e5
```
You will notice that the first `PreviousHash` is empty, as it is the `Genesis` block. If you rerun 
the blockchain system, the hashes remain the same. If the data is chained in any block, the hash
becomes something different.

## Conclusion
In this article, you gained a practical knowledge of how blockchain technology is implemented. We looked
at how blocks are created and added to a blockchain. We moved further than creating the logic for the 
genesis block and new blocks to creating the logic for the blockchain itself. In 
subsequent tutorials, we will see how to set up a wallet and implement the proof of work consensus 
protocol for our blockchain.