package tasks

import (
	"fmt"
	"time"
)

func work1() {
	fmt.Println("BEFORE WORK")
	time.Sleep(150 * time.Millisecond)
	fmt.Println("AFTER WORK")
}

func afterFunc(dur time.Duration, fn func()) (stop func() bool) {
	cancel := make(chan struct{})
	done := make(chan struct{})

	go func() {
		select {
		case <-time.After(dur):
			close(done)
			fn()
		case <-cancel:
			cancel = nil
			return
		}
	}()

	stop = func() bool {
		select {
		case <-done:
			return false
		default:
			if cancel != nil {
				cancel <- struct{}{}
				close(done)
			}

			return true
		}
	}

	return
}

func withDelay(dur time.Duration, fn func() int) *time.Timer {
	timer := time.AfterFunc(dur, work1)

	return timer
}

func MainTimer() {
	now := time.Now()
	fmt.Println(now)
	dur := 300 * time.Millisecond
	stop := afterFunc(dur, work1)

	time.Sleep(400 * time.Millisecond)
	fmt.Println(stop())
	fmt.Println(stop())
	fmt.Println(stop())
	fmt.Println(stop())
	fmt.Println(stop())
	fmt.Println(stop())
	fmt.Println(stop())
	//
	//timer.Reset(dur)
	//
	//value := <-timer.C
	//
	//fmt.Println(value)

	//timer := withDelay(dur, work)
	//time.Sleep(550 * time.Millisecond)
	//fmt.Println(timer.Stop())
	////fmt.Println(v, err)
	//fmt.Println(time.Since(now))

}
