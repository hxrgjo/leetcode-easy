package main

func removeDuplicates(nums []int) int {
	var result int

	m := make(map[int]interface{})

	for i := range nums {
		if _, ok := m[nums[i]]; !ok {
			nums[result] = nums[i]
			result++
			m[nums[i]] = 1
		}
	}

	return result
}

func removeDuplicatesFaster(nums []int) int {
	var result int = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if i != result {
			nums[result] = nums[i]
		}
		result++
	}

	return result
}
