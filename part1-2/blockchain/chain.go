package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
)

// Ablockchain interface
type Ablockchain interface {
	PrintBlockChain()
	Init()
	GenerateBlocks(n int)
	//AddBlock(b Block)
	//AddBlock(b Tblock)
}

// BlockChain hold data for both light and full clients.
type BlockChain struct {
	LightClient []BlockHeader
	FullClient  []Block
	lastIndex   int
}

//TblockChain is a blockchain consisting of Tblocks
type TblockChain struct {
	LightClient []BlockHeader
	FullClient  []Tblock
	lastIndex   int
}

// PrintBlockChain prints all blocks in the blockchain
func (bc *BlockChain) PrintBlockChain() {
	for i := 0; i < len(bc.LightClient); i++ {
		//bc.FullClient[i].PrintBlock()
		bc.LightClient[i].PrintBlockHeader()
	}
}

// PrintBlockChain prints all blocks in the blockchain
func (bc *TblockChain) PrintBlockChain() {
	for i := 0; i < len(bc.LightClient); i++ {
		//bc.FullClient[i].PrintBlock()
		bc.LightClient[i].PrintBlockHeader()
	}
}

// Init initializes the blockchain with one block (genesis)
func (bc *BlockChain) Init() {
	// Generate genesis block
	genesisBlock := &Block{}
	// Initialize fields
	genesisBlock.Header.Index = 0
	genesisBlock.Data = rand.Int()
	genesisBlock.Header.Level = 65
	c := strconv.Itoa(genesisBlock.Data)
	h := sha256.New()
	h.Write([]byte(c))
	hashed := h.Sum(nil)
	genesisBlock.Header.Hash = hex.EncodeToString(hashed)
	bc.AddBlock(*genesisBlock)
}

// Init initializes the blockchain with one block (genesis)
func (bc *TblockChain) Init() {
	// Generate genesis block
	genesisBlock := &Tblock{}
	// Initialize fields
	genesisBlock.Header.Index = 0
	//genesisBlock.Data = rand.Int()
	sumData := 0
	for i := 0; i < len(genesisBlock.Data); i++ {
		genesisBlock.Data[i] = rand.Int()
		sumData += genesisBlock.Data[i]
	}
	genesisBlock.Header.Level = 65
	c := strconv.Itoa(sumData)
	h := sha256.New()
	h.Write([]byte(c))
	hashed := h.Sum(nil)
	genesisBlock.Header.Hash = hex.EncodeToString(hashed)
	bc.AddTblock(*genesisBlock)
}

// AddBlock adds a block to the blockchain
// assigns an index to that block
// assigns LevelPrevHash to that block
func (bc *BlockChain) AddBlock(b Block) {
	// append block to blockchain
	bc.LightClient = append(bc.LightClient, b.Header)
	bc.FullClient = append(bc.FullClient, b)
	// assign index to each block for light and full clients
	bc.LightClient[bc.lastIndex].Index = bc.lastIndex
	bc.FullClient[bc.lastIndex].Header.Index = bc.lastIndex
	// assign the previous hash of the higher level
	prevLevelIndex := FindPrevLevelBlockIndex(bc.LightClient, bc.lastIndex)
	pli := bc.LightClient[prevLevelIndex].Hash
	bc.LightClient[bc.lastIndex].LevelPrevHash = pli

	bc.lastIndex++
}

// AddTblock adds a block to the blockchain
// assigns an index to that block
// assigns LevelPrevHash to that block
func (bc *TblockChain) AddTblock(b Tblock) {
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
func (bc *BlockChain) GenerateBlocks(n int) {
	lastIndex := len(bc.LightClient) - 1
	lastBlock := &bc.FullClient[lastIndex]
	for i := 0; i < n; i++ {
		newBlock := &Block{}
		newBlock.Init(lastBlock)
		bc.AddBlock(*newBlock)
		lastBlock = newBlock
	}
}

// GenerateBlocks will take a number and add that number of blocks to the chain
func (bc *TblockChain) GenerateBlocks(n int) {
	lastIndex := len(bc.LightClient) - 1
	lastBlock := &bc.FullClient[lastIndex]
	for i := 0; i < n; i++ {
		newBlock := &Tblock{}
		newBlock.InitTblock(lastBlock)
		bc.AddTblock(*newBlock)
		lastBlock = newBlock
	}
}

// FindBlockByIndex will return a full block of a given index
func (bc *TblockChain) FindBlockByIndex(blockIndex int) *Tblock {
	for i := 0; i < bc.lastIndex; i++ {
		if blockIndex == bc.LightClient[i].Index {
			return &bc.FullClient[i]
		}
	}
	return nil
}
