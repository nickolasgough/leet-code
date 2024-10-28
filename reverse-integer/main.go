package main

import (
	"strconv"
)

// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	sign := ""
	maxN := "2147483647"
	xStr := strconv.Itoa(x)
	if x < 0 {
		sign = "-"
		maxN = "2147483648"
		xStr = xStr[1:]
	}

	maxLen := len(maxN)
	xLen := len(xStr)
	result := make([]rune, xLen)
	for i := xLen - 1; i >= 0; i -= 1 {
		result[xLen-1-i] = rune(xStr[i])
	}
	resultStr := string(result)
	if xLen >= maxLen && resultStr > maxN {
		// Could implement custom lexicographical comparison here.
		return 0
	}
	resultStr = sign + resultStr
	resultInt, _ := strconv.Atoi(resultStr)
	return resultInt
}
