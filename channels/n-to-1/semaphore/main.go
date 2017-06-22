package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i  	// this will only put a value to the channel if there is something to receive that value
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {     //this has to be in a goroutine, otherwise it will keep waiting for done to be true,
					// which will never happen
		<-done
		<-done
		close(c)
	}()

	for n:= range c {	   //range clause keeps waiting for the value to arrive
							// if the sender and receiver are not in sync, we have a dead lock
		fmt.Println(n)
	}
}