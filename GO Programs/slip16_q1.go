package main

import (
	rectangle "GOLang/rect"
	"fmt"
)

func main() {
	var l, b int
	fmt.Print("Enter the length and width of rectangle: ")
	fmt.Scanln(&l, &b)
	area := rectangle.Area(l, b)
	fmt.Printf("Area of rectangle with length %d and width %d is %.d\n", l, b, area)
}
