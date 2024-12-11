package main

func longestValidParentheses(s string) int {
	nLeftBraces := 0
	nRightBraces := 0
	maxN := 0
	prevC := ""
	for _, rawC := range s {
		c := string(rawC)
		if c == "(" {
			nLeftBraces += 1
		} else if c == ")" {

		}
	}
}
