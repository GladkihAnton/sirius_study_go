package array

func MainMap() {
	// Input: nums = [0,9,11,12], target = 9
	// Output: [1,3]
	// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

	temp := map[string]int{
		"temp": 1,
	}
	temp["asdasd"] = 1231
}

// without len BenchmarkTwoSum/ArraySize-100000-10    	     231	   5257689 ns/op	 5445658 B/op	    2323 allocs/op
// with len BenchmarkTwoSum/ArraySize-100000-10    	     362	   3273566 ns/op	 2512759 B/op	      13 allocs/op
// TODO B/op что это?

func TwoSum(nums []int, target int) []int {
	valueToIndex := make(map[int]int) // TOOD check benchmark
	//valueToIndex := make(map[int]int) // TOOD check benchmark
	for idx, value := range nums {
		findValue := target - value

		index, ok := valueToIndex[findValue]

		if ok {
			return []int{index, idx}
		} else {
			valueToIndex[value] = idx
		}
	}

	return []int{0, 0}
}
