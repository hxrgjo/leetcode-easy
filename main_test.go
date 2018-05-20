package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {

	actual := reverse(12333)
	assert.Equal(t, 33321, actual)

	actual = reverse(10000)
	assert.Equal(t, 1, actual)

	actual = reverse(1534236469)
	assert.Equal(t, 0, actual)
}

func Test_isPalindrome(t *testing.T) {
	actual := isPalindrome(12321)
	assert.Equal(t, true, actual)
}

func Test_romanToInt(t *testing.T) {

	actual := romanToInt("III")
	assert.Equal(t, 3, actual)

	actual = romanToInt("MCMXCIV")
	assert.Equal(t, 1994, actual)
}

func Test_longestCommonPrefix(t *testing.T) {
	actual := longestCommonPrefix([]string{"flower", "flow", "flight"})
	assert.Equal(t, "fl", actual)

	actual = longestCommonPrefix([]string{"dogg", "racecar", "car"})
	assert.Equal(t, "", actual)
}

func Test_isValid(t *testing.T) {
	actual := isValid("()")
	assert.Equal(t, true, actual)

	actual = isValid("()[]{}")
	assert.Equal(t, true, actual)

	actual = isValid("([)]")
	assert.Equal(t, false, actual)

	actual = isValid("{[]}")
	assert.Equal(t, true, actual)

	actual = isValid("){")
	assert.Equal(t, false, actual)
}
