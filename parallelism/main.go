package main

import (
	"sync"
	"runtime"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func  init()  {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo()  {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar()  {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}