package interfaces

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// // main находит максимальное число из переданных на вход программы.
func MainTask1() {
	nums := readInput()
	fmt.Println(nums)

	iterator := newIntIterator(nums)
	fmt.Println(iterator)
	//it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(iterator, weight)
	fmt.Println(m)
}

//
//// max возвращает максимальный элемент в последовательности.
//// Для сравнения элементов используется вес, который возвращает
//// функция weight.

func max(it iterator, weight weightFunc) element {
	var maxEl element

	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "x" {
			break
		}

		num, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	return nums
}
