package main

func main()  {
	c := incrementor()
	for n := range c {
		println(n)
		if n >= 5 {
			break
		}
	}

	cSum := puller(c)
	for n := range cSum {
		println("sum =", n)
	}

	for n := range c {
		println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			out <- i
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
			if n >= 8 {
				break
			}
		}
		out <- sum
		close(out)
	}()
	return out
}