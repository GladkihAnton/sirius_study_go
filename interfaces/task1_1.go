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

	iter.value = iter.nextIter.value
	iter.nextIter = iter.nextIter.nextIter

	return true
}

func (iter *intIterator) val() int {
	return iter.value
}

// методы intIterator, которые реализуют интерфейс iterator

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
	root := &intIterator{0, nil}

	curStep := root
	for _, value := range src {
		temp := &intIterator{value, nil}
		curStep.nextIter = temp
		curStep = temp
	}

	return root
}
