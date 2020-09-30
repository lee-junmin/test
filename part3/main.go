package main

func main() {

	transfersim.readCSV()

	// test := [][]string{}
	// testRow1 := []string{"a", "b", "c", "d"}
	// testRow2 := []string{"j", "df", "asd", "ggg"}
	// test = append(test, testRow1)
	// test = append(test, testRow2)

	//sidechaintransfer.StartSimulationClock()
	//NetworkSimulation(1)
	//data, _ := sidechaintransfer.NeatTransferSim(5000, 1000)
	//sidechaintransfer.ExportCSV("./visualisations/scv-fail.csv", data)
	//sidechaintransfer.ExportCSV("./visualisations/spv-fail.csv", sidechaintransfer.NeatTransferSimSPV(5000, 1000))
	//sidechaintransfer.NeatTransferSim(5000, 1000)

	//sidechaintransfer.ExportCSV("./visualisations/fail.csv", sidechaintransfer.NetworkFailureTest(1000, 100, 60))

}

// func sublintest(step int, max int, rep int) {
// 	sublinearverification.ExportCSV("./visualisations/scv-time.csv", sublinearverification.TestScvTime(step, max, rep))
// 	sublinearverification.ExportCSV("./visualisations/scv-step.csv", sublinearverification.TestScvStep(step, max, rep))
// 	sublinearverification.ExportCSV("./visualisations/spv-time.csv", sublinearverification.TestSpvTime(step, max, rep))
// 	sublinearverification.ExportCSV("./visualisations/spv-step.csv", sublinearverification.TestSpvStep(step, max, rep))
//}
