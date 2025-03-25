package main

import "fmt"

type student struct {
	roll       int
	name       string
	m1, m2, m3 float64
}

func main() {
	var n int
	var total [10]float64
	var avg [10]float64
	fmt.Print("Enter the number of students: ")
	fmt.Scanln(&n)
	s1 := make([]student, n)
	for i := 0; i < n; i++ {
		fmt.Println("\nEnter the roll number: ")
		fmt.Scanln(&s1[i].roll)
		fmt.Println("Enter the student name: ")
		fmt.Scanln(&s1[i].name)
		fmt.Println("Enter the marks of three subjects: ")
		fmt.Scanln(&s1[i].m1, &s1[i].m2, &s1[i].m3)
		total[i] = s1[i].m1 + s1[i].m2 + s1[i].m3
		avg[i] = total[i] / 3
	}
	for j := 0; j < n; j++ {
		fmt.Println("\nRoll number:", s1[j].roll)
		fmt.Println("Name:", s1[j].name)
		fmt.Println("Marks:", s1[j].m1, "", s1[j].m2, "", s1[j].m3)
		fmt.Println("Total:", total[j])
		fmt.Println("Average:", avg[j])
	}
}
