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
