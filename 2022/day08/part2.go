package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// not optimized at all
func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	numOfRows := len(lines)
	numOfColumns := len(lines[0])

	treeMatrix := make([][]int, numOfRows)

	// initialize matrices
	for i := 0; i < numOfRows; i++ {
		treeMatrix[i] = make([]int, numOfColumns)
	}

	// populate treeMatrix from input
	for i, line := range lines {
		for j := range line {
			treeMatrix[i][j] = int(line[j] - '0')
		}
	}

	// find the answer
	maxScenicScore := 0
	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfColumns; j++ {
			currentScenicScore := findViewingDistanceUp(treeMatrix, i, j) *
				findViewingDistanceRight(treeMatrix, i, j) *
				findViewingDistanceDown(treeMatrix, i, j) *
				findViewingDistanceLeft(treeMatrix, i, j)
			if currentScenicScore > maxScenicScore {
				maxScenicScore = currentScenicScore
			}
		}
	}

	printMatrix(treeMatrix, 0, 0)

	fmt.Printf("Max scenic score is: %d", maxScenicScore)
}

func findViewingDistanceUp(matrix [][]int, i int, j int) int {
	distance := 0
	for k := i - 1; k >= 0; k-- {
		if matrix[k][j] < matrix[i][j] {
			distance++
		} else {
			distance++
			break
		}
	}
	return distance
}

func findViewingDistanceRight(matrix [][]int, i int, j int) int {
	distance := 0
	for k := j + 1; k < len(matrix[i]); k++ {
		if matrix[i][k] < matrix[i][j] {
			distance++
		} else {
			distance++
			break
		}
	}
	return distance
}

func findViewingDistanceDown(matrix [][]int, i int, j int) int {
	distance := 0
	for k := i + 1; k < len(matrix); k++ {
		if matrix[k][j] < matrix[i][j] {
			distance++
		} else {
			distance++
			break
		}
	}
	return distance
}

func findViewingDistanceLeft(matrix [][]int, i int, j int) int {
	distance := 0
	for k := j - 1; k >= 0; k-- {
		if matrix[i][k] < matrix[i][j] {
			distance++
		} else {
			distance++
			break
		}
	}
	return distance
}

func printMatrix(matrix [][]int, iHighlight int, jHightlight int) {
	numOfRows := len(matrix)
	numOfColumns := len(matrix[0])

	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfColumns; j++ {
			if i == iHighlight && j == jHightlight {
				fmt.Printf("\033[0;31m")
			}
			fmt.Printf("%d ", matrix[i][j])
			if i == iHighlight && j == jHightlight {
				fmt.Printf("\033[0m")
			}
		}
		fmt.Printf("\n")
	}
}
