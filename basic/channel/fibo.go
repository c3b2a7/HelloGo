package channel

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64) float64 {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	return x
}

func fibParallelRun(number float64, ch chan<- string) (value float64) {
	value = fib(number)
	ch <- fmt.Sprintf("Fib(%v)=%v\n", number, value)
	return
}

func fibParallelRun2(fib chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case fib <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}
