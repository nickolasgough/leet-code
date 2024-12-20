package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	prevRLE := countAndSay(n - 1)
	rle := ""
	prevC := string(prevRLE[0])
	cCount := 1
	for _, r := range prevRLE[1:] {
		c := string(r)
		if prevC == c {
			cCount += 1
			continue
		}

		rle += strconv.Itoa(cCount) + prevC
		prevC = c
		cCount = 1
	}
	rle += strconv.Itoa(cCount) + prevC
	return rle
}
