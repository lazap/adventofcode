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

	count := 0

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		assigments := strings.Split(line, ",")
		firstAssigment := strings.Split(assigments[0], "-")
		secondAssigment := strings.Split(assigments[1], "-")
		firstAssigmentX, _ := strconv.Atoi(firstAssigment[0])
		firstAssigmentY, _ := strconv.Atoi(firstAssigment[1])
		secondAssigmentX, _ := strconv.Atoi(secondAssigment[0])
		secondAssigmentY, _ := strconv.Atoi(secondAssigment[1])

		if (firstAssigmentX >= secondAssigmentX && firstAssigmentX <= secondAssigmentY) ||
			(firstAssigmentY >= secondAssigmentX && firstAssigmentY <= secondAssigmentY) ||
			(secondAssigmentX >= firstAssigmentX && secondAssigmentX <= firstAssigmentY) ||
			(secondAssigmentY >= firstAssigmentX && secondAssigmentY <= firstAssigmentY) {
			count++
		}

	}

	fmt.Println(count)

}
