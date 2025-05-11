package main

import (
	"math/big"
	"strconv"
	"strings"
)

func multiplyRecursion(num1 string, num2 string) string {
	if len(num1) != len(num2) {
		n := max(len(num1), len(num2))

		if len(num1) < len(num2) {
			num1 = strings.Repeat("0", n-len(num1)) + num1
		} else {
			num2 = strings.Repeat("0", n-len(num2)) + num2
		}
	}

	if len(num1)%2 != 0 {
		num1 = "0" + num1
	}

	if len(num2)%2 != 0 {
		num2 = "0" + num2
	}

	return recursion(num1, num2)
}

func recursion(num1 string, num2 string) string {
	if (len(num1) == 1) && (len(num2) == 1) {
		i, _ := strconv.Atoi(num1)
		j, _ := strconv.Atoi(num2)

		return strconv.Itoa(i * j)
	}

	divider := len(num1) / 2
	// coef = 10^divider
	coef := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(divider)), nil)
	coefSquared := new(big.Int).Mul(coef, coef)

	// abcd = ab*100 + cd
	// efgh = ef*100 + gh

	// abcd*efgh = (ab*100 + cd)(ef*100 + gh) = (100*100*ab*ef + 100*ab*gh + 100*cd*ef + cd*gh)
	ab, cd := num1[:divider], num1[divider:]
	ef, gh := num2[:divider], num2[divider:]

	// first = ab*ef
	first := new(big.Int)
	first.SetString(recursion(ab, ef), 10)

	// second = ab*gh
	second := new(big.Int)
	second.SetString(recursion(ab, gh), 10)

	// second = cd*ef
	third := new(big.Int)
	third.SetString(recursion(cd, ef), 10)

	// fourth = cd*gh
	fourth := new(big.Int)
	fourth.SetString(recursion(cd, gh), 10)

	// result = coef^2 * ab * ef  + coef * ((ab * gh) + (cd * ef)) + cd * gh

	result := new(big.Int).Mul(first, coefSquared)

	result.Add(
		result, new(big.Int).Mul(second.Add(second, third), coef),
	).Add(result, fourth)

	return result.String()
}
