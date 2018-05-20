package main

func isValid(s string) bool {
	var result = true

	if s == "" {
		return true
	}

	if len(s) < 2 {
		return false
	}

	if len(s)%2 != 0 {
		return false
	}

	leftSymbols := map[string]string{"(": "", "{": "", "[": ""}
	symbolMaps := map[string]string{")": "(", "}": "{", "]": "["}

	if _, ok := leftSymbols[string(s[0])]; !ok {
		return false
	}

	tempArray := []string{}

	for i := range s {
		if _, ok := leftSymbols[string(s[i])]; ok {
			tempArray = append(tempArray, string(s[i]))
		} else {
			symbol := tempArray[len(tempArray)-1]
			m := symbolMaps[string(s[i])]

			if symbol != m {
				return false
			}

			tempArray = tempArray[:len(tempArray)-1]
		}
	}

	if len(tempArray) > 0 {
		return false
	}

	return result
}
