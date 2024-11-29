package main

import (
	"strconv"
)

// https://leetcode.com/problems/integer-to-roman/

type romanNumeral struct {
	Letter string
	Value  int
}

func intToRoman(n int) string {
	resultStr := ""
	nStr := strconv.Itoa(n)
	nLen := len(nStr)
	romanNumerals := []*romanNumeral{
		{Letter: "I", Value: 1},
		{Letter: "V", Value: 5},
		{Letter: "X", Value: 10},
		{Letter: "L", Value: 50},
		{Letter: "C", Value: 100},
		{Letter: "D", Value: 500},
		{Letter: "M", Value: 1000},
	}
	for i := 0; i < nLen; i += 1 {
		r := 2 * i
		c := nStr[nLen-1-i]
		if c == '0' {
			continue
		}
		cInt, _ := strconv.Atoi(string(c))
		rn := romanNumerals[r]
		lettersToAdd := ""
		if cInt == 4 {
			nextRn := romanNumerals[r+1]
			lettersToAdd = rn.Letter + nextRn.Letter
		} else if cInt == 9 {
			nextRn := romanNumerals[r+2]
			lettersToAdd = rn.Letter + nextRn.Letter
		} else if cInt >= 5 {
			nextRn := romanNumerals[r+1]
			lettersToAdd += nextRn.Letter
			for range cInt - 5 {
				lettersToAdd += rn.Letter
			}
		} else {
			for range cInt {
				lettersToAdd += rn.Letter
			}
		}
		resultStr = lettersToAdd + resultStr
	}
	return resultStr
}
