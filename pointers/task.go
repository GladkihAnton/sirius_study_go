package pointers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	result := make([]int, len(iterable))
	cnt := 0
	for _, value := range iterable {
		if predicate(value) {
			result[cnt] = value
			cnt += 1
		}
	}
	return result
}

func mainTask1() {
	src := readInput()
	res := filter(func(value int) bool {
		return value%2 == 0
	}, src)

	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	// res := filter(...)

	fmt.Println(res)
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "x" {
			break
		}
		num, _ := strconv.Atoi(text)
		nums = append(nums, num)

		//fmt.Println(nums)
		//break
	}
	fmt.Println(nums)
	return nums
}
