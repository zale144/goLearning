package main

import (
	"fmt"
	"math/big"
)

func main()  {
	var n int64 = 7
	f := factorialFast(n)
	fmt.Println("factorial of", n, "is", f)
}

func factorial(n int64) *big.Int {
	var fact big.Int
	c := make(chan int64)
	done := make(chan bool)

	go func() {
		for i := n; i > 0; i-- {
			c <- i
		}
		close(c)
	}()

	go func() {
		fact.MulRange(1, n)
		done <- true
		close(done)
	}()

	<-done
	return &fact
}

func factorialSlow(n int64) *big.Int {
	var f big.Int
	f.MulRange(1, n)
	return &f
}

func factorialFast(n int64) *big.Int {
	var fact big.Int = *big.NewInt(1)
	c := make(chan *big.Int)
	done := make(chan bool)

	go func() {
		c <- big.NewInt(n)
		done <- true
	}()

	for i := int64(1); i < n; i++ {
		go func(ind int64) {
			c <- fact.Mul(big.NewInt(ind), <-c)
			done <- true
		}(i)
	}

	<-done
	<-done

	return <-c
}
