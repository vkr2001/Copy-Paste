package main

import (
	"fmt"
)

type shape interface {
	area() float32
	vol() float32
}

type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side * s.side
}

func (s square) vol() float32 {
	return s.side * s.side * s.side
}

type rectangle struct {
	l, w, h float32
}

func (r rectangle) area() float32 {
	return r.l * r.w
}

func (r rectangle) vol() float32 {
	return r.l * r.w * r.h
}

func main() {
	sqr := square{side: 5}
	rect := rectangle{l: 4, w: 6, h: 3}

	fmt.Println("Area of Square:", sqr.area())
	fmt.Println("Volume of Cube:", sqr.vol())
	fmt.Println("Area of Rectangle:", rect.area())
	fmt.Println("Volume of Cuboid:", rect.vol())
}