package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a number:")

	var input int
	var result string

	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for input > 0 {
		digit := input % 10

		r := digit + int(rune('0'))
		result = string(rune(r)) + result

		input = input - digit
		input = input / 10
	}

	fmt.Println(result)
}
