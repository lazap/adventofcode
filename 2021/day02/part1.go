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

	commands := strings.Split(string(input), "\n")
	horizontalPosition := 0
	depth := 0
	for _, command := range commands {
		commandParts := strings.Split(command, " ")
		if len(commandParts) != 2 {
			continue
		}
		commandValue, _ := strconv.Atoi(commandParts[1])
		switch commandParts[0] {
		case "forward":
			horizontalPosition += commandValue
		case "down":
			depth += commandValue
		case "up":
			depth -= commandValue
		}
	}

	fmt.Println(horizontalPosition * depth)
}
