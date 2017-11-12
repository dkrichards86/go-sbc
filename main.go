package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/dkrichards86/sbc"
	"os"
)

func getValues(bonds []sbc.BondData) {
	total := sbc.SumTotal(bonds)
	interest := sbc.SumInterest(bonds)

	fmt.Println("Values")
	fmt.Println("=====")
	fmt.Printf("Interest Accrued: $%.2f\n", interest)
	fmt.Printf("Total Value: $%.2f\n", total)
	fmt.Println()
}

func flagNotes(bonds []sbc.BondData) {
	fmt.Println("Notes")
	fmt.Println("=====")
	for _, note := range sbc.TakeNote(bonds) {
		fmt.Println(note)
	}
	fmt.Println()
}

func main() {
	filename := flag.String("filename", "bonds.csv", "Name of the CSV file with bond data")
	flag.Parse()

	// Open the CSV
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the CSV
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	// For each line in the CSV, assemble a BondParam and pass it to the scraper
	// The scraper's return value is then appended to a slice for future computation
	bonds := make([]sbc.BondData, 0)
	for _, line := range lines {
		data := sbc.BondParam{
			Series:       line[0],
			SerialNumber: line[1],
			IssueDate:    line[2],
			Denomination: line[3],
		}
		bonds = append(bonds, sbc.Scrape(data))
	}

	getValues(bonds)
	flagNotes(bonds)
}
