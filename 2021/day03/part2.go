package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	reports := strings.Split(string(input), "\n")

	numberOfReports := len(reports)
	numberOfBits := len(reports[0])

	oxygenReports := make([]string, numberOfReports)
	for i, report := range reports {
		oxygenReports[i] = report
	}

	co2Reports := make([]string, numberOfReports)
	for i, report := range reports {
		co2Reports[i] = report
	}

	currentBitPosition := 0
	for len(oxygenReports) > 1 && currentBitPosition < numberOfBits {
		numberOfReportsHalved := float64(len(oxygenReports)) / 2
		bitCounts := calculateBitCounts(oxygenReports, numberOfBits)
		fmt.Println(bitCounts, len(oxygenReports), numberOfReportsHalved)

		newOxygenReports := make([]string, 0)
		for _, report := range oxygenReports {
			if float64(bitCounts[currentBitPosition]) >= numberOfReportsHalved {
				if report[currentBitPosition] == '1' {
					newOxygenReports = append(newOxygenReports, report)
				}
			} else {
				if report[currentBitPosition] == '0' {
					newOxygenReports = append(newOxygenReports, report)
				}
			}
		}

		oxygenReports = newOxygenReports
		currentBitPosition++
	}

	currentBitPosition = 0
	for len(co2Reports) > 1 && currentBitPosition < numberOfBits {
		numberOfReportsHalved := float64(len(co2Reports)) / 2
		bitCounts := calculateBitCounts(co2Reports, numberOfBits)

		newCo2Reports := make([]string, 0)
		for _, report := range co2Reports {
			if float64(bitCounts[currentBitPosition]) >= numberOfReportsHalved {
				if report[currentBitPosition] == '0' {
					newCo2Reports = append(newCo2Reports, report)
				}
			} else {
				if report[currentBitPosition] == '1' {
					newCo2Reports = append(newCo2Reports, report)
				}
			}
		}

		co2Reports = newCo2Reports
		currentBitPosition++
	}

	oxygenResult, _ := strconv.ParseInt(oxygenReports[0], 2, 64)
	co2Result, _ := strconv.ParseInt(co2Reports[0], 2, 64)

	fmt.Println(oxygenResult, co2Result)
	fmt.Println(oxygenResult * co2Result)
}

func calculateBitCounts(reports []string, numberOfBits int) []int {
	bitCounts := make([]int, numberOfBits)
	for _, report := range reports {
		for i, bit := range report {
			bitCounts[i] += int(bit - '0')
		}
	}
	return bitCounts
}
