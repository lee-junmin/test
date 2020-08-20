package sidechaintransfer

import (
	"github.com/lee-junmin/thesis-blockchain/blockchain"
)

// Transfer represents a data transformation case
type Transfer struct {
	Certificate string
	Data        int
}

// SideChain is chain that requests data from the main chain
type SideChain struct {
	chain             blockchain.TblockChain
	lockBox           Transfer
	addedTransactions []int
}

// MainChain is chain that requests data from the main chain
type MainChain struct {
	chain blockchain.TblockChain
	//lockBox TransferBox
}

type chain interface {
	Init(size int)
}

// Init will generate a sidechain
func (s *SideChain) Init(size int) {
	sidechain := blockchain.TblockChain{}
	sidechain.Init()
	sidechain.GenerateBlocks(size - 1)
	s.chain = sidechain
	//s.chain.PrintBlockChain()
}

// Init will generate a sidechain
func (m *MainChain) Init(size int) {
	mainchain := blockchain.TblockChain{}
	mainchain.Init()
	mainchain.GenerateBlocks(size - 1)
	m.chain = mainchain
	//m.chain.PrintBlockChain()
}

// ExecuteTransfer will
func (s *SideChain) ExecuteTransfer(source *MainChain, block int, transaction int) (bool, int) {
	// 0.initiate a transfer stuct
	sourcechain := source.chain
	sourceblock := sourcechain.FindBlockByIndex(block)
	data := sourceblock.Data[transaction]
	t := &Transfer{Data: data}
	// 1.generate certificate
	t.Certificate = GenerateCertificate(sourceblock, sourcechain.LightClient, block)
	// 2.locks the tranfer struct into the lock box of the sidechain
	s.lockBox = *t
	// 3.verify certificate
	isVerified := VerifyTransfer(sourcechain.LightClient, block, t)
	// 4.add transaction from mainchain to sidechain
	success := false
	if isVerified {
		s.AddTransaction(t.Data) // add transaction to sidechain
		//fmt.Println(s.addedTransactions)
		success = true
	}
	return success, t.Data
}

//AddTransaction will add the verified transaction to the sidechain
func (s *SideChain) AddTransaction(data int) {
	s.addedTransactions = append(s.addedTransactions, data)
}

//VerifyTransfer will return false if certificate doesn't match the requester's verification
func VerifyTransfer(lc []blockchain.BlockHeader, index int, t *Transfer) bool {
	// make a slice of hashes up that follow up to the genesis
	hashSlice := []string{lc[index].Hash}
	for index > 0 {
		prevLevelIndex := blockchain.FindPrevLevelBlockIndex(lc, index)
		if lc[index].LevelPrevHash == lc[prevLevelIndex].Hash {
			index = prevLevelIndex
			hashSlice = append(hashSlice, lc[prevLevelIndex].Hash)
		}
	}
	toHash := ""
	for _, v := range hashSlice {
		toHash += v
	}

	if blockchain.CreateHashFromString(toHash) == t.Certificate {
		//fmt.Println("match")
		return true
	}
	return false

}

// GenerateCertificate will take in light client data and a block index and return a hash made from the list of hashes that follow up to the genesis block
func GenerateCertificate(b *blockchain.Tblock, lc []blockchain.BlockHeader, index int) string {

	// make a slice of hashes up that follow up to the genesis
	hashSlice := []string{b.HashFromBlock()}
	for index > 0 {
		prevLevelIndex := blockchain.FindPrevLevelBlockIndex(lc, index)
		if lc[index].LevelPrevHash == lc[prevLevelIndex].Hash {
			index = prevLevelIndex
			hashSlice = append(hashSlice, lc[prevLevelIndex].Hash)
		}
	}

	// generate hash for hash slice
	toHash := ""
	for _, v := range hashSlice {
		toHash += v
	}
	//fmt.Println(toHash)
	return blockchain.CreateHashFromString(toHash)
}
