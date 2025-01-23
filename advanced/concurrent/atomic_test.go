package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func testAtomic(times uint64, threads uint64) time.Duration {
	start := time.Now()

	var counter uint64
	var wg sync.WaitGroup
	wg.Add(int(threads))

	for range threads {
		go func() {
			defer wg.Done()
			for range times {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()

	if threads*times != counter {
		panic("wrong counter")
	}
	return time.Now().Sub(start)
}

func testMutex(times uint64, threads uint64) time.Duration {
	start := time.Now()

	var counter uint64
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(int(threads))

	for range threads {
		go func() {
			defer wg.Done()
			for range times {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()

	if threads*times != counter {
		panic("wrong counter")
	}
	return time.Now().Sub(start)
}

func TestAtomicAndMutexPerf(t *testing.T) {
	var times uint64 = 10000000
	var threads uint64 = 10

	atomicElapsed := testAtomic(times, threads)
	mutexElapsed := testMutex(times, threads)

	fmt.Printf("AtomicAndMutexPerf: threads:%d,times:%d, atomic_elapsed:%dms,mutex_elapsed:%dms\n",
		threads, times, atomicElapsed.Milliseconds(), mutexElapsed.Milliseconds())
}
