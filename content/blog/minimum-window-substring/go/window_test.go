package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstringWindow(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected string
	}{
		{
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
		{
			s:        "a",
			t:        "a",
			expected: "a",
		},
		{
			s:        "a",
			t:        "aa",
			expected: "",
		},
		{
			s:        "abcabc",
			t:        "abc",
			expected: "abc",
		},
		{
			s:        "cba",
			t:        "abc",
			expected: "cba",
		},
		{
			s:        "cb",
			t:        "abc",
			expected: "",
		},
		{
			s:        "acaba",
			t:        "abc",
			expected: "cab",
		},
		{
			s:        "abc",
			t:        "abcc",
			expected: "",
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", idx+1), func(t *testing.T) {
			result := minWindow(tc.s, tc.t)
			assert.Equal(t, tc.expected, result)
		})
	}
}
