package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_multiplyRecursion(t *testing.T) {
	cases := []struct {
		name     string
		num1     string
		num2     string
		expected string
	}{
		{
			name:     "12*34",
			num1:     "12",
			num2:     "34",
			expected: "408",
		},
		{
			name:     "123×456",
			num1:     "123",
			num2:     "456",
			expected: "56088",
		},
		{
			name:     "1234*5678",
			num1:     "1234",
			num2:     "5678",
			expected: "7006652",
		},
		{
			name:     "123×4567",
			num1:     "123",
			num2:     "4567",
			expected: "561741",
		},
		{
			name:     "3141592653589793238462643383279502884197169399375105820974944592×2718281828459045235360287471352662497757247093699959574966967627",
			num1:     "3141592653589793238462643383279502884197169399375105820974944592",
			num2:     "2718281828459045235360287471352662497757247093699959574966967627",
			expected: "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			res := multiplyRecursion(testCase.num1, testCase.num2)
			assert.Equal(t, testCase.expected, res)
		})
	}
}

// Benchmark_multiplyRecursion-10    	    1548	    763014 ns/op
func Benchmark_multiplyRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplyRecursion(
			"3141592653589793238462643383279502884197169399375105820974944592",
			"2718281828459045235360287471352662497757247093699959574966967627",
		)
	}
}
