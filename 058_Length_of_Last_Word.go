package main

func lengthOfLastWord(s string) int {
	var result int = 0

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			result = 0
		} else {
			result++
		}
	}

	return result
}

func lengthOfLastWordLeetCodeAnswer(s string) int {
	var result int = 0
	var tempresult int = 0

	for index := range s {
		if string(s[index]) == " " {
			tempresult = 0
		} else {
			tempresult++
			result = tempresult
		}
	}

	return result
}
