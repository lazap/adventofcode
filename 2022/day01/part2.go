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
	elfCalories := make([]int, 0)

	currentCalories := 0

	for _, calorie := range calories {
		if calorie != "" {
			calorieInt, _ := strconv.Atoi(calorie)
			currentCalories += calorieInt
		} else {
			elfCalories = append(elfCalories, currentCalories)
			currentCalories = 0
		}
	}

	elfCaloriesLength := len(elfCalories)

	// can be done with: sort.Ints(elfCalories)
	// but thats not the point
	for i := 0; i < elfCaloriesLength-1; i++ {
		for j := i + 1; j < elfCaloriesLength; j++ {
			if elfCalories[j] < elfCalories[i] {
				tmp := elfCalories[i]
				elfCalories[i] = elfCalories[j]
				elfCalories[j] = tmp
			}
		}
	}

	fmt.Println(elfCalories[elfCaloriesLength-1] + elfCalories[elfCaloriesLength-2] + elfCalories[elfCaloriesLength-3])
}
