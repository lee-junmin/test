package iotsim

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

var DATA []IotRow

type Iblock struct {
	Header BlockHeader
	Data   []IotRow
}

// BlockHeader represents each header of the blocks in the blockchain
type BlockHeader struct {
	Index         int
	Level         int
	Hash          string
	PrevHash      string
	LevelPrevHash string
}

//IblockChain is a blockchain consisting of Tblocks
type IblockChain struct {
	LightClient []BlockHeader
	FullClient  []Iblock
	lastIndex   int
}

//InitIblock initializes a block that has transactions
func (b *Iblock) InitIblock(prevBlock *Iblock) {
	//fmt.Println("DATA length", len(DATA))
	sumData := ""
	b.Data = make([]IotRow, 100, 100)
	for i := 0; i < 100; i++ {
		//fmt.Println(DATA[i])
		b.Data[i] = DATA[i]

		rowstring := DATA[i].periodStart.String() + DATA[i].temp + DATA[i].humidity
		sumData = sumData + rowstring

	}

	DATA = DATA[100:]

	b.Header.PrevHash = prevBlock.Header.Hash
	b.Header.Hash = b.HashFromBlock()
	b.Header.Level = strings.Count(b.Header.Hash, "0")
}

// PrintBlock will print all the fields in the block
func (b *Iblock) PrintIBlock() {
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
func (b *BlockHeader) PrintIBlockHeader() {
	fmt.Println("Index", b.Index, "==========")
	fmt.Println("Level:", b.Level)
	fmt.Println("Hash:", b.Hash)
	fmt.Println("PrevHash:", b.PrevHash)
	fmt.Println("PrevLevelHash:", b.LevelPrevHash)

	//fmt.Printf("%d,%d\n", b.Index, b.Level)
}

// PrintBlockChain prints all blocks in the blockchain
func (bc *IblockChain) PrintIBlockChain() {
	for i := 0; i < len(bc.LightClient); i++ {
		//bc.LightClient[i].PrintIBlockHeader()
	}
}

// // Init initializes the blockchain with one block (genesis)
func (bc *IblockChain) Init() {
	// Generate genesis block
	genesisBlock := &Iblock{}
	// Initialize fields
	genesisBlock.Header.Index = 0
	genesisBlock.Data = make([]IotRow, 100, 100)

	//genesisBlock.Data = rand.Int()
	sumData := ""
	for i := 0; i < 100; i++ {
		genesisBlock.Data[i] = DATA[i]
		rowstring := DATA[i].periodStart.String() + DATA[i].temp + DATA[i].humidity
		sumData = sumData + rowstring

	}
	DATA = DATA[100:]
	genesisBlock.Header.Level = 65
	c := sumData
	h := sha256.New()
	h.Write([]byte(c))
	hashed := h.Sum(nil)
	genesisBlock.Header.Hash = hex.EncodeToString(hashed)
	bc.AddIblock(*genesisBlock)
}

// AddIblock adds a block to the blockchain
// assigns an index to that block
// assigns LevelPrevHash to that block
func (bc *IblockChain) AddIblock(b Iblock) {
	// append block to blockchain
	bc.LightClient = append(bc.LightClient, b.Header)
	bc.FullClient = append(bc.FullClient, b)
	// assign index to each block for light and full clients
	bc.LightClient[bc.lastIndex].Index = bc.lastIndex
	bc.FullClient[bc.lastIndex].Header.Index = bc.lastIndex
	// assign the previous hash of the higher level
	prevLevelIndex := FindPrevLevelBlockIndex(bc.LightClient, bc.lastIndex)
	//fmt.Println(prevLevelIndex)
	pli := bc.LightClient[prevLevelIndex].Hash
	bc.LightClient[bc.lastIndex].LevelPrevHash = pli

	//fmt.Println(FindPrevLevelBlockIndex(bc.LightClient, bc.lastIndex))
	bc.lastIndex++
	//bc.FindBlockByIndex(bc.lastIndex).Header.LevelPrevHash = pli
}

// GenerateBlocks will take a number and add that number of blocks to the chain
func (bc *IblockChain) GenerateBlocks(n int) {
	lastIndex := len(bc.LightClient) - 1
	lastBlock := &bc.FullClient[lastIndex]
	for i := 0; i < n; i++ {
		newBlock := &Iblock{}
		newBlock.InitIblock(lastBlock)
		bc.AddIblock(*newBlock)
		lastBlock = newBlock
		//fmt.Println("Generate", i)
	}
}

// FindBlockByIndex will return a full block of a given index
func (bc *IblockChain) FindBlockByIndex(blockIndex int) *Iblock {
	for i := 0; i < bc.lastIndex; i++ {
		if blockIndex == bc.LightClient[i].Index {
			return &bc.FullClient[i]
		}
	}
	return nil
}

// HashFromBlock will recalculate the hash and RETURN the hash string
func (b *Iblock) HashFromBlock() string {

	sumData := ""
	for i := 0; i < 100; i++ {
		rowstring := b.Data[i].periodStart.String() + b.Data[i].temp + b.Data[i].humidity
		sumData = sumData + rowstring
	}

	c := sumData + b.Header.PrevHash //+ strconv.Itoa(b.Header.Index)
	result := CreateHashFromString(c)
	return result
}

// CreateHashFromString will create a hash from a string
func CreateHashFromString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	result := hex.EncodeToString(hashed)
	return result
}

// FindPrevLevelBlockIndex will take in a block index and
// RETURN the block index of the closest higher level block
func FindPrevLevelBlockIndex(lc []BlockHeader, n int) int {
	if n < 2 {
		return 0
	}
	result := n - 1
	for lc[n].Level > lc[result].Level {
		result--
	}
	return result
}
