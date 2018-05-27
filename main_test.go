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

func Test_mergeTwoLists(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}

	l2 := new(ListNode)
	l2.Val = 1
	l2.Next = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
		},
	}

	actual := mergeTwoLists(l1, l2)

	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, expected, actual)

}

func Test_fibonacci(t *testing.T) {
	actual := fibonacci(10)
	assert.Equal(t, 10, actual)
}

func Test_removeDuplicates(t *testing.T) {

	actual := removeDuplicates([]int{1, 1, 2})
	assert.Equal(t, 2, actual)

	actual = removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	assert.Equal(t, 5, actual)
}

func Test_removeDuplicatesFaster(t *testing.T) {

	actual := removeDuplicatesFaster([]int{1, 1, 2})
	assert.Equal(t, 2, actual)

	actual = removeDuplicatesFaster([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	assert.Equal(t, 5, actual)
}

func Test_searchInsert(t *testing.T) {
	actual := searchInsert([]int{1, 3, 5, 6}, 5)
	assert.Equal(t, 2, actual)

	actual = searchInsert([]int{1, 3, 5, 6}, 2)
	assert.Equal(t, 1, actual)

	actual = searchInsert([]int{1, 3, 5, 6}, 7)
	assert.Equal(t, 4, actual)

	actual = searchInsert([]int{1, 3, 5, 6}, 0)
	assert.Equal(t, 0, actual)
}

func Test_countAndSay(t *testing.T) {
	// 1.     1
	// 2.     11
	// 3.     21
	// 4.     1211
	// 5.     111221

	actual := countAndSay(1)
	assert.Equal(t, "1", actual)

	actual = countAndSay(2)
	assert.Equal(t, "11", actual)

	actual = countAndSay(3)
	assert.Equal(t, "21", actual)

	actual = countAndSay(4)
	assert.Equal(t, "1211", actual)

	actual = countAndSay(5)
	assert.Equal(t, "111221", actual)

}
