package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAreaWindow(t *testing.T) {
	testCases := []struct {
		height         []int
		expectedVolume int
	}{
		{
			height: []int{
				1, 8, 6, 2, 5, 4, 8, 3, 7,
			},
			expectedVolume: 49,
		},
		{
			height: []int{
				1, 1,
			},
			expectedVolume: 1,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Example %d", idx+1), func(t *testing.T) {
			result := maxAreaWindow(tc.height)
			assert.Equal(t, tc.expectedVolume, result)
		})
	}
}
