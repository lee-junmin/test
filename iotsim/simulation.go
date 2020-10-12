package iotsim

//TODO: need to sort the iotsim testing out
import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	//"time"
)

var SimulationClockOn bool
var SimulationClock time.Time
var Network bool
var NetworkFailRate int

//NeatTransferSim will run a simulation with no attacks
func NeatTransferSim(r int, sourceSize int) ([][]string, float32) {
	m := &IMainChain{}
	m.Init()
	test := &ISideChain{}
	test.Init(100)
	result := [][]string{}
	row := []string{"tranfer_success", "block_index", "transaction_index", "data-time", "data-temp", "data-humidity", "transfer_time"}
	result = append(result, row)

	successCount := 0
	//startTime := time.Now()
	for i := 0; i < r; i++ {
		//start transfer
		block := rand.Intn(sourceSize)
		//block := sourceSize - 1
		transaction := rand.Intn(100)
		//startTime := time.Now()
		s, d, t := test.ExecuteTransfer(m, block, transaction)
		//SimulationClock = SimulationClock.Add(time.Since(startTime))

		newrow := []string{strconv.FormatBool(s), strconv.Itoa(block), strconv.Itoa(transaction), d.temp, d.humidity, t.Format(time.RFC3339)}
		result = append(result, newrow)
		if s {
			successCount++
		} // end of trar := rand.Int
		randomInt := time.Duration(rand.Intn(10) + 100)
		SimulationClock = SimulationClock.Add(time.Minute * randomInt)
	}
	successRate := float32(successCount) / float32(r)
	fmt.Println("scv:", SimulationClock)
	return result, successRate
}

//NeatTransferSimSPV will run a simulation with no attacks
func NeatTransferSimSPV(r int, sourceSize int) ([][]string, float32) {
	m := &IMainChain{}
	m.Init()
	test := &ISideChain{}
	test.Init(10)
	result := [][]string{}
	row := []string{"tranfer_success", "block_index", "transaction_index", "data-time", "data-temp", "data-humidity", "transfer_time"}
	result = append(result, row)
	successCount := 0
	for i := 0; i < r; i++ {
		//	block := sourceSize - 1
		block := rand.Intn(sourceSize)
		transaction := rand.Intn(100)
		//startTime := time.Now()
		s, d, t := test.ExecuteTransferSPV(m, block, transaction)
		//SimulationClock = SimulationClock.Add(time.Since(startTime))

		newrow := []string{strconv.FormatBool(s), strconv.Itoa(block), strconv.Itoa(transaction), d.temp, d.humidity, t.Format(time.RFC3339)}
		result = append(result, newrow)
		if s {
			successCount++
		}

		randomInt := time.Duration(rand.Intn(10) + 100)
		SimulationClock = SimulationClock.Add(time.Minute * randomInt)
	}

	successRate := float32(successCount) / float32(r)
	fmt.Println("spc:", SimulationClock)
	return result, successRate
}

// ExportCSV into a csv file
func ExportCSV(title string, data [][]string) {
	file, err := os.Create(title)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, v := range data {
		err := writer.Write(v)
		checkError("Cannot write to file", err)
	}
}

//Check Error
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

//StartSimulationClock will start a global simulation clock, parameter will be the mintues per day for network failure
func StartSimulationClock() {
	SimulationClockOn = true
	//start := time.Now()
	SimulationClock = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	//for SimulationClockOn {
	//SimulationClock = SimulationClock.Add(time.Since(start))
	//start = SimulationClock
	//SimulationClock = SimulationClock.Add(time.Microsecond * 100)
	//time.Sleep(time.Microsecond)
	//fmt.Println(SimulationClock)
	//NetworkSimulation(NetworkFailRate)
	//NetworkSimulation()
	//}

}

// NetworkSimulation will mock a network failure
func NetworkSimulation() {

	if SimulationClock.Hour() == 1 && SimulationClock.Minute() <= NetworkFailRate-1 {
		//fmt.Println(SimulationClock)
		Network = false
	} else {
		Network = true
	}

}

//NetworkFailureTest will mock a network failure
func NetworkFailureTest(r int, sourceSize int, maxTime int) [][]string {

	result := [][]string{}
	row := []string{"failure min per day", "scv success rate", "spv success rate"}
	result = append(result, row)
	for i := 0; i < maxTime; i++ {
		SimulationClock = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		NetworkFailRate = i
		//NetworkSimulation()
		_, scvrate := NeatTransferSim(r, sourceSize)
		_, spvrate := NeatTransferSimSPV(r, sourceSize)
		newrow := []string{
			strconv.Itoa(i + 1),
			strconv.FormatFloat(float64(scvrate), 'f', -1, 32),
			strconv.FormatFloat(float64(spvrate), 'f', -1, 32),
		}
		result = append(result, newrow)
		fmt.Println("simluation", i, "min done...", scvrate, spvrate)

	}
	return result
}

func NeatRecordTransferSim(r int) ([][]string, float32) {
	m := &IMainChain{}
	m.Init()
	test := &ISideChain{}
	test.Init(100)
	result := [][]string{}
	row := []string{"product-id", "data-time", "data-temp", "data-humidity"}
	result = append(result, row)

	successCount := 0
	//startTime := time.Now()
	for i := 0; i < r; i++ {
		id := strconv.Itoa(rand.Intn(19734) + 100000)

		s, d, _ := test.ExecuteRecordTransfer(id, m)
		//SimulationClock = SimulationClock.Add(time.Since(startTime))
		for i := 0; i < len(s); i++ {
			newrow := []string{s[i], d[i].periodStart.Format("2006/01/02 15:04"), d[i].temp, d[i].humidity}
			result = append(result, newrow)

		}
		randomInt := time.Duration(rand.Intn(10) + 100)
		SimulationClock = SimulationClock.Add(time.Minute * randomInt)
	}
	successRate := float32(successCount) / float32(r)
	fmt.Println("scv:", SimulationClock)
	return result, successRate
}

func NormalityNetworkFailureTest(min int, r int, sourceSize int, rounds int) [][]string {

	result := [][]string{}
	row := []string{"sublinear", "linear"}
	result = append(result, row)
	for i := 0; i < rounds; i++ {
		SimulationClock = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		NetworkFailRate = min
		//NetworkSimulation()
		_, scvrate := NeatTransferSim(r, sourceSize)
		_, spvrate := NeatTransferSimSPV(r, sourceSize)
		newrow := []string{
			strconv.FormatFloat(float64(scvrate), 'G', -1, 32),
			strconv.FormatFloat(float64(spvrate), 'G', -1, 32),
		}
		result = append(result, newrow)
		fmt.Println("simluation", min, "min done...", scvrate, spvrate, 100*i/rounds, "%")
	}
	return result
}
