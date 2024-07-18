package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter a number:")

	var input string
	var result int

	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, r := range input {
		digit := int(r) - int('0')

		if digit < 0 || digit > 9 {
			fmt.Println("not a number")
			os.Exit(1)
		}

		result = result*10 + digit
	}

	fmt.Printf("Parsed input: %d\n", result)
}
