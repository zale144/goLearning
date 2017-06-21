package main

import (
	"fmt"
	"bytes"
	"strconv"
)

type Person struct {
	first string
	last string
	age int
}

type DoubleZero struct {
	Person
	Label string
	LicenceToKill bool
}

type foo int

func (f foo) square() foo {
	return f*f
}

func (p Person) stringify() string {
	var str bytes.Buffer
	str.WriteString("Firstname: ")
	str.WriteString(p.first)
	str.WriteString("\nLastname: ")
	str.WriteString(p.last)
	str.WriteString("\nAge: ")
	str.WriteString(strconv.Itoa(p.age))
	return str.String()
}

func main() {
	var e foo = 12

	fmt.Println(e.square())

	p007 := DoubleZero{
		Person: Person {
			first: "James",
			last: "Bond",
			age: 20,
		},
		Label: "Double Zero Seven",
		LicenceToKill: true,
	}
	p007.first = "double o seven"
	fmt.Println(p007)
	p1 := Person{"James", "Bond", 20}
	p2 := Person{"Miss", "Moneypenny", 18}

	people := []Person{p1, p2}

	for _, v := range people {
		fmt.Println(v.stringify(), "\n")
	}
}
