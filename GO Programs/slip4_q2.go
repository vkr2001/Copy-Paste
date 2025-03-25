package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int
	fmt.Println("Enter the size of the array:")
	fmt.Scanln(&size)
	arr := make([]int, size)
	fmt.Println("Enter array elements:")
	for i := 0; i < size; i++ {
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scanln(&arr[i])
	}
	sort.Ints(arr)
	fmt.Println("Sorted array in ascending order:", arr)
}
