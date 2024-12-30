package main

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"sync"
)

type Blockchain struct {
	Blocks     []*Block
	Difficulty int
	mu         sync.Mutex
}

func (bc *Blockchain) AddBlock(data string) *Block {

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := newBlock(prevBlock.Index+1, []byte(data), prevBlock.Hash, bc.Difficulty)

	if bc.isValidBlock(newBlock, prevBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
		return newBlock
	}
	return nil
}

func (bc *Blockchain) isValidBlock(newBlock, prevBlock *Block) bool {
	if newBlock.Index != prevBlock.Index+1 {
		log.Println("Invalid index")
		return false
	}

	if !bytes.Equal(newBlock.PrevBlockHash, prevBlock.Hash) {
		log.Println("Invalid prev hash")
		return false

	}
	if !bytes.Equal(newBlock.Hash, newBlock.CalculateHash()) {
		log.Println("Invalid hash")
		return false
	}

	if !bytes.HasPrefix(newBlock.Hash, bytes.Repeat([]byte{0}, bc.Difficulty)) {
		log.Println("Block does not meet difficulty requirements")
		return false
	}
	return true

}

func NewBlockchain(difficulty int) *Blockchain {
	block := CreateGenesisBlock(difficulty)
	return &Blockchain{
		Blocks:     []*Block{block},
		Difficulty: difficulty,
	}
}

func (bc *Blockchain) String() string {
	var builder strings.Builder
	builder.WriteString("BlockChain:\nDifficuilty: " + strconv.Itoa(bc.Difficulty) + "\n")
	builder.WriteString("Blocks:\n")
	for _, block := range bc.Blocks {
		builder.WriteString(block.String())
		builder.WriteString("\n")
	}
	return builder.String()
}
