package tasks

import (
	"fmt"
	"sync/atomic"
	"time"

	//"sync/atomic"
	"sync"
)

func exampleInt() {
	wg := sync.WaitGroup{}

	now := time.Now()

	//mx := sync.Mutex{}
	//j := 0

	j := atomic.Int32{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			j.Add(1)
			// mx.Lock()
			//j += 1
			//	mx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Since: %v, j: %v", time.Since(now), j.Load())
}

func exampleSlice() {
	n := 100
	ints := make([]int, n)
	wg := sync.WaitGroup{}

	now := time.Now()

	wg.Add(n)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			ints[i%n] += 1
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Since: %v, even: %v, odd: %v", time.Since(now), ints[0], ints[1])
}
func exampleRlock() {
	n := 5
	map_ := make(map[int]int, 10)
	//j := 0
	//map_ := make(map[int]int, n)

	wg := sync.WaitGroup{}
	mx := sync.RWMutex{}

	now := time.Now()

	wg.Add(5 * n)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			mx.Lock()
			map_[i%n] += 1
			mx.Unlock()
			wg.Done()
		}()
	}
	for i := 0; i < 4*n; i++ {
		go func() {
			mx.RLock()
			time.Sleep(50 * time.Millisecond)
			fmt.Println(map_)
			mx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Since: %v", time.Since(now))
}

func MainRace() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	//pool := make(chan struct{})
	//pool <- struct{}{}

	go func() {
		fmt.Println("first")
		wg.Done()
	}()
	go func() {
		fmt.Println("second")
		wg.Done()
	}()
	go func() {
		fmt.Println("third")
		wg.Done()
	}()
	wg.Wait()
}
