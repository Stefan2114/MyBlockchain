package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index         int
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) CalculateHash() []byte {
	var data bytes.Buffer
	binary.Write(&data, binary.LittleEndian, int64(b.Index))
	binary.Write(&data, binary.LittleEndian, b.Timestamp)
	data.Write(b.Data)
	data.Write(b.PrevBlockHash)
	binary.Write(&data, binary.LittleEndian, int64(b.Nonce))
	hash := sha256.Sum256(data.Bytes())
	return hash[:]
}

func (b *Block) MineBlock(difficulty int) {
	target := bytes.Repeat([]byte{0}, difficulty)
	for {
		b.Hash = b.CalculateHash()
		if bytes.HasPrefix(b.Hash, target) {
			break
		}
		b.Nonce++
	}
}

func newBlock(index int, data []byte, prevBlockHash []byte, difficulty int) *Block {
	block := &Block{
		Index:         index,
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: prevBlockHash,
		Nonce:         0,
	}
	block.MineBlock(difficulty)
	return block
}

func CreateGenesisBlock(difficulty int) *Block {
	block := &Block{Index: 0,
		Timestamp:     time.Now().Unix(),
		Data:          []byte("Genesis block"),
		PrevBlockHash: []byte{},
		Nonce:         0,
	}
	block.MineBlock(difficulty)
	return block
}

func (b *Block) String() string {
	var builder strings.Builder
	builder.WriteString("Index: " + strconv.Itoa(b.Index) + "\n")
	builder.WriteString("Timestamp: " + strconv.FormatInt(b.Timestamp, 10) + "\n")
	builder.WriteString("Data: ")
	builder.Write(b.Data)
	builder.WriteString("\nPrevBlockHash: " + fmt.Sprintf("%x", b.PrevBlockHash) + "\n")
	builder.WriteString("Hash: " + fmt.Sprintf("%x", b.Hash) + "\n")
	builder.WriteString("Nonce: " + strconv.Itoa(b.Nonce))
	return builder.String()

}
