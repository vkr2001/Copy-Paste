package main

import "fmt"

type shape interface {
	areac() float32
	peric() float32
	arear() float32
	perir() float32
}

type shape_info struct {
	r, l, b float32
}

func (s shape_info) areac() float32 {
	return ((3.14) * (s.r) * (s.r))
}

func (s shape_info) peric() float32 {
	return (2 * 3.14 * (s.r))
}

func (s shape_info) arear() float32 {
	return (s.l) * (s.b)
}

func (s shape_info) perir() float32 {
	return 2 * ((s.l) + (s.b))
}

func main() {
	sh := shape_info{r: 3.2, l: 4.0, b: 6.0}
	fmt.Println("Area of Circle:", sh.areac())
	fmt.Println("Circumference of Circle:", sh.peric())
	fmt.Println("Area of Rectangle:", sh.arear())
	fmt.Println("Perimeter of rectangle:", sh.perir())
}
