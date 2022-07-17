# Implementing a Blockchain with Go

## Introduction
In the first tutorial, we started building a Blockchain with Go. We went as far as creating a Genesis block, initializing the blockchain,
and adding relevant functions to the code. We will move on by refactoring the existing code to make the project properly structured.

## Refactoring the Existing Code
To continue the blockchain, we will create a `blockchain` package with two files: a `block.go` and a `PoW.go` file. All existing logic
in the `main.go` file except the `main` function should be copied to the `block.go` file. Go ahead to add visibility to the `blocks` field in 
the `Blockchain` struct and wherever it is used in the code. Next, import only the `fmt` package and the `blockchain` package you just created,
to the `main.g`o file. We will now create the consensus algorithm for the Blockchain.

## Implementing a Consensus Algorithm
There are different types of consensus algorithms to choose from. However, we will be adopting the Proof-of-Work algorithm used by
bitcoin. PoW forces nodes in the network to do work in order to add a block to the chain. The work here is actually compute power as evident
from the case of miners. The job of the miners in running the PoW algorithm contributes to the overall safety of the blockchain. 

> if a user does work by signing a block, they need to show proof of this work done.

The procedure involves the following steps:
- collecting a data or transaction from the block, 
- creating a counter (or nonce) initialized at 0,
- creating a hash using the data and the counter,
- checking the hash to see if it meets some standard requirements,
- if it does not meet the requirements, we recreate another hash repeatedly until it does.

The requirements include the following:
- the first few bytes must contain 0s

In the `PoW.go` file, we will create a `Difficulty` constant set to 12.