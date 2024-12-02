package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", generateParenthesis(4))
}

func generateParenthesis(n int) []string {
	resultMap := generateParenthesisHelper(n)
	result := make([]string, 0)
	for k, _ := range resultMap {
		result = append(result, k)
	}
	return result
}

func generateParenthesisHelper(n int) map[string]bool {
	if n == 1 {
		return map[string]bool{
			"()": true,
		}
	}
	result := make(map[string]bool)
	parentheses := generateParenthesisHelper(n - 1)
	for p, _ := range parentheses {
		newPs := []string{
			"(" + p + ")",
			"()" + p,
			p + "()",
		}
		for _, p := range newPs {
			result[p] = true
		}
	}

	for tempN := n - 1; tempN > 0; tempN -= 1 {
		temp1 := generateParenthesisHelper(tempN)
		temp2 := generateParenthesisHelper(n - tempN)
		for k1, _ := range temp1 {
			for k2, _ := range temp2 {
				newPs := []string{
					k1 + k2,
					k2 + k1,
				}
				for _, p := range newPs {
					result[p] = true
				}
			}
		}
	}
	return result
}
