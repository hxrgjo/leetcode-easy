package main

func reverseVowels(s string) string {
	first := 0
	last := len(s) - 1
	ret := make([]byte, len(s))

	for first <= last {
		if isVowels(s[first]) && isVowels(s[last]) {
			ret[first], ret[last] = s[last], s[first]
			first++
			last--
		} else if !isVowels(s[first]) {
			ret[first] = s[first]
			first++
		} else if !isVowels(s[last]) {
			ret[last] = s[last]
			last--
		}
	}

	return string(ret)
}

func isVowels(c byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _, v := range vowels {
		if c == v {
			return true
		}
	}

	return false
}
