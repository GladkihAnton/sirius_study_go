package array

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkTwoSum(b *testing.B) {
	rand.Seed(0)
	target := 9
	for _, size := range []int{1, 10, 100, 1000, 10000, 100000} {
		array := randomArray(size)
		name := fmt.Sprintf("ArraySize-%d", size)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				TwoSum(array, target)
			}
		})
	}
}

func randomArray(size int) []int {
	array := make([]int, size+2)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(100000) + 10
	}
	array[size] = 0
	array[size+1] = 9
	return array
}
