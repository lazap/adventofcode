package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	plays := strings.Split(string(input), "\n")
	score := 0
	for _, play := range plays {
		score += getRpsResult(play)
	}

	fmt.Println(score)

}

func getRpsResult(play string) int {
	switch play {
	case "A X":
		return 3
	case "A Y":
		return 4
	case "A Z":
		return 8
	case "B X":
		return 1
	case "B Y":
		return 5
	case "B Z":
		return 9
	case "C X":
		return 2
	case "C Y":
		return 6
	case "C Z":
		return 7
	}
	return 0
}
