package concurrent

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestNormal(t *testing.T) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	itemCh := PrepareItem(ctx)
	packages := PackItems(ctx, itemCh, 0)

	numPackages := 0
	for range packages {
		numPackages++
	}
	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}

func TestFanIn(t *testing.T) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	itemCh := PrepareItem(ctx)
	workerNum := runtime.NumCPU()
	workers := make([]<-chan int, workerNum)
	for idx := range workers {
		workers[idx] = PackItems(ctx, itemCh, idx)
	}

	numPackages := 0
	for range FanIn[int](ctx, workers...) {
		numPackages++
	}
	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
