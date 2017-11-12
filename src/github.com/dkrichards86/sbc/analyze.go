package sbc

import "math"

// This round() implementation is taken from a post on StackOverflow
// https://stackoverflow.com/a/29786394/6508952
func round(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	product := num * output
	signed := int(product + math.Copysign(0.5, product))
	
	return float64(signed) / output
}

func SumInterest(bonds []BondData) float64 {
	total := float64(0)

	for _, bond := range bonds {
		total = total + bond.Interest
	}

	return round(total, 2)
}

func SumTotal(bonds []BondData) float64 {
	total := float64(0)

	for _, bond := range bonds {
		total = total + bond.Value
	}

	return round(total, 2)
}

func TakeNote(bonds []BondData) []string {
	notes := make([]string, 0)

	for _, bond := range bonds {
		if bond.Note != "" {
			noteStr := bond.SerialNumber + ": " + bond.Note
			notes = append(notes, noteStr)
		}
	}
	
	return notes
}