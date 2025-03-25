package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func main() {
	var rows, cols int
	fmt.Println("Enter number of rows and columns of the matrix:")
	fmt.Scanln(&rows, &cols)

	fmt.Println("Enter the elements of the matrix:")
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	fmt.Println("Original Matrix:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	transposed := transpose(matrix)

	fmt.Println("Transpose Matrix:")
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			fmt.Printf("%d ", transposed[i][j])
		}
		fmt.Println()
	}
}
