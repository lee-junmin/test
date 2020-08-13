package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"

	//"math/rand"
	"strconv"
	"strings"
)

var r = rand.New(rand.NewSource(99))

type Ablock interface {
	PrintBlock()
}

// BlockHeader represents each header of the blocks in the blockchain
type BlockHeader struct {
	Index         int
	Level         int
	Hash          string
	PrevHash      string
	LevelPrevHash string
}

// Block represents a block that will form a blockchain
type Block struct {
	Header BlockHeader
	Data   int
}

// TBlock is a block but has  transactions
type Tblock struct {
	Header BlockHeader
	Data   [100]int
}

// Init initializes a block,  calculates the hash and level of the block and store radom data
func (b *Block) Init(prevBlock *Block) {
	b.Data = r.Int()
	b.Header.PrevHash = prevBlock.Header.Hash
	c := strconv.Itoa(b.Data) + b.Header.PrevHash + strconv.Itoa(b.Header.Index) + strconv.Itoa(b.Header.Level) + b.Header.LevelPrevHash
	h := sha256.New()
	h.Write([]byte(c))
	hashed := h.Sum(nil)
	b.Header.Hash = hex.EncodeToString(hashed)
	b.Header.Level = strings.Count(b.Header.Hash, "0") //+ strings.Count(b.Header.Hash, "1") 	//+ strings.Count(b.Header.Hash, "2") 	+ strings.Count(b.Header.Hash, "3") 	+ strings.Count(b.Header.Hash, "4") 	+ strings.Count(b.Header.Hash, "5")
}

//InitTblock initializes a block that has transactions
func (b *Tblock) InitTblock(prevBlock *Tblock) {
	dataSum := 0
	for i := 0; i < len(b.Data); i++ {
		b.Data[i] = rand.Int()
		dataSum += b.Data[i]
	}

	b.Header.PrevHash = prevBlock.Header.Hash

	c := strconv.Itoa(dataSum) + b.Header.PrevHash + strconv.Itoa(b.Header.Index) + strconv.Itoa(b.Header.Level) + b.Header.LevelPrevHash
	h := sha256.New()
	h.Write([]byte(c))
	hashed := h.Sum(nil)
	b.Header.Hash = hex.EncodeToString(hashed)
	b.Header.Level = strings.Count(b.Header.Hash, "0") //+ strings.Count(b.Header.Hash, "1") 	//+ strings.Count(b.Header.Hash, "2") 	+ strings.Count(b.Header.Hash, "3") 	+ strings.Count(b.Header.Hash, "4") 	+ strings.Count(b.Header.Hash, "5")
}

// PrintBlock will print all the fields in the block
func (b *Block) PrintBlock() {
	fmt.Println("(full)========== INDEX", b.Header.Index, "==========")
	fmt.Println("Level        :", b.Header.Level)
	fmt.Println("Hash         :", b.Header.Hash)
	fmt.Println("PrevHash     :", b.Header.PrevHash)
	fmt.Println("LevelPrevHash:", b.Header.LevelPrevHash)
	fmt.Println("Data         :", b.Data)
}

// PrintBlock will print all the fields in the block
func (b *Tblock) PrintBlock() {
	fmt.Println("(full)========== INDEX", b.Header.Index, "==========")
	fmt.Println("Level        :", b.Header.Level)
	fmt.Println("Hash         :", b.Header.Hash)
	fmt.Println("PrevHash     :", b.Header.PrevHash)
	fmt.Println("LevelPrevHash:", b.Header.LevelPrevHash)
	fmt.Println("Data")
	for i := 0; i < len(b.Data); i++ {
		fmt.Print(b.Data[i], " ")
	}
	fmt.Println("")

}

// PrintBlockHeader will print all the files in a block header
func (b *BlockHeader) PrintBlockHeader() {
	//fmt.Println("INDEX", b.Index, "LEVEL", b.Level, "HASH", b.Hash[:3], "L.HASH", b.LevelPrevHash[:3])
	//fmt.Println("Index", b.Index, "==========")
	//fmt.Println("Level:", b.Level)
	//fmt.Println("Hash:", b.Hash)
	//fmt.Println("PrevHash:", b.PrevHash)
	//fmt.Println("PrevLevelHash", b.LevelPrevHash)
	fmt.Println("Index", b.Index, "Level:", b.Level)
}
