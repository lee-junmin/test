package main

import (
	"github.com/lee-junmin/thesis-blockchain/iotsim"
	"github.com/lee-junmin/thesis-blockchain/sidechaintransfer"
	"github.com/lee-junmin/thesis-blockchain/sublinearverification"
)

func main() {
	//iotsim.ReadIot()
	IoTnetworktest()
	//m := &iotsim.IblockChain{}
	//m.Init()
	//m.GenerateBlocks(180)
	//m.Init(180)
	//fmt.Println(m.Chain.FullClient[0])
	//gethistogramoflevels(1000, 1000, 1)
	//normalityNetworkTest()
	// t := &blockchain.TblockChain{}
	// t.Init()
	// t.GenerateBlocks(1000)
	// t.PrintBlockChain()
}

// DONE part 1 basic comparison plotting
func sublintest(step int, max int, rep int) {
	sublinearverification.ExportCSV("./visualisations/scv-time.csv", sublinearverification.TestScvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/scv-step.csv", sublinearverification.TestScvStep(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/spv-time.csv", sublinearverification.TestSpvTime(step, max, rep))
	sublinearverification.ExportCSV("./visualisations/spv-step.csv", sublinearverification.TestSpvStep(step, max, rep))
}

func gethistogramoflevels(step int, max int, rep int) {
	sublinearverification.ExportCSV("./visualisations/scv-step-level.csv", sublinearverification.TestScvStep(step, max, rep))
}

func networktest() {
	sidechaintransfer.StartSimulationClock()
	sidechaintransfer.NetworkSimulation()
	data1, _ := sidechaintransfer.NeatTransferSim(5000, 1000)
	data2, _ := sidechaintransfer.NeatTransferSimSPV(5000, 1000)
	sidechaintransfer.ExportCSV("./visualisations/part2/scv-neat.csv", data1)
	sidechaintransfer.ExportCSV("./visualisations/part2/spv-neat.csv", data2)
	sidechaintransfer.NeatTransferSim(5000, 1000)
	sidechaintransfer.ExportCSV("./visualisations/part2/fail-test.csv", sidechaintransfer.NetworkFailureTest(1000, 100, 60))
}

func normalityNetworkTest() {
	sidechaintransfer.StartSimulationClock()
	sidechaintransfer.NetworkSimulation()
	// transfer 5000 times for random transaction of a blockchain size of 1000
	// data1, _ := sidechaintransfer.NeatTransferSim(5000, 1000)
	// data2, _ := sidechaintransfer.NeatTransferSimSPV(5000, 1000)
	// sidechaintransfer.ExportCSV("./visualisations/part2/normality.csv", data1)
	// sidechaintransfer.ExportCSV("./visualisations/part2/spv-neat.csv", data2)
	// sidechaintransfer.NeatTransferSim(5000, 1000)

	//sidechaintransfer.ExportCSV("./visualisations/part2/stat-test-15.csv", sidechaintransfer.NormalityNetworkFailureTest(15, 1500, 300, 1000))
	//sidechaintransfer.ExportCSV("./visualisations/part2/stat-test-30.csv", sidechaintransfer.NormalityNetworkFailureTest(30, 1500, 300, 1000))
	sidechaintransfer.ExportCSV("./visualisations/part2/stat-test-15-5000.csv", sidechaintransfer.NormalityNetworkFailureTest(15, 5000, 200, 500))
	sidechaintransfer.ExportCSV("./visualisations/part2/stat-test-30-5000.csv", sidechaintransfer.NormalityNetworkFailureTest(30, 5000, 200, 500))

}

//part 3

func IoTnetworktest() {
	iotsim.ReadIot()
	iotsim.StartSimulationClock()
	iotsim.NetworkSimulation()
	//data1, _ := iotsim.NeatTransferSim(5000, 180)
	//data2, _ := iotsim.NeatTransferSimSPV(5000, 180)

	//fmt.Println("START=====length", len(iotsim.DATA))
	//data3, _ := iotsim.NeatRecordTransferSim(100)
	//	iotsim.ExportCSV("./visualisations/part3/scv-fail.csv", data1)
	//	iotsim.ExportCSV("./visualisations/part3/spv-fail.csv", data2)
	//iotsim.ExportCSV("./visualisations/part3/iot.csv", data3)
	//sidechaintransfer.NeatTransferSim(5000, 1000)
	//iotsim.ExportCSV("./visualisations/part3/fail-test-15.csv", iotsim.NetworkFailureTest(5000, 180, 15))
	//iotsim.ReadStorage()
	iotsim.ExportCSV("./visualisations/part3/stat-test-15-5000.csv", iotsim.NormalityNetworkFailureTest(15, 5000, 180, 500))
}
