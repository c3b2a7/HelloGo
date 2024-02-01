package concurrent

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Item struct {
	ID            int
	Name          string
	PackingEffort time.Duration
}

func PrepareItem(ctx context.Context) <-chan Item {
	itemCh := make(chan Item)
	items := []Item{
		{0, "Shirt", 0 * time.Second},
		{1, "Legos", 1 * time.Second},
		{2, "TV", 2 * time.Second},
		{3, "Bananas", 3 * time.Second},
		{4, "Hat", 4 * time.Second},
		{5, "Phone", 5 * time.Second},
		{6, "Plates", 6 * time.Second},
		{7, "Computer", 7 * time.Second},
		{8, "Pint Glass", 8 * time.Second},
		{9, "Watch", 9 * time.Second},
	}

	go func() {
		defer close(itemCh)
		for _, item := range items {
			select {
			case itemCh <- item:
			case <-ctx.Done():
				return
			}
		}
	}()

	return itemCh
}

func PackItems(ctx context.Context, itemCh <-chan Item, workerId int) <-chan int {
	packages := make(chan int)
	go func() {
		defer close(packages)
		for item := range itemCh {
			select {
			case <-ctx.Done():
				return
			case packages <- item.ID:
				time.Sleep(item.PackingEffort)
				fmt.Printf("Worker #%d: Shipping package no.%d, took %s to pack %s\n", workerId, item.ID, item.PackingEffort.String(), item.Name)
			}
		}
	}()
	return packages
}

func FanIn[T any](ctx context.Context, chans ...<-chan T) <-chan T {
	out := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch <-chan T) {
			defer wg.Done()
			for t := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- t:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
