package sublinearverification 


import (
	"time"
	//"fmt"
)

// SPVTime will verify a given block
// takes in a block index and RETURN true if the block is valid
func SPVTime(lc []BlockHeader, index int) int {
	start := time.Now()
	for index > 0 {

		if lc[index].PrevHash == lc[index-1].Hash {
			//time.Sleep(time.Nanosecond)
			index--
		} else {
			return 0
		}

	}
	//return fmt.Sprintf("%s", time.Since(start))
	return timeFormatMicro(time.Since(start))
}

func timeFormatMicro(t time.Duration) int {
	return int(t)
}

// SPVSteps takes in a slice of headers and RETURNs the number of steps for verification
func SPVSteps(lc []BlockHeader, index int) int {
	result:=0
	for index > 0 {

		if lc[index].PrevHash == lc[index-1].Hash {
		  result++
			index--
		} else {
			return 0
		}

	}

	return result
}
