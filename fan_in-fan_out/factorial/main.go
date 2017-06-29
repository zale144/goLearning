package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const ln int = 1000

func main() {
	in := gen()

	f := factorial(in)
	out := merge(f)

	for n:= range out {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i:= 0; i < 1000; i++ {
			out <- rand.Intn(18) + 2
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) [ln]chan int {
	var chans [ln]chan int

	for i := 0; i < ln; i++ {
		chans[i] = make(chan int)
		go func(ind int) {
			chans[ind] <- fact(<-in)
			close(chans[ind])
		}(i)
	}
	return chans
}

func merge(chans [ln]chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(ln)
	for i := 0; i < ln; i++ {
		go func(ind int) {
			out <- <- chans[ind]
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
