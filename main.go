package main

import "fmt"

func main(){
var r = rand.New(rand.NewSource(99))
func main() {

	test(100,10000,100)


func test(step int, max int, rep int){
	exportCSV("scv-time.csv", TestScvTime(step, max, rep))
	exportCSV("scv-step.csv", TestScvStep(step, max, rep))
	exportCSV("spv-time.csv", TestSpvTime(step, max, rep))
	exportCSV("spv-step.csv", TestSpvStep(step, max, rep))	

}