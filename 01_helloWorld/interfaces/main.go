package main

import (
	"math"
	"fmt"
)

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) circumfence() float64 {
	return s.side * 4
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) circumfence() float64 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	area() float64
	circumfence() float64
}

func info(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.area())
	fmt.Println("Circumfence", s.circumfence())
}

func main() {
	var s1 Shape = Square{side: 12}
	var s2 Shape = Circle{radius: 10}
	info(s1)
	info(s2)
}
