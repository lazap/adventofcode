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

	sum := 0

	rucksacks := strings.Split(string(input), "\n")
	for _, rucksack := range rucksacks {
		itemNumber := len(rucksack)
		seenItems := make(map[byte]bool)
		for i := 0; i < itemNumber/2; i++ {
			seenItems[rucksack[i]] = true
		}
		for i := itemNumber / 2; i < itemNumber; i++ {
			if seenItems[rucksack[i]] {
				if rucksack[i] > 96 && rucksack[i] < 123 {
					fmt.Printf("%c - %d\n", rucksack[i], rucksack[i]-96)
					sum += int(rucksack[i] - 96)
				} else if rucksack[i] > 64 && rucksack[i] < 91 {
					fmt.Printf("%c - %d\n", rucksack[i], rucksack[i]-38)
					sum += int(rucksack[i] - 38)
				}
				break
			}
		}
	}

	fmt.Println(sum)
}
