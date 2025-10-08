package main

import "fmt"

func lengthOfLongestSubstringWindow(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxSeqLen, i, j := 1, 0, 1
	visitedChars := map[byte]int{
		s[i]: i,
	}

	for j < len(s) && i < j {
		if idx, ok := visitedChars[s[j]]; ok && idx >= i {
			maxSeqLen = max(maxSeqLen, j-i)
			i = idx + 1
		}

		visitedChars[s[j]] = j
		j++

		if j == len(s) {
			fmt.Println(i)
			maxSeqLen = max(maxSeqLen, j-i)
		}
	}

	if len(visitedChars) == len(s) {
		return len(visitedChars)
	}

	return maxSeqLen
}
