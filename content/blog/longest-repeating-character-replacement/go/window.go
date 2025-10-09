package main

func characterReplacement(s string, k int) int {
	if len(s) <= 1 || len(s)-1 <= k {
		return len(s)
	}

	count := make(map[byte]int)
	res, maxFrequent, left := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]] += 1

		maxFrequent = max(maxFrequent, count[s[right]])

		if (right-left+1)-maxFrequent > k {
			count[s[left]] -= 1
			left += 1
		}

		res = max(res, right-left+1)
	}

	return res
}
