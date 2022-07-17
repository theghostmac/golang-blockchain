# Implementing a Blockchain with Go

## Introduction
In the first tutorial, we started building a Blockchain with Go. We went as far as creating a Genesis block, initializing the blockchain,
and adding relevant functions to the code. We will move on by refactoring the existing code to make the project properly structured.

## Refactoring the Existing Code
To continue the blockchain, we will create a `blockchain` package with two files: a `block.go` and a `PoW.go` file. All existing logic
in the `main.go` file except the `main` function should be copied to the `block.go` file. Go ahead to add visibility to the `blocks` field in 
the `Blockchain` struct and wherever it is used in the code. Next, import only the `fmt` package and the `blockchain` package you just created,
to the `main.g`o file. We will now create the consensus algorithm for the Blockchain.

## Implementing a Consensus Mechanism
