package tasks

import (
	"fmt"
	"time"
)

func MainTicker() {
	ticker := time.NewTicker(150 * time.Millisecond)

	now := time.Now()

	go func() {
		for {
			<-ticker.C

			fmt.Println(fmt.Printf("Gor 1 %v", time.Since(now)))
		}
	}()
	go func() {
		for {
			<-ticker.C

			fmt.Println(fmt.Printf("Gor 2 %v", time.Since(now)))
		}
	}()
	time.Sleep(10 * time.Second)
}
