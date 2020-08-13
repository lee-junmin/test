package main

import (
	"math/rand"

	"github.com/lee-junmin/thesis-blockchain/sublinearverification"
)

var r = rand.New(rand.NewSource(99))

func main() {
	test(100, 10000, 100)
}

func test(step int, max int, rep int) {
	sublinearverification.ExportCSV("./visualisations/scv-time.csv", TestScvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/scv-step.csv", TestScvStep(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/spv-time.csv", TestSpvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/spv-step.csv", TestSpvStep(step, max, rep))
}
