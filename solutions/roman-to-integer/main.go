package main

type romanNumeral struct {
	Letter  rune
	Value   int
	OtherRn rune
}

func romanToInt(s string) int {
	sLen := len(s)
	resultInt := 0
	romanNumerals := map[rune]*romanNumeral{
		'I': {Letter: 'I', Value: 1},
		'V': {Letter: 'V', Value: 5, OtherRn: 'I'},
		'X': {Letter: 'X', Value: 10, OtherRn: 'I'},
		'L': {Letter: 'L', Value: 50, OtherRn: 'X'},
		'C': {Letter: 'C', Value: 100, OtherRn: 'X'},
		'D': {Letter: 'D', Value: 500, OtherRn: 'C'},
		'M': {Letter: 'M', Value: 1000, OtherRn: 'C'},
	}
	for i := sLen - 1; i >= 0; i -= 1 {
		c := rune(s[i])
		rn := romanNumerals[c]
		resultInt += rn.Value

		if rn.OtherRn != 0 && i > 0 {
			otherRn := romanNumerals[rn.OtherRn]
			if rune(s[i-1]) == otherRn.Letter {
				i -= 1
				resultInt -= otherRn.Value
			}
		}
	}
	return resultInt
}
