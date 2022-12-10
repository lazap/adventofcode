package main

import (
	"fmt"
	"io/ioutil"
)

// not optimal at all!
func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	wantedNumberOfChars := 14

	numberOfChars := 0
	seenChars := make([]byte, wantedNumberOfChars)
	for i, char := range []byte(input) {
		if !contains(char, seenChars) {
			seenChars[numberOfChars] = char
			numberOfChars++
		} else {
			seenCharIndex := find(char, seenChars)
			toCopy := seenChars[seenCharIndex+1:]
			seenChars = make([]byte, wantedNumberOfChars)
			copy(toCopy, seenChars)
			numberOfChars -= (seenCharIndex + 1)

			seenChars[numberOfChars] = char
			numberOfChars++
		}
		if numberOfChars == wantedNumberOfChars {
			fmt.Println(i + 1)
			break
		}
	}
}

func contains(wantedItem byte, items []byte) bool {
	for _, item := range items {
		if item == wantedItem {
			return true
		}
	}
	return false
}

func find(wantedItem byte, items []byte) int {
	for i, item := range items {
		if item == wantedItem {
			return i
		}
	}
	return -1
}

func copy(from []byte, to []byte) {
	for i, elem := range from {
		to[i] = elem
	}
}
