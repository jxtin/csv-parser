package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	// implemented using encoding/csv

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func myCsvReader(filePath string) [][]string {
	// implemented using strings.Split
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	content := string(file)

	records := [][]string{}

	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	maxColumns := 0

	for _, line := range lines {
		columns := strings.Split(line, ",")
		if len(columns) > maxColumns {
			maxColumns = len(columns)
		}
		records = append(records, columns)
	}

	finalRecords := [][]string{}
	for _, record := range records {
		if len(record) < maxColumns {
			for i := len(record); i < maxColumns; i++ {
				record = append(record, "nan")
			}
		}
		finalRecords = append(finalRecords, record)
	}
	return finalRecords

}

func displayPrettyTable(records [][]string) {
	maxLengths := make([]int, len(records[0]))
	for _, record := range records {
		for i, column := range record {
			if len(column) > maxLengths[i] {
				maxLengths[i] = len(column)
			}
		}
	}

	for _, record := range records {
		for i, column := range record {
			fmt.Print(column)
			for j := 0; j < maxLengths[i]-len(column); j++ {
				fmt.Print(" ")
			}
			fmt.Print(" | ")
		}
		fmt.Println()
	}
}

func main() {
	var filePath string
	fmt.Print("Enter file path: ")
	fmt.Scanln(&filePath)
	displayPrettyTable(myCsvReader(filePath))
	println("---------------------------------")
	displayPrettyTable(readCsvFile(filePath))
}
