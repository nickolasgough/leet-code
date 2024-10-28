package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	if strLen == 0 || strLen == 1 {
		return strLen
	}

	encounteredChars := make(map[rune]int)
	currPosition := 0
	longestSubstr := 0
	for i, c := range s {
		charPosition, ok := encounteredChars[c]
		if ok {
			if charPosition >= currPosition {
				currPosition = charPosition + 1
			}
			encounteredChars[c] = i
		} else {
			encounteredChars[c] = i
		}

		currSubstr := i - (currPosition - 1)
		if currSubstr > longestSubstr {
			longestSubstr = currSubstr
		}
	}
	return longestSubstr
}
