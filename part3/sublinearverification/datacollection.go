package sublinearverification

import (
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lee-junmin/thesis-blockchain/blockchain"
)

func TestScvTime(step int, max int, rep int) [][]string {
	result := [][]string{}
	result = append(result, []string{"Sublinear Complexity Verification (time)"})
	//TODO: Generate rep amount of blockchains
	blockChainSlice := make([]blockchain.BlockChain, 0, rep)

	for i := 0; i < rep; i++ {
		newBlockChain := blockchain.BlockChain{}
		newBlockChain.Init()
		newBlockChain.GenerateBlocks(max)
		blockChainSlice = append(blockChainSlice, newBlockChain)
	}

	//TODO: Iterate steps and get the average of steps of each blockchain
	iter := int(max / step)
	for i := 0; i < iter; i++ {
		s := ""
		sum := 0
		for j := 0; j < rep; j++ {
			//fmt.Println("slice",blockChainSlice[j].lastIndex)
			sum += SCVTime(blockChainSlice[j].LightClient, i*step)
		}
		s = strconv.Itoa(sum / rep)
		row := []string{s}
		result = append(result, row)
	}
	return result
}

//TestScvStep will return a 2D slice holding number of steps
func TestScvStep(step int, max int, rep int) [][]string {
	result := [][]string{}
	result = append(result, []string{"Sublinear Complexity Verification (step)"})
	//TODO: Generate rep amount of blockchains
	blockChainSlice := make([]blockchain.BlockChain, 0, rep)

	for i := 0; i < rep; i++ {
		newBlockChain := blockchain.BlockChain{}
		newBlockChain.Init()
		newBlockChain.GenerateBlocks(max)
		blockChainSlice = append(blockChainSlice, newBlockChain)
	}

	//TODO: Iterate steps and get the average of steps of each blockchain
	iter := int(max / step)
	for i := 0; i < iter; i++ {
		s := ""
		sum := 0
		for j := 0; j < rep; j++ {
			sum += SCVSteps(blockChainSlice[j].LightClient, i*step)
		}
		s = strconv.Itoa(sum / rep)
		row := []string{s}
		result = append(result, row)
	}
	return result
}

func TestSpvTime(step int, max int, rep int) [][]string {
	result := [][]string{}
	result = append(result, []string{"Sublinear Complexity Verification (time)"})
	//TODO: Generate rep amount of blockchains
	blockChainSlice := make([]blockchain.BlockChain, 0, rep)

	for i := 0; i < rep; i++ {
		newBlockChain := blockchain.BlockChain{}
		newBlockChain.Init()
		newBlockChain.GenerateBlocks(max)
		blockChainSlice = append(blockChainSlice, newBlockChain)
	}

	//TODO: Iterate steps and get the average of steps of each blockchain
	iter := int(max / step)
	for i := 0; i < iter; i++ {
		s := ""
		sum := 0
		for j := 0; j < rep; j++ {
			//fmt.Println("slice",blockChainSlice[j].lastIndex)
			sum += SPVTime(blockChainSlice[j].LightClient, i*step)
		}
		s = strconv.Itoa(sum / rep)
		row := []string{s}
		result = append(result, row)
	}
	return result
}

//TestScvStep will return a 2D slice holding number of steps
func TestSpvStep(step int, max int, rep int) [][]string {
	result := [][]string{}
	result = append(result, []string{"Sublinear Complexity Verification (step)"})
	//TODO: Generate rep amount of blockchains
	blockChainSlice := make([]blockchain.BlockChain, 0, rep)

	for i := 0; i < rep; i++ {
		newBlockChain := blockchain.BlockChain{}
		newBlockChain.Init()
		newBlockChain.GenerateBlocks(max)
		blockChainSlice = append(blockChainSlice, newBlockChain)
	}

	//TODO: Iterate steps and get the average of steps of each blockchain
	iter := int(max / step)
	for i := 0; i < iter; i++ {
		s := ""
		sum := 0
		for j := 0; j < rep; j++ {
			sum += SPVSteps(blockChainSlice[j].LightClient, i*step)
		}
		s = strconv.Itoa(sum / rep)
		row := []string{s}
		result = append(result, row)
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
