package main

func characterReplacementIncorrect(s string, k int) int {
	if len(s) <= 1 || len(s)-1 <= k {
		return len(s)
	}

	maxLength := 1
	for i := 0; i < len(s)-1; i++ {
		stack := make([]byte, 0, len(s))
		stack = append(stack, s[i])
		currentK := k
		for j := i + 1; j < len(s); j++ {
			if s[j] == stack[0] {
				stack = append(stack, s[j])

				continue
			}

			if currentK == 0 {
				break
			}

			stack = append(stack, s[j])
			currentK--
		}

		maxLength = max(maxLength, len(stack))
		if maxLength > len(s[i:]) {
			break
		}
	}

	return maxLength
}
