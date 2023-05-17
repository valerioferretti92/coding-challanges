package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("123", "456"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	maxLength := 0
	partials := make([]string, 0)

	// Building partial multiplications as strings
	for i := 0; i < len(num1); i++ {
		carry := 0
		pos1 := len(num1) - i - 1
		val1, _ := strconv.Atoi(string(num1[pos1]))
		partials = append(partials, zeros(i))

		for j := 0; j < len(num2); j++ {
			pos2 := len(num2) - j - 1
			val2, _ := strconv.Atoi(string(num2[pos2]))

			partial := val1*val2 + carry
			partials[i] = fmt.Sprintf("%d%s", partial%10, partials[i])
			carry = partial / 10
		}

		if carry != 0 {
			partials[i] = fmt.Sprintf("%d%s", carry, partials[i])
		}

		if len(partials[i]) > maxLength {
			maxLength = len(partials[i])
		}
	}

	// Prefixing partials with 0
	for i := 0; i < len(partials); i++ {
		length := maxLength - len(partials[i])
		partials[i] = zeros(length) + partials[i]
	}

	// Summing partial multiplications
	carry := 0
	result := ""
	for i := maxLength - 1; i >= 0; i-- {
		partial := 0
		for j := 0; j < len(partials); j++ {
			d, _ := strconv.Atoi(string(partials[j][i]))
			partial = partial + d
		}

		digit := (partial + carry) % 10
		carry = (partial + carry) / 10
		result = fmt.Sprintf("%d%s", digit, result)
	}

	if carry != 0 {
		result = fmt.Sprintf("%d%s", carry, result)
	}

	return result
}

func zeros(i int) string {
	initilizer := ""

	for i > 0 {
		initilizer = initilizer + "0"
		i--
	}

	return initilizer
}
