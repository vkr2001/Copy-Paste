package main

import "fmt"

func main() {
	var matrix1 [3][3]int
	var matrix2 [3][3]int
	var result [3][3]int

	fmt.Println("Enter elements of the first matrix:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Enter element [%d][%d]: ", i, j)
			fmt.Scan(&matrix1[i][j])
		}
	}

	fmt.Println("Enter elements of the second matrix:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Enter element [%d][%d]: ", i, j)
			fmt.Scan(&matrix2[i][j])
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	fmt.Println("Result of matrix multiplication:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
