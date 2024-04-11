package structs

import "fmt"

type MyCircularDeque struct {
	Buf []int

	head int
	tail int

	size int
	cap  int
}

func Constructor(k int) MyCircularDeque {
	Buf := make([]int, k)
	return MyCircularDeque{Buf, 0, 0, 0, k}
}

func (deque *MyCircularDeque) InsertFront(value int) bool {
	if deque.IsFull() {
		return false
	}

	if deque.IsEmpty() {
		deque.Buf[deque.head] = value
		deque.size += 1
		return true
	}

	deque.size += 1
	deque.head = deque.next(deque.head)
	deque.Buf[deque.head] = value
	return true
}

func (deque *MyCircularDeque) InsertLast(value int) bool {
	if deque.IsFull() {
		return false
	}

	if deque.IsEmpty() {
		deque.Buf[deque.tail] = value
		deque.size += 1
		return true
	}

	deque.size += 1
	deque.tail = deque.prev(deque.tail)
	deque.Buf[deque.tail] = value
	return true
}

//
//func (deque *MyCircularDeque) DeleteFront() bool {
//
//}
//
//
//func (deque *MyCircularDeque) DeleteLast() bool {
//
//}

func (deque *MyCircularDeque) GetFront() int {
	return deque.Buf[deque.head]
}

func (deque *MyCircularDeque) GetTail() int {
	return deque.Buf[deque.tail]
}

func (deque *MyCircularDeque) IsEmpty() bool {
	return deque.size == 0
}

func (deque *MyCircularDeque) IsFull() bool {
	return deque.cap == deque.size
}

func (deque *MyCircularDeque) next(step int) int {
	if step == deque.size-1 {
		return 0
	} else {
		return step + 1
	}
}

// 5
// 0 5 4 3 2 1 0 5 4 3 2 1 0
func (deque *MyCircularDeque) prev(step int) int {
	if step == 0 {
		return deque.cap - 1
	} else {
		return step - 1
	}
}

//1
//-1
//2
//-2
//3
//-3

// 1 3 2

// <-
// ->
// 1 0 2
// getFirst() -> isEmpty
// getFirst() -> buf[tail]
// getLast() -> buf[head]
// getLast() -> buf[head]
// getLast() -> buf[head]
// deleteLast() -> delete(buf[head])

func MainDeque() {
	deque := Constructor(3)
	deque.InsertFront(1)
	fmt.Println(deque.Buf)
	deque.InsertLast(2)
	fmt.Println(deque.Buf)
	deque.InsertLast(3)
	fmt.Println(deque.Buf)
	fmt.Println(deque.GetTail())
	fmt.Println(deque.GetFront())
}
