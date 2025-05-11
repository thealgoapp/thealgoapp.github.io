package main

import (
	"math/big"
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	// x = ab = 10*a + b
	// y = cd = c*10 + d
	// x * y = (10*a + b) * (10*c + d)
	// x * y = 100 * ac + 10 * (*ad + bc*) + bd

	// (a + b) * (c + d) = ac + ad + bc + bd
	// *ad + bc* = (a + b) * (c + d) - ac - bd

	// x * y = 100 * ac + 10 * ((a + b) * (c + d) - ac - bd) + bd

	return recursionKaratsuba(num1, num2)
}

func recursionKaratsuba(num1 string, num2 string) string {
	if (len(num1) == 1) && (len(num2) == 1) {
		i, _ := strconv.Atoi(num1)
		j, _ := strconv.Atoi(num2)

		return strconv.Itoa(i * j)
	}

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

	divider := len(num1) / 2
	// coef = 10^divider
	coef := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(divider)), nil)
	coefSquared := new(big.Int).Mul(coef, coef)

	// ab = a*10 + b
	// cd = c*10 + d

	// ab*cd = 100 * ac + 10 * ((a + b) * (c + d) - ac - bd) + bd
	// ab*cd = 100 * ac + bd + 10 * ((a + b) * (c + d) - ac - bd)
	a, b := num1[:divider], num1[divider:]
	c, d := num2[:divider], num2[divider:]

	// first = ac
	first := new(big.Int)
	first.SetString(recursionKaratsuba(a, c), 10)

	// fourth = b*d
	fourth := new(big.Int)
	fourth.SetString(recursionKaratsuba(b, d), 10)

	// 100 * ac + bd
	result := new(big.Int).Add(new(big.Int).Mul(first, coefSquared), fourth)

	bigA := new(big.Int)
	bigA.SetString(a, 10)

	bigB := new(big.Int)
	bigB.SetString(b, 10)

	bigC := new(big.Int)
	bigC.SetString(c, 10)

	bigD := new(big.Int)
	bigD.SetString(d, 10)

	// (a + b) * (c + d)
	subMultiply := new(big.Int)
	subMultiply.SetString(recursionKaratsuba(
		new(big.Int).Add(bigA, bigB).String(),
		new(big.Int).Add(bigC, bigD).String(),
	), 10)

	subMultiply.Add(
		subMultiply, new(big.Int).Neg(first),
	).Add(
		subMultiply, new(big.Int).Neg(fourth),
	)

	result.Add(
		result, subMultiply.Mul(subMultiply, coef),
	)

	return result.String()
}

func multiplyOptimised(num1 string, num2 string) string {
	// x = ab = 10*a + b
	// y = cd = c*10 + d
	// x * y = (10*a + b) * (10*c + d)
	// x * y = 100 * ac + 10 * (*ad + bc*) + bd

	// (a + b) * (c + d) = ac + ad + bc + bd
	// *ad + bc* = (a + b) * (c + d) - ac - bd

	// x * y = 100 * ac + 10 * ((a + b) * (c + d) - ac - bd) + bd

	if len(num1) <= 15 && len(num2) <= 15 {
		return recursionNative(num1, num2)
	}

	return recursionKaratsubaOptimised(num1, num2)
}

func recursionNative(num1 string, num2 string) string {
	if (len(num1) == 1) && (len(num2) == 1) {
		i := int(num1[0] - '0')
		j := int(num2[0] - '0')

		return strconv.Itoa(i * j)
	}

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

	divider := len(num1) / 2
	a, b := num1[:divider], num1[divider:]
	c, d := num2[:divider], num2[divider:]

	first := recursionKaratsubaOptimised(a, c)

	second := recursionKaratsubaOptimised(a, d)

	third := recursionKaratsubaOptimised(b, c)

	// fourth = b*d
	fourth := recursionKaratsubaOptimised(b, d)

	res := sumStrings(
		sumStrings(
			first+strings.Repeat("0", 2*divider),
			sumStrings(second, third)+strings.Repeat("0", divider),
		),
		fourth,
	)

	res = strings.TrimLeft(res, "0")
	if res == "" {
		return "0"
	}

	return res
}

func recursionKaratsubaOptimised(num1 string, num2 string) string {
	if (len(num1) == 1) && (len(num2) == 1) {
		i := int(num1[0] - '0')
		j := int(num2[0] - '0')

		return strconv.Itoa(i * j)
	}

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

	divider := len(num1) / 2
	// coef = 10^divider

	// ab = a*10 + b
	// cd = c*10 + d

	// ab*cd = 100 * ac + 10 * ((a + b) * (c + d) - ac - bd) + bd
	// ab*cd = 100 * ac + bd + 10 * ((a + b) * (c + d) - ac - bd)
	a, b := num1[:divider], num1[divider:]
	c, d := num2[:divider], num2[divider:]

	first := recursionKaratsubaOptimised(a, c)

	// fourth = b*d
	fourth := recursionKaratsubaOptimised(b, d)

	res := sumStrings(
		first+strings.Repeat("0", 2*divider),
		fourth,
	)

	sub := subtractStrings(
		subtractStrings(
			recursionKaratsubaOptimised(
				sumStrings(a, b),
				sumStrings(c, d),
			),
			first,
		),
		fourth,
	)

	res = strings.TrimLeft(sumStrings(res, sub+strings.Repeat("0", divider)), "0")
	if res == "" {
		return "0"
	}

	return res
}

func sumStrings(a, b string) string {
	main, complementary := a, b
	if len(main) < len(complementary) {
		main, complementary = complementary, main
	}

	carry := 0
	result := make([]rune, len(main)+1)
	i, j := len(main)-1, len(complementary)-1
	for j >= 0 {
		digitA, digitB := int(main[i]-'0'), int(complementary[j]-'0')
		sum := carry + digitA + digitB
		result[i+1] = rune(sum%10) + '0'
		carry = sum / 10

		i--
		j--
	}

	for ; i >= 0; i-- {
		digit := int(main[i] - '0')
		sum := carry + digit
		result[i+1] = rune(sum%10) + '0'
		carry = sum / 10
	}

	if carry != 0 {
		result[0] = '1'
	} else {
		result = result[1:]
	}

	return string(result)
}

func subtractStrings(a, b string) string {
	a, b = strings.TrimLeft(a, "0"), strings.TrimLeft(b, "0")

	result := make([]rune, len(a))
	borrow := 0
	i, j := len(a)-1, len(b)-1

	for j >= 0 {
		digitA := int(a[i] - '0')
		digitB := int(b[j] - '0')
		diff := digitA - digitB - borrow

		if diff < 0 {
			diff += 10
			borrow = 1
		} else {
			borrow = 0
		}

		result[i] = rune(diff + '0')
		i--
		j--
	}

	for ; i >= 0; i-- {
		digitA := int(a[i] - '0')
		diff := digitA - borrow

		if diff < 0 {
			diff += 10
			borrow = 1
		} else {
			borrow = 0
		}

		result[i] = rune(diff + '0')
	}

	resStr := strings.TrimLeft(string(result), "0")
	if resStr == "" {
		return "0"
	}
	return resStr
}
