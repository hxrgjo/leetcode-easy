package main

import "strconv"

func addBinary(a string, b string) string {
	var result string
	var next int

	aLen := len(a)
	bLen := len(b)

	for aLen-1 >= 0 || bLen-1 >= 0 {
		tempResult := next
		strTempresult := ""

		if aLen-1 >= 0 {
			tempa, _ := strconv.Atoi(string(a[aLen-1]))
			tempResult += tempa
			aLen--
		}

		if bLen-1 >= 0 {
			tempb, _ := strconv.Atoi(string(b[bLen-1]))
			tempResult += tempb
			bLen--
		}

		switch tempResult {
		case 0:
			strTempresult = "0"
			next = 0
		case 1:
			strTempresult = "1"
			next = 0
		case 2:
			strTempresult = "0"
			next = 1
		case 3:
			strTempresult = "1"
			next = 1
		}

		result = strTempresult + result
	}

	if next == 1 {
		result = "1" + result
	}

	return result
}
