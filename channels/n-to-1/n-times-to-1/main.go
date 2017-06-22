package main

import (
	"fmt"
)

func main() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i  	// this will only put a value to the channel if there is something to receive that value
			}
			done <- true
		}()
	}

	go func() {     //this has to be in a goroutine, otherwise it will keep waiting for done to be true,
					// which will never happen
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for i:= range c {	   //range clause keeps waiting for the value to arrive
		// if the sender and receiver are not in sync, we have a dead lock
		fmt.Println(i)
	}
}