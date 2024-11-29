package main

import "strings"

// https://leetcode.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := make([][]string, numRows)
	for i := 0; i < numRows; i += 1 {
		result[i] = make([]string, 0)
	}

	row := 0
	zig := true
	for _, c := range s {
		result[row] = append(result[row], string(c))

		if zig && row == (numRows-1) {
			zig = false
		} else if !zig && row == 0 {
			zig = true
		}

		if zig {
			row += 1
		} else {
			row -= 1
		}
	}

	resultStr := ""
	for _, r := range result {
		resultStr += strings.Join(r, "")
	}
	return resultStr
}
