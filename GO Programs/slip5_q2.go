package main

import "fmt"

type Employee struct {
	eno, sal int
	ename    string
}

func main() {
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scanln(&n)
	emp := make([]Employee, n)
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for employee %d:\n", i+1)
		fmt.Print("Employee Number: ")
		fmt.Scanln(&emp[i].eno)
		fmt.Print("Employee Name: ")
		fmt.Scanln(&emp[i].ename)
		fmt.Print("Salary: ")
		fmt.Scanln(&emp[i].sal)
	}

	minSalary := emp[0].sal
	for i := 1; i < n; i++ {
		if emp[i].sal < minSalary {
			minSalary = emp[i].sal
		}
	}

	fmt.Println("\nEmployees with minimum salary")
	for _, emp := range emp {
		if emp.sal == minSalary {
			fmt.Printf("Employee Number: %d\n", emp.eno)
			fmt.Printf("Employee Name: %s\n", emp.ename)
			fmt.Printf("Salary: %d\n\n", emp.sal)
		}
	}
}
