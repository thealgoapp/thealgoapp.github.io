package main

func lengthOfLongestSubstringBruteforce(s string) int {
	maxSeqLen := 0
	for i := 0; i < len(s); i++ {
		visitedChars := make(map[byte]struct{})
		visitedChars[s[i]] = struct{}{}
		seqLen := 1
		for j := i + 1; j < len(s); j++ {
			if _, ok := visitedChars[s[j]]; ok {
				break
			}

			visitedChars[s[j]] = struct{}{}
			seqLen++
		}

		maxSeqLen = max(maxSeqLen, seqLen)
	}

	return maxSeqLen
}
