package main

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	sLen := len(s)
	if sLen == 0 {
		return 0
	}

	maxN := "2147483647"
	sign := ""
	if s[0] == '+' || s[0] == '-' {
		if sLen < 2 {
			return 0
		}

		if s[0] == '-' {
			maxN = "2147483648"
			sign = "-"
		}
		s = s[1:]
	}

	resultStr := ""
	for _, c := range s {
		if resultStr == "" && c == '0' {
			continue
		}
		if c < '0' || c > '9' {
			break
		}

		resultStr += string(c)
	}
	if resultStr == "" {
		return 0
	}

	sLen = len(resultStr)
	maxLen := len(maxN)
	if sLen > maxLen || (sLen == maxLen && resultStr > maxN) {
		resultStr = maxN
	}
	resultInt, _ := strconv.Atoi(sign + resultStr)
	return resultInt
}
