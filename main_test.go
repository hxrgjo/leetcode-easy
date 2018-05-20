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
