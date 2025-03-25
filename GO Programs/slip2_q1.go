package main

import "fmt"

func main() {
	var f1, f2, f3, n int
	f1 = 1
	f2 = 2
	f3 = f1 + f2
	fmt.Print("Enter the number of terms to print:")
	fmt.Scanln(&n)
	fmt.Print("Fibonacci Series:")
	fmt.Print(f1, " ", f2)
	for i := 3; i <= n; i++ {
		fmt.Print(" ", f3)
		f1 = f2
		f2 = f3
		f3 = f1 + f2
	}
}
