package main

import "math"

func constructRectangle(area int) []int {
	r := 0
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			r = i
		}
	}
	return []int{area / r, r}
}

func constructRectangleLeetCode(area int) []int {
	if area == 1 {
		return []int{1, 1}
	}
	var i int
	for i = int(math.Sqrt(float64(area))); i >= 1; i-- {
		if area%i == 0 {
			break
		}
	}
	c := area / i
	if c > i {
		return []int{c, i}
	}
	return []int{i, c}
}
