package main

// https://leetcode.com/problems/integer-to-roman/

type romanNumeral struct {
	Letter string
	Value  int
}

func intToRoman(n int) string {
	resultStr := ""
	remainder := n
	romanNumerals := []*romanNumeral{
		{Letter: "M", Value: 1000},
		{Letter: "D", Value: 500},
		{Letter: "C", Value: 100},
		{Letter: "L", Value: 50},
		{Letter: "X", Value: 10},
		{Letter: "V", Value: 5},
		{Letter: "I", Value: 1},
	}
	for i, rn := range romanNumerals {
		isDivisor := i%2 == 0
		if !isDivisor {
			continue
		}

		nLetters := int(remainder / rn.Value)
		if nLetters > 0 {
			remainder = remainder % rn.Value

			lettersToAdd := ""
			if nLetters == 4 {
				otherRn := romanNumerals[i-1]
				lettersToAdd = rn.Letter + otherRn.Letter
			} else if nLetters == 9 {
				otherRn := romanNumerals[i-2]
				lettersToAdd = rn.Letter + otherRn.Letter
			} else if nLetters >= 5 {
				otherRn := romanNumerals[i-1]
				lettersToAdd += otherRn.Letter
				for range nLetters - 5 {
					lettersToAdd += rn.Letter
				}
			} else {
				for range nLetters {
					resultStr += rn.Letter
				}
			}
			resultStr += lettersToAdd
		}
	}
	return resultStr
}
