package main

import (
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"strconv"
)

func TestScvTime(step int, max int, rep int) [][]string{
	result := [][]string{}
	result = append(result, []string{"Sublinear Complexity Verification (time)"})
	//TODO: Generate rep amount of blockchains
	blockChainSlice := make([]BlockChain,0,rep)

	for i:=0;i<rep;i++ {
		newBlockChain := BlockChain{}
		newBlockChain.Init()
		newBlockChain.GenerateBlocks(max)
		blockChainSlice = append(blockChainSlice,newBlockChain)    
	}

	//TODO: Iterate steps and get the average of steps of each blockchain
	iter := int(max / step)
	for i:=0;i<iter;i++{
		s := ""
		sum := 0
		for j:=0;j<rep;j++{
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
func TestScvStep(step int, max int, rep int) [][]string{
	result := [][]string{}
	result = append(result, []string{"Sublinear Complexity Verification (step)"})
	//TODO: Generate rep amount of blockchains
	blockChainSlice := make([]BlockChain,0,rep)

	for i:=0;i<rep;i++ {
		newBlockChain := BlockChain{}
		newBlockChain.Init()
		newBlockChain.GenerateBlocks(max)
		blockChainSlice = append(blockChainSlice,newBlockChain)    
	}

	//TODO: Iterate steps and get the average of steps of each blockchain
	iter := int(max / step)
	for i:=0;i<iter;i++{
		s := ""
		sum := 0
		for j:=0;j<rep;j++{
			sum += SCVSteps(blockChainSlice[j].LightClient, i*step)
		}
		s = strconv.Itoa(sum / rep)
		row := []string{s}
		result = append(result, row)
	}
	return result
}



func TestSpvTime(step int, max int, rep int) [][]string{
	result := [][]string{}
	result = append(result, []string{"Sublinear Complexity Verification (time)"})
	//TODO: Generate rep amount of blockchains
	blockChainSlice := make([]BlockChain,0,rep)

	for i:=0;i<rep;i++ {
		newBlockChain := BlockChain{}
		newBlockChain.Init()
		newBlockChain.GenerateBlocks(max)
		blockChainSlice = append(blockChainSlice,newBlockChain)    
	}

	//TODO: Iterate steps and get the average of steps of each blockchain
	iter := int(max / step)
	for i:=0;i<iter;i++{
		s := ""
		sum := 0
		for j:=0;j<rep;j++{
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
func TestSpvStep(step int, max int, rep int) [][]string{
	result := [][]string{}
	result = append(result, []string{"Sublinear Complexity Verification (step)"})
	//TODO: Generate rep amount of blockchains
	blockChainSlice := make([]BlockChain,0,rep)

	for i:=0;i<rep;i++ {
		newBlockChain := BlockChain{}
		newBlockChain.Init()
		newBlockChain.GenerateBlocks(max)
		blockChainSlice = append(blockChainSlice,newBlockChain)    
	}

	//TODO: Iterate steps and get the average of steps of each blockchain
	iter := int(max / step)
	for i:=0;i<iter;i++{
		s := ""
		sum := 0
		for j:=0;j<rep;j++{
			sum += SPVSteps(blockChainSlice[j].LightClient, i*step)
		}
		s = strconv.Itoa(sum / rep)
		row := []string{s}
		result = append(result, row)
	}
	return result
}

func exportCSV(title string, data [][]string) {
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

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
// sTest set which as 1 for sublinear complexity verfication
// and 0 for simple payment verfication
// func sTest(which int, step int, max int, avg int) [][]string {

// 	result := [][]string{}
// 	if which == 0 {
// 		result = append(result, []string{"Simple Payment Verification (time)"})

// 	} else if which == 1 {
// 		result = append(result, []string{"Sublinear Complexity Verification (time)"})
// 	} else if which == 2 {
// 		result = append(result, []string{"Simple Payment Verification (steps)"})
// 	} else if which == 3 {
// 		result = append(result, []string{"Sublinear Complexity Verification (steps)"})
// 	}else {
// 		fmt.Println("first argument needs to be 0~3")
// 	}


// 	iter := int(max / step)
// 	testchain := BlockChain{}
// 	testchain.Init()
// 	testchain.GenerateBlocks(max)
// 	//testchain.PrintBlockChain()

// 	for i := 1; i <= iter; i++ {
// 		s := ""
// 		sum := 0
// 		for j := 0; j < avg; j++ {

// 			if which == 0 {
// 				sum += SimplePaymentVerification(testchain.LightClient, i*step)
// 			} else if which == 1 {
// 				sum += SublinearComplexityVerification(testchain.LightClient, i*step)
// 			} else if which == 2 {
// 				sum += SPVSteps(testchain.LightClient, i*step)
// 			} else if which == 3 {
// 				sum += SCVSteps(testchain.LightClient, i*step)
// 			} else{
// 				fmt.Println("which needs to be 0~3")
// 				sum += 0
// 			}
// 		}
// 		s = strconv.Itoa(sum / avg)
// 		row := []string{s}
// 		//scv := strconv.FormatInt(SublinearComplexityVerification(testchain.LightClient, i*step), 10)
// 		result = append(result, row)
// 	}
// 	return result
// }
//TestScvTime will return a 2D slice holding cpu time