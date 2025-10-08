package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringWindow(t *testing.T) {
	testCases := []struct {
		s           string
		expectedLen int
	}{
		{
			s:           "abcabcbb",
			expectedLen: 3,
		},
		{
			s:           "bbbbb",
			expectedLen: 1,
		},
		{
			s:           "pwwkew",
			expectedLen: 3,
		},
		{
			s:           "au",
			expectedLen: 2,
		},
		{
			s:           "aab",
			expectedLen: 2,
		},
		{
			s:           "abba",
			expectedLen: 2,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", idx+1), func(t *testing.T) {
			result := lengthOfLongestSubstringWindow(tc.s)
			assert.Equal(t, tc.expectedLen, result)
		})
	}
}
