package tasks

import (
	"errors"
	"fmt"
	"sync"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// начало решения

// Queue - FIFO-очередь на n элементов
type Queue struct {
	ch chan int
}

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.

func (q Queue) Get(wait bool) (int, error) {
	if wait {
		return <-q.ch, nil
	}

	select {
	case x := <-q.ch:
		return x, nil
	default:
		return 0, ErrFull
	}
}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, wait bool) error {
	if wait {
		q.ch <- val
		return nil
	}

	select {
	case q.ch <- val:
		return nil
	default:
		return ErrFull
	}
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	return Queue{make(chan int, n)}
}

// конец решения

func queueTask() {
	q := MakeQueue(2)

	err := q.Put(1, false)
	fmt.Println("put 1:", err)

	err = q.Put(2, false)
	fmt.Println("put 2:", err)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		err = q.Put(3, true)
		fmt.Println("put 3:", err)
		wg.Done()
	}()

	res, err := q.Get(false)
	fmt.Println("get:", res, err)

	res, err = q.Get(false)
	fmt.Println("get:", res, err)

	go func() {
		res, err = q.Get(true)
		fmt.Println("get:", res, err)

		wg.Done()
	}()

	wg.Wait()

}
