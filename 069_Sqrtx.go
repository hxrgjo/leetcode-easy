package main

func mySqrt(x int) int {
	start := 0
	end := x

	if x == 1 {
		return 1
	}

	for start <= end {
		mid := (start + end) / 2

		if mid == 0 {
			return 0
		}

		if mid*mid < x {
			start = mid + 1
		} else if mid*mid > x {
			end = mid - 1
		} else {
			return mid
		}
	}

	return (start + end) / 2
}

func mySqrtLeetCodeSample(y int) int {
	l := 0
	r := y + 1 // [0, y)

	ans := 0
	for l < r {
		mid := l + (r-l)/2

		if guess(mid, y) {
			l = mid + 1
			ans = mid
		} else {
			r = mid
		}
	}

	return ans
}

func guess(x, y int) bool {
	return x*x <= y
}
