package main

import (
	"errors"
	"math"
	"log"
	"fmt"
)

func main() {
	n, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return math.Sqrt(f), nil
}