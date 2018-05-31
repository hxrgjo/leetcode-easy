package main

func removeElement(nums []int, val int) int {
	var result = 0
	var counter = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[counter] = nums[i]
			counter++
			result++
		}
	}

	return result
}
