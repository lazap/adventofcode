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
	depths := make([]int, len(depthStrings))

	for i := 0; i < len(depthStrings); i++ {
		depths[i], _ = strconv.Atoi(depthStrings[i])
	}

	result := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			result++
		}
	}

	fmt.Println(result)
}
