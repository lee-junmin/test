package iotsim

import (
	"fmt"
	"time"
)

var RECORD []StorageRow

// ITransfer represents a data transformation case
type ITransfer struct {
	Certificate string
	Data        IotRow
}

// ISideChain is chain that requests data from the main chain
type ISideChain struct {
	Chain             IblockChain
	lockBox           ITransfer
	addedTransactions []IotRow
}

// IMainChain is chain that requests data from the main chain
type IMainChain struct {
	Chain IblockChain
	//lockBox TransferBox
}

//TODO: init sidechain with random data
// Init will generate a sidechain
func (s *ISideChain) Init(size int) {

	sidechain := IblockChain{}
	ReadIot()
	sidechain.Init()
	sidechain.GenerateBlocks(size - 1)
	s.Chain = sidechain
	//s.chain.PrintBlockChain()
}

// Init will generate a mainchain
func (m *IMainChain) Init() {
	//problem
	ReadIot()
	mainchain := IblockChain{}
	mainchain.Init()
	mainchain.GenerateBlocks(180)
	m.Chain = mainchain
	m.Chain.PrintIBlockChain()
}

func (s *ISideChain) ExecuteRecordTransfer(id string, source *IMainChain) ([]string, []IotRow, []time.Time) {

	bo := []string{}
	io := []IotRow{}
	ti := []time.Time{}
	ReadStorage()

	var in, out time.Time
	for _, v := range RECORD {
		if id == v.productId {
			in = v.in
			out = v.out
			fmt.Println("product ID: ", id)
		}
	}

	for i := 0; i < 180; i++ {
		for j := 0; j < 100; j++ {
			pStart := source.Chain.FullClient[i].Data[j].periodStart
			inTen := in.Add(time.Minute * (-10))
			if inTen.Before(pStart) && out.After(pStart) {
				_, r2, r3 := s.ExecuteTransfer(source, i, j)
				bo = append(bo, id)
				io = append(io, r2)
				ti = append(ti, r3)

			}
		}
	}

	return bo, io, ti

}

// ExecuteTransfer will
func (s *ISideChain) ExecuteTransfer(source *IMainChain, block int, transaction int) (bool, IotRow, time.Time) {
	// 0.initiate a transfer stuct
	//fmt.Println("start transfer")
	sourcechain := source.Chain
	sourceblock := sourcechain.FindBlockByIndex(block)
	data := sourceblock.Data[transaction]
	t := &ITransfer{Data: data}
	// 1.generate certificate
	t.Certificate = GenerateCertificate(sourceblock, sourcechain.LightClient, block)
	// 2.locks the tranfer struct into the lock box of the sidechain
	s.lockBox = *t
	// 3.verify certificate
	isVerified := VerifyTransfer(sourcechain.LightClient, block, t)
	// 4.add transaction from mainchain to sidechain
	success := false
	if isVerified && Network {
		s.AddTransaction(t.Data) // add transaction to sidechain
		//fmt.Println(s.addedTransactions)
		success = true
	}
	return success, t.Data, SimulationClock
}

// ExecuteTransferSPV will
func (s *ISideChain) ExecuteTransferSPV(source *IMainChain, block int, transaction int) (bool, IotRow, time.Time) {

	// 0.initiate a transfer stuct
	sourcechain := source.Chain
	sourceblock := sourcechain.FindBlockByIndex(block)
	data := sourceblock.Data[transaction]
	t := &ITransfer{Data: data}
	// 1.generate certificate
	t.Certificate = GenerateCertificateSPV(sourceblock, sourcechain.LightClient, block)
	// 2.locks the tranfer struct into the lock box of the sidechain
	s.lockBox = *t
	// 3.verify certificate
	isVerified := VerifyTransferSPV(sourcechain.LightClient, block, t)
	// 4.add transaction from mainchain to sidechain
	success := false
	if isVerified {
		s.AddTransaction(t.Data) // add transaction to sidechain
		//fmt.Println(s.addedTransactions)
		success = true
	}
	//fmt.Println(timeFormatMicro(time.Since(startTime))
	return success, t.Data, SimulationClock
}

//AddTransaction will add the verified transaction to the sidechain
func (s *ISideChain) AddTransaction(data IotRow) {
	s.addedTransactions = append(s.addedTransactions, data)
	//fmt.Println("added")
}

//VerifyTransfer will return false if certificate doesn't match the requester's verification
func VerifyTransfer(lc []BlockHeader, index int, t *ITransfer) bool {
	// make a slice of hashes up that follow up to the genesis
	hashSlice := []string{lc[index].Hash}
	for index > 0 {
		start := time.Now()
		prevLevelIndex := FindPrevLevelBlockIndex(lc, index)
		if lc[index].LevelPrevHash == lc[prevLevelIndex].Hash {
			index = prevLevelIndex
			hashSlice = append(hashSlice, lc[prevLevelIndex].Hash)
		}
		SimulationClock.Add(time.Since(start))
		NetworkSimulation()
		if Network == false {
			return false
		}
	}
	toHash := ""
	for _, v := range hashSlice {
		toHash += v
	}

	if CreateHashFromString(toHash) == t.Certificate {

		return true
	}
	return false

}

// GenerateCertificate will take in light client data and a block index and return a hash made from the list of hashes that follow up to the genesis block
func GenerateCertificate(b *Iblock, lc []BlockHeader, index int) string {

	// make a slice of hashes up that follow up to the genesis
	hashSlice := []string{b.HashFromBlock()}
	for index > 0 {
		prevLevelIndex := FindPrevLevelBlockIndex(lc, index)
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
	return CreateHashFromString(toHash)
}

//VerifyTransferSPV will return false if certificate doesn't match the requester's verification
func VerifyTransferSPV(lc []BlockHeader, index int, t *ITransfer) bool {
	// make a slice of hashes up that follow up to the genesis
	hashSlice := []string{lc[index].Hash}
	for index > 0 {
		start := time.Now()
		if lc[index].PrevHash == lc[index-1].Hash {
			hashSlice = append(hashSlice, lc[index-1].Hash)
			index--
		}
		SimulationClock.Add(time.Since(start))
		NetworkSimulation()
		if Network == false {
			return false
		}
	}
	toHash := ""
	for _, v := range hashSlice {
		toHash += v
	}

	if CreateHashFromString(toHash) == t.Certificate {
		//fmt.Println("match")
		return true
	}
	return false

}

// GenerateCertificateSPV will take in light client data and a block index and return a hash made from the list of hashes that follow up to the genesis block
func GenerateCertificateSPV(b *Iblock, lc []BlockHeader, index int) string {

	// make a slice of hashes up that follow up to the genesis
	hashSlice := []string{b.HashFromBlock()}
	for index > 0 {
		if lc[index].PrevHash == lc[index-1].Hash {
			hashSlice = append(hashSlice, lc[index-1].Hash)
			index--
		}
	}

	// generate hash for hash slice
	toHash := ""
	for _, v := range hashSlice {
		toHash += v
	}
	//fmt.Println(toHash)
	return CreateHashFromString(toHash)
}
