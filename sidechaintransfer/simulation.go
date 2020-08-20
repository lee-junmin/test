package sidechaintransfer

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

//NeatTransferSim will run a simulation with no attacks
func NeatTransferSim(r int, sourceSize int) [][]string {
	m := &MainChain{}
	m.Init(sourceSize)
	test := &SideChain{}
	test.Init(10)
	result := [][]string{}
	row := []string{"tranfer_success", "block_index", "transaction_index", "data"}
	result = append(result, row)
	for i := 0; i < r; i++ {
		block := rand.Intn(sourceSize)
		transaction := rand.Intn(100)
		s, d := test.ExecuteTransfer(m, block, transaction)
		newrow := []string{strconv.FormatBool(s), strconv.Itoa(block), strconv.Itoa(transaction), strconv.Itoa(d)}
		result = append(result, newrow)
	}
	return result
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
