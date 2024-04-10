package gouroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func temp(wg *sync.WaitGroup, ch chan int, i int) {
	var fromChannel int
	if i == 1 {
		ch <- i
	} else {
		fromChannel = <-ch
	}
	fmt.Println(fromChannel)

	time.Sleep(1 * time.Second)
	fmt.Println("test")
}

func exampleWithDoneChannel() {
	words := []string{
		"a",
		"b",
		"c",
	}
	ch := make(chan string)
	done := make(chan struct{})

	go func() {
		for _, word := range words {
			ch <- word
		}
		close(ch)
	}()

	go func() {
		for word := range ch {
			if word != "" {
				fmt.Println(word)
			}
		}
		done <- struct{}{}
	}()

	for i := 0; i < 2; i++ {
		<-done
	}
}

func exampleWithWaitGroup() {
	words := []string{
		"a",
		"b",
		"c",
	}
	ch := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for _, word := range words {
			ch <- word
		}
		wg.Done()
		close(ch)
	}()

	wg.Add(1)
	go func() {
		for word := range ch {
			if word != "" {
				fmt.Println(word)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
func exampleSelect() {
	cancel := make(chan int)
	second := make(chan int)
	go func() {
		x := rand.Intn(2)
		time.Sleep(time.Duration(x) * time.Second)
		cancel <- 1
	}()
	go func() {
		x := rand.Intn(2)
		time.Sleep(time.Duration(x) * time.Second)
		second <- 2
	}()

	select {
	case <-cancel:
		fmt.Println("cancel")
		return
	case y := <-second:
		fmt.Println(y)
	}
}

func channelBuff() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	x, ok := <-ch
	fmt.Println(x, ok)
	for i := range ch {
		fmt.Println(i)
	}

	x1, ok1 := <-ch
	fmt.Println(x1, ok1)
}

func intGen(cancel chan struct{}, start, stop int) chan int {
	result := make(chan int)

	go func() {
		defer func() {
			close(result)
		}()

		for i := start; i < stop; i++ {
			select {
			case result <- i:
				time.Sleep(300 * time.Millisecond)
			case <-cancel:
				return
			}
		}
	}()

	return result
}

func intGenLeak() {
	cancel := make(chan struct{})
	in := intGen(cancel, 10, 20)
	for i := range in {
		if i == 13 {
			cancel <- struct{}{}
			break
		}
		fmt.Println(i)
	}
	fmt.Println("finished")
}

func MainGoroutine() {
	queueTask()
}
