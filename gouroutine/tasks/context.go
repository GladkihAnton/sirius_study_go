package tasks

import (
	"context"
	"fmt"
	"time"
)

// count отправляет в канал числа от start до бесконечности
func countCtx(ctx context.Context, start int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			fmt.Println("count finished")
			close(out)
		}()

		for i := start; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// take выбирает первые n чисел из in и отправляет в выходной канал
func takeCtx(ctx context.Context, in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			fmt.Println("take finished")
		}()

		for i := 0; i < n; i++ {
			select {
			case x := <-in:
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
		fmt.Println("finisned")
	}()
	return out
}

func MainContext() {
	ctx := context.Background()

	ctx1, cancel := context.WithCancel(ctx)

	stream := takeCtx(ctx1, countCtx(ctx1, 5), 5)
	cancel()
	first := <-stream
	second := <-stream
	third := <-stream

	time.Sleep(400 * time.Millisecond)

	fmt.Println(first, second, third)
	time.Sleep(10000 * time.Millisecond)
}
