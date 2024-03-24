package array

import "fmt"

func Main() {
	array := [4]int{1, 2, 3, 4}
	slice1 := array[:]
	name := slice1[1:]

	fmt.Println(name, array)
	name = appendInt(name, 10)
	fmt.Println(name, name[:cap(name)], array)
	name = appendInt(name, 11)
	// name = append(name, 11)

	fmt.Println(name, name[:cap(name)], array)
}

func appendInt(array []int, value int) []int {
	if len(array) < cap(array) {
		array = array[:len(array)+1]
		array[len(array)-1] = value
	} else {
		newArray := make([]int, len(array)+1, len(array)*2)
		copy(newArray, array)
		newArray[len(newArray)-1] = value
		return newArray
	}
	return array
}
