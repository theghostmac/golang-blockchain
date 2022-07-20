## Persisting Data in the Blockchain
There's no specification about the type of database to be used in blockchains, it was not mentioned in the bitcoin 
white paper. Any kind of database can be used. Bitcoin uses LevelDB. But I use BadgerDB, pending the time when I complete
the [BacchusDB](https://github.com/theghostmac/BacchusDB) I am building in Go.

BadgerDB, however, is a native Golang database. It uses a key-value storage db. It doesn't use sql tables, no query 
statements or rows or columns. It accepts arrays or slices of bytes, and stores everything as byte keys and values.

## Procedure
Create file `blockchain.go` and transport the struct, `InitBlockchain()` and `AddBlock()` functions to it. Create a way to
serialize and deserialize the block data structure into bytes.