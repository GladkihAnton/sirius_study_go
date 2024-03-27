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

func twoSum(nums []int, target int) []int {
	valueToIndex := make(map[int]int, len(nums)) // TOOD check benchmark
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
