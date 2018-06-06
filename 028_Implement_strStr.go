package main

func strStr(haystack string, needle string) int {
	var result = -1

	lenHaystack := len(haystack)
	lenNeedle := len(needle)

	if haystack == "" && needle == "" {
		return 0
	}

	if needle == "" {
		return -1
	}

	for i := 0; i <= lenHaystack-lenNeedle; i++ {
		if haystack[i] == needle[0] {
			tempIndex := i + 1
			tag := false

			for j := 1; j < lenNeedle; j++ {
				if haystack[tempIndex] != needle[j] {
					tag = true
					break
				}
				tempIndex++
			}

			if tag {
				continue
			}

			return i
		}
	}

	return result
}
