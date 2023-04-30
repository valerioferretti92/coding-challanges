package main

import (
	"fmt"
	"math"
	"strconv"
)

// 1010
func main() {
	fmt.Println(reverse(100))
}

func reverse(x int) int {
	var sign, abs int

	if x < 0 {
		sign = -1
		abs = x * sign
	} else {
		sign = 1
		abs = x
	}

	var str, inv string
	str = fmt.Sprintf("%d", abs)
	for _, char := range str {
		inv = string(char) + inv
	}

	if sign == -1 {
		inv = "-" + inv
	}

	x, err := strconv.Atoi(inv)
	if err != nil {
		return 0
	}
	if x > int(math.Pow(2, 31))-1 || x < -int(math.Pow(2, 31)) {
		return 0
	}
	return x
}

func bitwsieReverse(x int) int {
	xbin := make([]int, 32)

	for i := 0; i < 32; i++ {
		pow := int(math.Pow(2, float64(i)))
		xbin[i] = (x & pow) / pow

		if xbin[i] == 0 {
			xbin[i] = 1
		} else {
			xbin[i] = 0
		}
	}

	var result int = 0
	for i := 0; i < 32; i++ {
		pow := int(math.Pow(2, float64(i)))
		result += xbin[i] * pow
	}

	return result
}
