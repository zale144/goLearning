package main

import (
	"fmt"
	"Tutorial/01_helloWorld/utils"
)


var b string
var c = 144

func main() {
	var c bool
	fmt.Printf("%T\t%v\n", c, c)
	/*for i := 60; i <= 170; i++ {
		fmt.Printf("%b \t %d \t %x \t  %q\n", i, i, i, i)
	}*/

	Print()
}

func printC() int {
	return c
}