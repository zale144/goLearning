package main

import "fmt"

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

func main() {
	p := Person{"Alex"}
	o := Osoba{"Aleks"}

	p.printP()
	o.printP()
}
