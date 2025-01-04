# Crypton Blockchain

Crypton is a simple blockchain implementation in Go, designed to provide a basic understanding of blockchain concepts. It includes functionalities to add blocks, validate proof-of-work, and traverse the blockchain.

## Features

- Add blocks to the blockchain.
- Validate proof-of-work for blocks.
- Serialize and deserialize blocks.
- Iterate through the blockchain.
- Persistent storage using BadgerDB.

## Prerequisites

- [Go](https://golang.org/) (version 1.18 or later).
- [BadgerDB](https://github.com/dgraph-io/badger) for persistent storage.

## Installation

1.  git clone cd crypton
2.  go mod tidy
3.  go run main.go

## Usage

### Add a Block

To add a block to the blockchain, use the add command:

`go run main.go add -block ""`

Example:
`go run main.go add -block "Block 1 Data"`

### Print the Blockchain

To print all blocks in the blockchain, use the print command:

`go run main.go print`

## Contribution

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- [BadgerDB](https://github.com/dgraph-io/badger) for the storage engine.
- The Go programming language.

Feel free to explore and experiment with Crypton to learn more about blockchain technologies!

```

```
