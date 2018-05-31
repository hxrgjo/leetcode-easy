package main

import (
	"math"
)

func maxSubArray(nums []int) int {
	currentSum := 0
	maxSum := math.MinInt32

	for index := range nums {

		if currentSum < 0 {
			currentSum = nums[index]
		} else {
			currentSum += nums[index]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}

	}

	return maxSum
}
