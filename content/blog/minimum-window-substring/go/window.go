package main

func minWindow(s string, t string) string {
	expectedFrequencies := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		expectedFrequencies[t[i]]++
	}

	substr := ""
	left := 0
	currentFrequencies := make(map[byte]int)
	acceptableCriteriaCount := 0
	for right := 0; right < len(s); right++ {
		if _, exists := expectedFrequencies[s[right]]; !exists {
			continue
		}
		currentFrequencies[s[right]]++
		if acceptableCriteriaCount < len(expectedFrequencies) && currentFrequencies[s[right]] == expectedFrequencies[s[right]] {
			acceptableCriteriaCount++
		}

		if acceptableCriteriaCount == len(expectedFrequencies) {
			_, exists := expectedFrequencies[s[left]]
			for left < right && (!exists || currentFrequencies[s[left]] > expectedFrequencies[s[left]]) {
				currentFrequencies[s[left]]--
				left++
				_, exists = expectedFrequencies[s[left]]
			}

			if len(substr) == 0 || len(substr) > right-left+1 {
				substr = s[left : right+1]
			}
		}
	}

	return substr
}
