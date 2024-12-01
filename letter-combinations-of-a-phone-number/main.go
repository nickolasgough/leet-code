package main

var (
	digitsMap = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	digitsLen := len(digits)
	result := make([]string, 0)
	if digitsLen < 1 {
		return result
	}

	letters := make([][]string, digitsLen)
	for i, digit := range digits {
		letters[i] = digitsMap[string(digit)]
	}
	indices := make([]int, digitsLen)
	for {
		combo := ""
		for i, index := range indices {
			combo += letters[i][index]
		}
		result = append(result, combo)

		i := 0
		maxI := digitsLen - 1
		for {
			index := indices[i]
			index += 1
			maxIndex := len(letters[i]) - 1
			if index > maxIndex {
				indices[i] = 0
				i += 1
				if i > maxI {
					return result
				}
				continue
			}
			indices[i] = index
			break
		}
	}
}
