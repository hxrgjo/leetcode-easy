package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for index, value := range nums {
		i, ok := m[value]
		if !ok {
			m[value] = index
			continue
		}
		if index-i <= k {
			return true
		}
		m[value] = index
	}

	return false
}
