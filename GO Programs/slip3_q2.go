package main

import "fmt"

type Employee struct {
	Eno    int
	Ename  string
	Salary int
}

func main() {
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scanln(&n)
	emp := make([]Employee, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for Employee %d:\n", i+1)
		fmt.Print("Employee Number: ")
		fmt.Scanln(&emp[i].Eno)
		fmt.Print("Employee Name: ")
		fmt.Scanln(&emp[i].Ename)
		fmt.Print("Salary: ")
		fmt.Scanln(&emp[i].Salary)
	}

	maxsal := emp[0].Salary
	for i := 1; i < n; i++ {
		if emp[i].Salary > maxsal {
			maxsal = emp[i].Salary
		}
	}

	fmt.Println("\nEmployee with maximum salary:")
	for _, emp := range emp {
		if emp.Salary == maxsal {
			fmt.Printf("Employee Number: %d\n", emp.Eno)
			fmt.Printf("Employee Name: %s\n", emp.Ename)
			fmt.Printf("Salary: %d\n", emp.Salary)
			fmt.Println("-------------------------")
		}
	}
}
