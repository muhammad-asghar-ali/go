package crypton

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type (
	Block struct {
		Hash     []byte
		Data     []byte
		PrevHash []byte
		Nonce    int
	}
)

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func (b *Block) Create(data string, prev []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prev, 0}
	// block.DeriveHash()

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (b *Block) Genesis() *Block {
	return b.Create("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	res := bytes.Buffer{}
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	HandleError(err)

	return res.Bytes()
}

func (b *Block) Deserialize(data []byte) *Block {
	block := &Block{}

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(block)
	HandleError(err)

	return block
}
