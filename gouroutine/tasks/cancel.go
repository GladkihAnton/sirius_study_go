package tasks

import (
	"fmt"
)

// начало решения

// count отправляет в канал числа от start до бесконечности
func count(cancel chan struct{}, start int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			fmt.Println("count finished")
			close(out)
		}()

		for i := start; ; i++ {
			select {
			case out <- i:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// take выбирает первые n чисел из in и отправляет в выходной канал
func take(cancel chan struct{}, in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			fmt.Println("take finished")
		}()

		for i := 0; i < n; i++ {
			select {
			case x := <-in:
				out <- x
			case <-cancel:
				return
			}
		}
		fmt.Println("finisned")
	}()
	return out
}

// конец решения

func mainTakeCount() {
	cancel := make(chan struct{})

	stream := take(cancel, count(cancel, 5), 5)
	first := <-stream
	second := <-stream
	third := <-stream

	close(cancel)

	fmt.Println(first, second, third)
}
