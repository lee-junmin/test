package main

import (
	"github.com/lee-junmin/thesis-blockchain/blockchain"
	"github.com/lee-junmin/thesis-blockchain/sublinearverification"
)

func main() {
	// //sublintest(100, 10000, 100)
	// testBlock := blockchain.Tblock{}
	// fmt.Println(testBlock)
	// //testBlock.InitTblock()
	// testBlock.PrintBlock()

	test := &blockchain.TblockChain{}
	test.Init()
	test.PrintBlockChain()

}

func sublintest(step int, max int, rep int) {
	sublinearverification.ExportCSV("./visualisations/scv-time.csv", sublinearverification.TestScvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/scv-step.csv", sublinearverification.TestScvStep(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/spv-time.csv", sublinearverification.TestSpvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/spv-step.csv", sublinearverification.TestSpvStep(step, max, rep))
}
