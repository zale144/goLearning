package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int				{ return len(a) }
func (a ByAge) Swap(i, j int)			{ a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i, j int) bool		{ return a[i].Age < a[j].Age }

type ByName []Person

func (n ByName) Len() int				{ return len(n) }
func (n ByName) Swap(i, j int) 			{ n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool		{ return n[i].Name < n[j].Name }

type Reversed struct {
	ByName
}

func reverse(i interface{}) sort.Interface {
	t := i.(Reversed)
	return sort.Interface(Reversed(t))
}

func (r Reversed) Less(i, j int) bool	{ return r.ByName[j].Name < r.ByName[i].Name }

func Reverse(b ByName) Reversed {
	return Reversed{b}
}

func main() {

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jennie", 26},
	}
	s := []string{"Zeno", "John", "Al", "Jennie"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	//sort.Sort(ByAge(people))
	fmt.Println("ByAge: ")
	fmt.Println(people)
	sort.Sort(sort.Reverse(ByAge(people)))
	fmt.Println("ByAge reverse: ")
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println("ByName: ")
	fmt.Println(people)
	sort.Sort(reverse(Reverse(people)))
	fmt.Println("ByName reverse: ")
	fmt.Println(people)

	//sort.Strings(s)
	//sort.Sort(sort.StringSlice(s))
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	//sort.Ints(n)
	//sort.Sort(sort.IntSlice(n))
	sort.IntSlice(n).Sort()
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	fmt.Println("b" < "a")

}
