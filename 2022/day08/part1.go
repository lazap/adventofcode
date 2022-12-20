package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// optimized using memoization, not performance tested, but probably better
// then iterating whole rows and columns for every element of a input matrix
func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	numOfRows := len(lines)
	numOfColumns := len(lines[0])

	treeMatrix := make([][]int, numOfRows)
	downHeightMatrix := make([][]int, numOfRows)
	upHeightMatrix := make([][]int, numOfRows)
	rightHeightMatrix := make([][]int, numOfRows)
	leftHeightMatrix := make([][]int, numOfRows)

	// initialize matrices
	for i := 0; i < numOfRows; i++ {
		treeMatrix[i] = make([]int, numOfColumns)
		downHeightMatrix[i] = make([]int, numOfColumns)
		upHeightMatrix[i] = make([]int, numOfColumns)
		rightHeightMatrix[i] = make([]int, numOfColumns)
		leftHeightMatrix[i] = make([]int, numOfColumns)
	}

	// populate treeMatrix from input
	for i, line := range lines {
		for j := range line {
			treeMatrix[i][j] = int(line[j] - '0')
		}
	}

	// copy edge values from to heightMatrices from treeMatrix
	for i := 0; i < numOfColumns; i++ {
		downHeightMatrix[0][i] = treeMatrix[0][i]
		upHeightMatrix[numOfRows-1][i] = treeMatrix[numOfRows-1][i]
	}

	for i := 0; i < numOfRows; i++ {
		rightHeightMatrix[i][0] = treeMatrix[i][0]
		leftHeightMatrix[i][numOfColumns-1] = treeMatrix[i][numOfColumns-1]
	}

	// calculate remaining elements in height matrices
	for i := 1; i < numOfRows; i++ {
		for j := 1; j < numOfColumns; j++ {
			downHeightMatrix[i][j] = max(downHeightMatrix[i-1][j], treeMatrix[i][j])
			upHeightMatrix[numOfRows-i-1][j] = max(upHeightMatrix[numOfRows-i][j], treeMatrix[numOfRows-i-1][j])
			rightHeightMatrix[i][j] = max(rightHeightMatrix[i][j-1], treeMatrix[i][j])
			leftHeightMatrix[i][numOfColumns-j-1] = max(leftHeightMatrix[i][numOfColumns-j], treeMatrix[i][numOfColumns-j-1])
		}
	}

	// printMatrix(downHeightMatrix)

	// find the answer
	visibleTreeCount := 0
	for i := 1; i < numOfRows-1; i++ {
		for j := 1; j < numOfColumns-1; j++ {
			currentTreeHeight := treeMatrix[i][j]
			if currentTreeHeight > downHeightMatrix[i-1][j] ||
				currentTreeHeight > upHeightMatrix[i+1][j] ||
				currentTreeHeight > rightHeightMatrix[i][j-1] ||
				currentTreeHeight > leftHeightMatrix[i][j+1] {
				visibleTreeCount++
			}
		}
	}

	fmt.Printf("Result is: %d\n", visibleTreeCount+numOfColumns*2+numOfRows*2-4)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func printMatrix(matrix [][]int) {
	numOfRows := len(matrix)
	numOfColumns := len(matrix[0])

	for i := 0; i < numOfRows; i++ {
		for j := 0; j < numOfColumns; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}
