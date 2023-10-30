package concurrent

import (
	"context"
	"fmt"
	"time"
)

func CancelWithChannel() {
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	close(stop)
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(3 * time.Second)
}

func CancelWithCancelContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(3 * time.Second)
}

func CancelMultiGoroutineWithCancelContext() {
	ctx, cancel := context.WithCancel(context.Background())

	name := func(name string) func(ctx context.Context) string {
		return func(ctx context.Context) string {
			return name
		}
	}

	go watch(ctx, name("监控1"))
	go watch(ctx, name("监控2"))
	go watch(ctx, name("监控3"))

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(3 * time.Second)
}

func CancelMultiGoroutineWithNamedCancelContext() {
	ctx, cancel := context.WithCancel(context.Background())

	nameKey := "name"
	getName := func(ctx context.Context) string {
		return ctx.Value(nameKey).(string)
	}
	go watch(context.WithValue(ctx, nameKey, "监控1"), getName)
	go watch(context.WithValue(ctx, nameKey, "监控2"), getName)
	go watch(context.WithValue(ctx, nameKey, "监控3"), getName)

	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(3 * time.Second)
}

func watch(ctx context.Context, getName func(context.Context) string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s监控退出，停止...\n", getName(ctx))
			return
		default:
			fmt.Println("goroutine监控中...")
			time.Sleep(1 * time.Second)
		}
	}
}
