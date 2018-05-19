package main

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	var beforeReverseNumber = x
	var reverseNumber = 0

	for x != 0 {
		reverseNumber = 10*reverseNumber + (x % 10)
		x = x / 10
	}

	if reverseNumber == beforeReverseNumber {
		return true
	}

	return false
}
