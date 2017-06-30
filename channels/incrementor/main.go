package main

import "fmt"

func main()  {
	c1 := incrementor("Foo")
	c2 := incrementor("Bar")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter: ", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/*
	The optional <- operator specifies the channel direction, send or receive.
	If no direction is given, the channel is bidirectional.

	chan T - can be used to send and receive values of type T
	chan<- float64 - can only be used to send float64s
	<-chane int - can only be used to receive ints
*/