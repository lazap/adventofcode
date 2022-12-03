package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	reports := strings.Split(string(input), "\n")

	numberOfReports := len(reports)
	numberOfReportsHalved := numberOfReports / 2
	numberOfBits := len(reports[0])

	bitCounts := make([]int, numberOfBits)
	for _, report := range reports {
		for i, bit := range report {
			bitCounts[i] += int(bit - '0')
		}
	}

	fmt.Println(bitCounts, numberOfReports, numberOfReportsHalved)
	var gammaRate float64
	var epsilonRate float64
	binaryValue := math.Pow(2, float64(numberOfBits-1))
	for _, bitCount := range bitCounts {
		if bitCount > numberOfReportsHalved {
			gammaRate += binaryValue
		} else {
			epsilonRate += binaryValue
		}
		binaryValue /= 2
	}

	fmt.Println(int(gammaRate * epsilonRate))
}
