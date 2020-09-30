package sublinearverification

import (
	"fmt"
	"time"

	"github.com/lee-junmin/thesis-blockchain/part1-2/blockchain"
)

// SCVTime will verify a given block based on the level
// takes in a block index and RETURN true if the block is valid
func SCVTime(lc []blockchain.BlockHeader, index int) int {
	//fmt.Println("sublinear complexity verification called")
	start := time.Now()
	for index > 0 {
		prevLevelIndex := blockchain.FindPrevLevelBlockIndex(lc, index)
		if lc[index].LevelPrevHash == lc[prevLevelIndex].Hash {
			//fmt.Println("scv-path", index, lc[prevLevelIndex].Index, lc[prevLevelIndex].Level)
			index = prevLevelIndex
		} else {
			return 0
		}
	}
	return timeFormatMicro(time.Since(start))
}

// SCVSteps will verify a given block based on the level
// takes in a block index and RETURN true if the block is valid
func SCVSteps(lc []blockchain.BlockHeader, index int) int {
	fmt.Printf("\n")
	results := 0
	for index > 0 {
		prevLevelIndex := blockchain.FindPrevLevelBlockIndex(lc, index)
		if lc[index].LevelPrevHash == lc[prevLevelIndex].Hash {
			//fmt.Println("scv-path", index, lc[prevLevelIndex].Index, lc[prevLevelIndex].Level)
			//fmt.Printf("%d ",index)
			index = prevLevelIndex
			results++
		} else {
			return 0
		}
	}
	//fmt.Println("steps:",results)
	return results
}
