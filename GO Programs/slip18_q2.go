package main

import (
	calculator "GOLang/calc"
	"fmt"
)

func main() {
	var choice int
	var num1, num2, result float64

	fmt.Println("Choose operation:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Scanln(&choice)
	fmt.Println("Enter two numbers:")
	fmt.Scanln(&num1, &num2)

	switch choice {
	case 1:
		result = calculator.Add(num1, num2)
	case 2:
		result = calculator.Subtract(num1, num2)
	case 3:
		result = calculator.Multiply(num1, num2)
	case 4:
		result = calculator.Divide(num1, num2)
	default:
		fmt.Println("Invalid choice")
		return
	}

	fmt.Println("Result:", result)
}
