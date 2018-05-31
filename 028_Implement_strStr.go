package main

func strStr(haystack string, needle string) int {
	var result = -1

	lenHaystack := len(haystack)
	lenNeedle := len(needle)

	for i := 0; i <= lenHaystack-lenNeedle; i++ {
		if haystack[i:i+lenNeedle] == needle {
			return i
		}
	}

	return result
}
