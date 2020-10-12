package iotsim

import (
	//"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

type StorageRow struct {
	in        time.Time
	out       time.Time
	productId string
}

type IotRow struct {
	periodStart time.Time
	temp        string
	humidity    string
}

func ReadStorage() {
	// Open the file
	csvfile, err := os.Open("./iotsim/data-storage.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	var result []StorageRow
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		layout := "2/1/06 15:04"
		intime, _ := time.Parse(layout, record[0])
		outtime, _ := time.Parse(layout, record[1])
		newrow := &StorageRow{in: intime, out: outtime, productId: record[2]}
		//fmt.Println(newrow.in, newrow.out, newrow.productId)
		result = append(result, *newrow)

	}
	RECORD = result
}

func ReadIot() {
	// Open the file
	csvfile, err := os.Open("./iotsim/data-iot.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	var result []IotRow
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		layout := "2/1/06 15:04"
		time, _ := time.Parse(layout, record[0])
		newrow := &IotRow{
			periodStart: time,
			temp:        record[1],
			humidity:    record[2]}
		//fmt.Println(newrow.periodStart, newrow.temp, newrow.humidity)
		result = append(result, *newrow)

	}
	DATA = result
	//fmt.Println("length in data read", len(DATA))
}
