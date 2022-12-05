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
	elfsInGroup := 3

	rucksacks := strings.Split(string(input), "\n")
	rucksackMaps := make([]map[byte]bool, elfsInGroup-1)
	for i, rucksack := range rucksacks {
		fmt.Println(rucksack)
		elfIndex := i % elfsInGroup
		if elfIndex < elfsInGroup-1 {
			rucksackMaps[elfIndex] = make(map[byte]bool)
			for j := 0; j < len(rucksack); j++ {
				rucksackMaps[elfIndex][rucksack[j]] = true
			}
		} else {
			for j := 0; j < len(rucksack); j++ {
				item := rucksack[j]
				found := true
				for _, rucksackMap := range rucksackMaps {
					if !rucksackMap[item] {
						found = false
					}
				}
				if found {
					fmt.Printf("common: %c\n", item)
					if item > 96 && item < 123 {
						sum += int(item - 96)
					} else if item > 64 && item < 91 {
						sum += int(item - 38)
					}
					break
				}
			}
			rucksackMaps = make([]map[byte]bool, elfsInGroup-1)
		}
	}

	fmt.Println(sum)
}
