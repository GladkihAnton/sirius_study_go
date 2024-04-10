package gouroutine

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// say печатает фразу от имени обработчика
func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

// makePool создает пул на n обработчиков
// возвращает функции handle и wait
func makePool(n int, handler func(int, string)) (func(string), func()) {
	pool := make(chan int, n)

	for i := 0; i < n; i++ {
		pool <- i
	}

	handle := func(sentence string) {
		id := <-pool

		go func() {
			handler(id, sentence)
			pool <- id
		}()
	}

	wait := func() {
		for i := 0; i < n; i++ {
			<-pool
		}
	}

	return handle, wait
}

// makePool создает пул на n обработчиков
// возвращает функции handle и wait
func makeBusyPool(n int, handler func(int, string)) (func(string) error, func()) {
	pool := make(chan int, n)

	for i := 0; i < n; i++ {
		pool <- i
	}

	handle := func(sentence string) error {
		select {
		case id := <-pool:
			go func() {
				handler(id, sentence)
				pool <- id
			}()
		default:
			return errors.New("pool busy")
		}
		return nil
	}

	wait := func() {
		for i := 0; i < n; i++ {
			<-pool
		}
	}

	return handle, wait
}

// конец решения

func runPool() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}

	handle, wait := makeBusyPool(2, say)
	for _, phrase := range phrases {
		err := handle(phrase)
		fmt.Println(err)
	}
	wait()
}
