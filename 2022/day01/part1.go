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

	calories := strings.Split(string(input), "\n")
	maxCalories := 0
	currentCalories := 0

	for _, calorie := range calories {
		if calorie != "" {
			calorieInt, _ := strconv.Atoi(calorie)
			currentCalories += calorieInt
		} else {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		}
	}

	fmt.Println(maxCalories)

}
