package main

import "fmt"

func main() {
	c := make(chan int)

	func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n:= range c {	//range blocks it
		fmt.Println(n)
	}
}