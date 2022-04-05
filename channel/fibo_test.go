package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	start := time.Now()

	for i := 1; i < 15; i++ {
		n := fib(float64(i))
		fmt.Printf("Fib(%v): %v\n", i, n)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func TestFibParallelRun(t *testing.T) {
	start := time.Now()

	size := 15
	ch := make(chan string, size)

	for i := 0; i < size; i++ {
		go fibParallelRun(float64(i), ch)
	}

	for i := 0; i < size; i++ {
		fmt.Printf(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func TestFibParallelRun2(t *testing.T) {
	start := time.Now()

	command := ""
	data := make(chan int)
	quit := make(chan bool)

	// FIXME
	go fibParallelRun2(data, quit)

	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
