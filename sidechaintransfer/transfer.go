package sidechaintransfer

import "github.com/lee-junmin/thesis-blockchain/blockchain"

// Transfer represents a data transformation case
type Transfer struct {
	Certificate []string
	Data        int
}



// TransferRequest from the sender
func (requester *blockchain.Tblock) TransferRequest(source *TblockChain, bi int, ti int) {
	transfer := &Transfer{}
	transfer.Certificate = transfer.GenerateCertificate()
	transfer.Data = source.FindBlockByIndex(bi).Data[ti]
	//sender.Verify(tranfer)
	Println("transfer requested...")

}

// Lock will add the input transfer the requester's lockbox
func (requester *blockchain.Tblock) Lock(t *Transfer) {
	cert := t.Certificate
	data := t.Data
	transferBox := &TransferBox{Certificate:cert,Data:data}
	append(requester.LockBox,transferBox)
}

// GenerateCertificate will generate a slice of superblocks
func (t *Transfer) GenerateCertificate() []string {
	var result []string
    index:=  
	for index > 0 {
		prevLevelIndex := blockchain.FindPrevLevelBlockIndex(lc, index)
		if lc[index].LevelPrevHash == lc[prevLevelIndex].Hash {
			//fmt.Println("scv-path", index, lc[prevLevelIndex].Index, lc[prevLevelIndex].Level)
			index = prevLevelIndex
		} else {
			return 0
		}
	}
}

// Verify
func (sender *Tblock) Verify(t *Transfer) bool {
	// recalculate the hash
	t.
	// follow up to the genesis block

}

// Unlock will execute the tranfer to the sender blockchain
func (t *Transfer) Unlock(sender *Tblock) {

}
