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

## Code Overview

### Block Structure

```go
type (
    Block struct {
        Hash     []byte
        Data     []byte
        PrevHash []byte
        Nonce    int
    }
)
```

- **Hash**: The unique identifier of the block.
- **Data**: The data stored in the block.
- **PrevHash**: The hash of the previous block.
- **Nonce**: The nonce value used for proof-of-work.

### Proof-of-Work

Crypton uses a simple proof-of-work mechanism to validate blocks:

```go
func (pow *ProofOfWork) Run() (int, []byte) {
    var intHash big.Int
    hash := [32]byte{}
    nonce := 0

    for nonce < math.MaxInt64 {
        data := pow.InitData(nonce)
        hash = sha256.Sum256(data)
        intHash.SetBytes(hash[:])

        if intHash.Cmp(pow.Target) == -1 {
            break
        } else {
            nonce++
        }
    }

    return nonce, hash[:]
}
```

### Persistent Storage

BadgerDB is used to persist the blockchain data. The lh key stores the hash of the last block:

```go
err = db.Update(func(txn *badger.Txn) error {
    err = txn.Set(genesis.Hash, genesis.Serialize())
    HandleError(err)
    err = txn.Set([]byte("lh"), genesis.Hash)
    return err
})
```

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
