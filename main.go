package main

import (
	"github.com/lee-junmin/thesis-blockchain/sidechaintransfer"
	"github.com/lee-junmin/thesis-blockchain/sublinearverification"
)

func main() {
	// test := [][]string{}
	// testRow1 := []string{"a", "b", "c", "d"}
	// testRow2 := []string{"j", "df", "asd", "ggg"}
	// test = append(test, testRow1)
	// test = append(test, testRow2)
	sidechaintransfer.ExportCSV("./visualisations/test.csv", sidechaintransfer.NeatTransferSim(50000, 1000))

}

func sublintest(step int, max int, rep int) {
	sublinearverification.ExportCSV("./visualisations/scv-time.csv", sublinearverification.TestScvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/scv-step.csv", sublinearverification.TestScvStep(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/spv-time.csv", sublinearverification.TestSpvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/spv-step.csv", sublinearverification.TestSpvStep(step, max, rep))
}
