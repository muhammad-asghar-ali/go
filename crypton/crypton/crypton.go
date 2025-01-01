package crypton

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

const (
	path = "./tmp/blocks"
)

type (
	Crypton struct {
		LastHash []byte
		Database *badger.DB
	}

	CryptonIterator struct {
		CurrentHash []byte
		Database    *badger.DB
	}
)

func (c *Crypton) AddBlock(data string) {
	b := &Block{}
	lh := []byte{}

	err := c.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		HandleError(err)

		err = item.Value(func(val []byte) error {
			lh = append([]byte{}, val...)
			return nil
		})

		return err
	})

	HandleError(err)

	new := b.Create(data, lh)
	err = c.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(new.Hash, new.Serialize())
		HandleError(err)

		err = txn.Set([]byte("lh"), new.Hash)
		c.LastHash = new.Hash

		return err
	})

	HandleError(err)
}

func (c *Crypton) Init() *Crypton {
	b := &Block{}
	lh := []byte{}

	opts := badger.DefaultOptions(path)

	db, err := badger.Open(opts)
	HandleError(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := b.Genesis()

			fmt.Println("Genesis Proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			HandleError(err)

			err = txn.Set([]byte("lh"), genesis.Hash)

			lh = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			HandleError(err)

			err = item.Value(func(val []byte) error {
				lh = append([]byte{}, val...)
				return nil
			})

			return err
		}
	})

	HandleError(err)

	return &Crypton{lh, db}
}

func (c *Crypton) Iterator() *CryptonIterator {
	return &CryptonIterator{c.LastHash, c.Database}
}

func (iter *CryptonIterator) Next() *Block {
	block := &Block{}
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		HandleError(err)

		encoded := []byte{}
		err = item.Value(func(val []byte) error {
			encoded = append([]byte{}, val...)
			return nil
		})

		b := &Block{}
		block = b.Deserialize(encoded)
		return err
	})

	HandleError(err)

	iter.CurrentHash = block.PrevHash

	return block
}
