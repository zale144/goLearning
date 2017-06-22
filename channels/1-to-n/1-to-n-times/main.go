package main

import (
	//"fmt"
)

func main() {

	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		println("sending...")
		for i := 0; i < 10000; i++ {
			println("sending", i)
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		println("goroutine ", i, " receiving...")
		go func() {
			for a:= range c {
				println("received: ", a, "from goroutine ", i)
			}
			println("done")
			done <- true
		}()
	}

	//for i := 0; i < n; i++ {
		<-done
	//}

}