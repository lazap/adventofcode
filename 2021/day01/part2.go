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

	depthStrings := strings.Split(string(input), "\n")
	depths := make([]int, len(depthStrings)-2)

	for i := 2; i < len(depthStrings); i++ {
		first, _ := strconv.Atoi(depthStrings[i])
		second, _ := strconv.Atoi(depthStrings[i-1])
		third, _ := strconv.Atoi(depthStrings[i-2])
		depths[i-2] = first + second + third
	}

	result := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			result++
		}
	}

	fmt.Println(result)
}
