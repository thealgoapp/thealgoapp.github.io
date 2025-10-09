package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringWindow(t *testing.T) {
	testCases := []struct {
		s           string
		k           int
		expectedLen int
	}{
		{
			s:           "ABAB",
			k:           2,
			expectedLen: 4,
		},
		{
			s:           "AABABBA",
			k:           1,
			expectedLen: 4,
		},
		{
			s:           "ABC",
			k:           5,
			expectedLen: 3,
		},
		{
			s:           "A",
			k:           2,
			expectedLen: 1,
		},
		{
			s:           "ABA",
			k:           2,
			expectedLen: 3,
		},
		{
			s:           "AAA",
			k:           2,
			expectedLen: 3,
		},
		{
			s:           "ABBA",
			k:           1,
			expectedLen: 3,
		},
		{
			s:           "ABBA",
			k:           2,
			expectedLen: 4,
		},
		{
			s:           "AA",
			k:           0,
			expectedLen: 2,
		},
		{
			s:           "AA",
			k:           1,
			expectedLen: 2,
		},
		// (!) missed corner cases
		{
			s:           "ABBB",
			k:           1,
			expectedLen: 4,
		},
		{
			s:           "BAAAB",
			k:           2,
			expectedLen: 5,
		},
		{
			s:           "BAAABBAAAAAAAB",
			k:           2,
			expectedLen: 12,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", idx+1), func(t *testing.T) {
			result := characterReplacement(tc.s, tc.k)
			assert.Equal(t, tc.expectedLen, result)
		})
	}
}
