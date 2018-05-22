package main

func fibonacci(n int) (result int) {

	f := []int{0, 1, 1}

	for i := 2; i <= n; i++ {
		f[2] = f[0] + f[1]
		f[0] = f[1]
		f[1] = f[2]
	}

	return f[2]
}
