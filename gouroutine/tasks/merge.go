package tasks

import (
	"awesomeProject1/gouroutine"
	"fmt"
)

func merge(chs ...chan int) chan int {
	resultCh := make(chan int)
	done := make(chan struct{})

	for _, ch := range chs {
		go func(ch chan int) {
			for v := range ch {
				resultCh <- v
			}
			done <- struct{}{}
		}(ch)
	}

	go func() {
		for i := 0; i < len(chs); i++ {
			<-done
		}
		close(resultCh)
	}()

	return resultCh
}

func runMerge() {
	cancel := make(chan struct{})
	ch1 := gouroutine.IntGen(cancel, 1, 6)
	ch2 := gouroutine.IntGen(cancel, 4, 7)

	resultCh := merge(ch1, ch2)

	for value := range resultCh {
		fmt.Println(value)
	}
	//var x int
	//var ok bool
	//for {
	//	select {
	//	case x, ok = <-ch1:
	//		if !ok {
	//			ch1 = nil
	//		}
	//	case x, ok = <-ch2:
	//		if !ok {
	//			ch2 = nil
	//		}
	//	default:
	//		if ch1 == nil && ch2 == nil {
	//			return
	//		} else {
	//			continue
	//		}
	//	}
	//	fmt.Println(x)
	//}

}
