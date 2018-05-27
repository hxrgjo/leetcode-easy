package main

import "strconv"

func countAndSay(n int) string {
	var result = "1"

	if n == 1 {
		return "1"
	}

	for i := 2; i <= n; i++ {
		var temp = string(result[0])
		var tempCounter = 1

		tempResult := result

		result = ""

		for k := 1; k < len(tempResult); k++ {

			if string(tempResult[k]) != temp {
				result += strconv.Itoa(tempCounter)
				result += temp

				temp = string(tempResult[k])
				tempCounter = 1
			} else {
				tempCounter++
			}
		}

		result += strconv.Itoa(tempCounter)
		result += temp

	}

	return result
}
