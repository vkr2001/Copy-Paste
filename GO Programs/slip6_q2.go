package main

import "fmt"

func CopyArray(org []int, dst []int) {
	for i := 0; i < len(org); i++ {
		dst[i] = org[i]
	}
}

func main() {
	org := []int{1, 2, 3, 4, 5, 7, 6, 9}
	dst := make([]int, len(org))

	CopyArray(org, dst)

	fmt.Println("Original Array:", org)
	fmt.Println("Copied Array:", dst)
}
