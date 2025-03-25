package main

import "fmt"

func main() {
	var choice, num1, num2 int
	fmt.Println("Choose an arithmetic operation:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Print("Enter your choice:")
	fmt.Scanln(&choice)

	fmt.Print("Enter the two numbers: ")
	fmt.Scanln(&num1, &num2)

	switch choice {
	case 1:
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case 2:
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case 3:
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Cannot divide by zero!")
		}
	default:
		fmt.Println("Invalid choice")
	}
}
