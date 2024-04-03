package gouroutine

import (
	"fmt"
	"sync"
	"time"
)

func temp(wg *sync.WaitGroup, ch chan int, i int) {
	defer func() {
		wg.Done()
	}()
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

func MainGoroutine() {
	//start := time.Now()
	//wg := sync.WaitGroup{}
	//ch := make(chan int)
	//
	//wg.Add(1)
	//go temp(&wg, ch, 1)
	//
	////wg.Add(1)
	////go temp(&wg, ch, 2)
	//
	//wg.Wait()
	//fmt.Println(time.Now().Sub(start))
	mainTask2()
}
