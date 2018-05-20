package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	result := ""

	if len(strs) == 0 {
		return ""
	}

	for k := range strs[0] {
		for i := range strs {
			if !strings.HasPrefix(strs[i], string(strs[0][:k+1])) {
				return result
			}
		}
		result = string(strs[0][:k+1])
	}

	return result
}
