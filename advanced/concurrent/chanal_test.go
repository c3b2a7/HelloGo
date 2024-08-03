package concurrent

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

type (
	Job struct {
		ID  int
		Num int
	}
	Result struct {
		Job Job
		Sum int
	}
)

func createWorkerPool(workerNum int, jobCh chan Job, resultCh chan Result) {
	for i := 0; i < workerNum; i++ {
		go func(jobCh chan Job, resultCh chan Result) {
			for job := range jobCh {
				randNum := job.Num
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for randNum != 0 {
					tmp := randNum % 10
					sum += tmp
					randNum /= 10
				}

				resultCh <- Result{job, sum}
			}
		}(jobCh, resultCh)
	}
}

func TestChannel(t *testing.T) {
	jobCh := make(chan Job, 16)
	resultCh := make(chan Result, 16)

	go func(results <-chan Result) {
		for result := range results {
			fmt.Printf("JobID: %d, Num: %d, Sum: %d\n", result.Job.ID, result.Job.Num, result.Sum)
		}
	}(resultCh)

	createWorkerPool(16, jobCh, resultCh)

	var id int
	for id < 100 {
		id++
		randInt := rand.Int()
		jobCh <- Job{id, randInt}
	}
}

func TestConcurrent(t *testing.T) {
	var x int
	var wg sync.WaitGroup

	add := func() {
		for i := 0; i < 10000; i++ {
			x++
		}
		wg.Done()
	}

	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}

func TestConcurrent1(t *testing.T) {
	var x int
	var wg sync.WaitGroup
	var mu sync.Mutex

	add := func() {
		for i := 0; i < 10000; i++ {
			mu.Lock()
			x++
			mu.Unlock()
		}
		wg.Done()
	}

	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	assert.Equal(t, 20000, x)
}

func TestConcurrent2(t *testing.T) {
	var wg sync.WaitGroup
	var x atomic.Int64

	add := func() {
		for i := 0; i < 10000; i++ {
			x.Add(1)
		}
		wg.Done()
	}

	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	assert.EqualValues(t, 20000, x.Load())
}

func TestConcurrent3(t *testing.T) {
	var wg sync.WaitGroup
	var hello = func() {
		defer wg.Done()
		fmt.Println("Hello Goroutine!")
	}

	wg.Add(1)
	go hello()

	fmt.Println("main goroutine done!")
	wg.Wait()
}

func Test(t *testing.T) {
	i := 1
	defer fmt.Println("i = ", func(a int) int { return a }(i)) // expression
	defer func() {
		fmt.Println("i = ", i)
	}() // statement
	i = i * 2
}
