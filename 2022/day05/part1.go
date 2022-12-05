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

	lines := strings.Split(string(input), "\n")
	lines = lines[10:]

	stacks := make([][]byte, 9)
	stacks[0] = []byte{'P', 'F', 'M', 'Q', 'W', 'G', 'R', 'T'}
	stacks[1] = []byte{'H', 'F', 'R'}
	stacks[2] = []byte{'P', 'Z', 'R', 'V', 'G', 'H', 'S', 'D'}
	stacks[3] = []byte{'Q', 'H', 'P', 'B', 'F', 'W', 'G'}
	stacks[4] = []byte{'P', 'S', 'M', 'J', 'H'}
	stacks[5] = []byte{'M', 'Z', 'T', 'H', 'S', 'R', 'P', 'L'}
	stacks[6] = []byte{'P', 'T', 'H', 'N', 'M', 'L'}
	stacks[7] = []byte{'F', 'D', 'Q', 'R'}
	stacks[8] = []byte{'D', 'S', 'C', 'N', 'L', 'P', 'H'}

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		n, _ := strconv.Atoi(instruction[1])
		from, _ := strconv.Atoi(instruction[3])
		to, _ := strconv.Atoi(instruction[5])

		fromLength := len(stacks[from-1])
		tmp := stacks[from-1][fromLength-n:]
		reverseWithEffect(tmp)
		stacks[from-1] = stacks[from-1][:fromLength-n]
		stacks[to-1] = append(stacks[to-1], tmp...)
	}

	for _, stack := range stacks {
		stackSize := len(stack)
		fmt.Printf("%s", stack[stackSize-1:])
	}

}

func reverseWithEffect(arr []byte) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		tmp := arr[i]
		arr[i] = arr[length-i-1]
		arr[length-i-1] = tmp
	}
}
