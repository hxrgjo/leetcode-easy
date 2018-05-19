package main

func romanToInt(s string) int {

	result := 0

	num := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += num[string(s[i])]
			continue
		}

		if v, ok := num[string(s[i])+string(s[i+1])]; ok {
			result += v
			i++
			continue
		}

		result += num[string(s[i])]
	}

	return result
}

func romanToIntFaster(s string) int {
	if len(s) == 0 {
		return -1
	}

	var rst int
	var nextI byte

	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		default:
			return -1

		case 'I':
			if nextI == 'V' || nextI == 'X' {
				rst--
			} else {
				rst++
			}

		case 'X':
			if nextI == 'L' || nextI == 'C' {
				rst -= 10
			} else {
				rst += 10
			}

		case 'C':
			if nextI == 'D' || nextI == 'M' {
				rst -= 100
			} else {
				rst += 100
			}

		case 'V':
			rst += 5

		case 'L':
			rst += 50

		case 'D':
			rst += 500

		case 'M':
			rst += 1000

		}

		nextI = s[i]

	}

	return rst

}
