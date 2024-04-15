package tasks

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func work() int {
	n := rand.Intn(10)
	if n > 9 {
		time.Sleep(1 * time.Second)
		fmt.Println("long sleep")
	} else {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("short sleep")
	}
	return n
}

func withTimeout(dur time.Duration, fn func() int) (int, error) {
	//after := After(dur)
	after := time.After(dur)

	result := make(chan int, 1)
	//cancel := make(chan struct{})
	go func() {
		defer func() {
			fmt.Println("Closed channel")
			close(result)
		}()
		r := fn()

		//select {
		//case result <- r:
		//	return
		//case <-cancel:
		//	return
		//}
		result <- r
	}()

	select {
	case x := <-after:
		fmt.Println(x)
		//close(cancel)
		return 0, errors.New("too long")
	case v := <-result: // 2
		return v, nil
	}
}

func After(dur time.Duration) <-chan time.Time {
	ch := make(chan time.Time, 1)

	go func() {
		defer func() {
			fmt.Println("Closed after")
		}()

		time.Sleep(dur)
		ch <- time.Now()
	}()

	return ch
}

func bufExample() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 1
	fmt.Println(<-ch)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		ch <- 1 // 1
		wg.Done()
	}()

	go func() {
		fmt.Println(<-ch) // 2
		wg.Done()
	}()

	fmt.Println(<-ch)

	wg.Wait()
}

func nilExample() {
	ch := make(chan int, 2)
	ch = nil
	//close(ch)
	select {
	case <-ch:
		fmt.Println("TEST2")
	case <-time.After(50 * time.Millisecond):
		fmt.Println("TEST")
	}

}

func MainTimeout() {
	nilExample()
	//start := time.Now()
	//value, err := withTimeout(200*time.Millisecond, work)
	//fmt.Println(value, err)
	//value, err = withTimeout(200*time.Millisecond, work)
	//fmt.Println(value, err)
	//value, err = withTimeout(200*time.Millisecond, work)
	//fmt.Println(value, err)
	//
	//fmt.Printf("TIME: %v", time.Since(start))
	//time.Sleep(2 * time.Second)
	//work
}
