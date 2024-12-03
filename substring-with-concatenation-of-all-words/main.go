package main

import "fmt"

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	fmt.Printf("%+v\n", findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	sLen := len(s)
	wordsLen := len(words)
	wordLen := len(words[0])
	result := make([]int, 0)
	for i := 0; i < sLen; i += 1 {
		start := i
		end := i + wordLen
		wordMap := mapifyWords(words)
		wordsToFind := wordsLen
		for end <= sLen {
			subWord := s[start:end]
			value, ok := wordMap[subWord]
			if ok && value > 0 {
				wordMap[subWord] = value - 1
				wordsToFind -= 1
			} else {
				wordMap = mapifyWords(words)
				wordsToFind = wordsLen
				break
			}
			if wordsToFind == 0 {
				result = append(result, i)
			}

			start += wordLen
			end += wordLen
		}
	}
	return result
}

func mapifyWords(words []string) map[string]int {
	wordMap := make(map[string]int)
	for _, word := range words {
		value, ok := wordMap[word]
		if ok {
			wordMap[word] = value + 1
			continue
		}
		wordMap[word] = 1
	}
	return wordMap
}
