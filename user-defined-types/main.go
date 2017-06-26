package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string `json:"name"`
}

func (p Person) printP() {
	fmt.Println("Name:", p.name)
}

func (o Osoba) printP() {
	fmt.Println("Ime:", o.name)
}

type Osoba Person

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

type Shape interface{}

func area(s Shape) float64 {
	switch t := s.(type) {
	case Square:
		return math.Pow(Square(t).side, 2)
	case Circle:
		return math.Pow(Circle(t).radius, 2) * math.Pi
	default:
		return -1
	}
}

func main() {

	p := Person{"Alex"}
	o := Osoba{"Aleks"}

	p.printP()
	o.printP()

	var s Shape = Square{side: 12}
	t := &s
	b := *t
	fmt.Println(b)

	fmt.Printf("%T\t%d\n", s, area(s))

	var c Shape = Circle{radius: 13}
	fmt.Printf("%T\t%d\n", c, area(c))
}
