package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	//join two slice of bytes
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

type BlockChain struct {
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{GenesisBlock()}}
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis", []byte{})
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First blockchain after genesis")
	chain.AddBlock("second blockchain after genesis")
	chain.AddBlock("third block after genesis block")

	for key, val := range chain.blocks {
		fmt.Printf("the number %v  is %x \n", key, val)
	}
}
