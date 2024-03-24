package interfaces

// element - интерфейс элемента последовательности
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	next() bool
	val() int
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
	value    int
	nextIter *intIterator
}

func (iter *intIterator) next() bool {
	if iter.nextIter == nil {
		return false
	}

	//iter = iter.next
	return true
}

func (iter *intIterator) val() int {
	return iter.value
}

// методы intIterator, которые реализуют интерфейс iterator

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
	root := &intIterator{src[0], nil}

	curStep := root

	for idx, value := range src {
		if idx == 0 {
			continue
		}
		temp := &intIterator{value, nil}
		curStep.nextIter = temp
		curStep = temp
	}

	return root
}
