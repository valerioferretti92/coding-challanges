package main

import "fmt"

func main() {
	fmt.Println(convert1("paypalishiring", 3))
}

func convert1(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	currRow, dir := 0, -1

	for i := 0; i < len(s); i++ {
		rows[currRow] += string(s[i])

		if currRow == 0 || currRow == numRows-1 {
			dir *= -1
		}

		currRow += dir
	}

	var result string
	for _, row := range rows {
		result += row
	}
	return result
}

func convert2(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}

	var zigzag string = ""
	for row := 0; row < numRows; row++ {
		index := row
		zigzag = add(zigzag, s, index)

		for index < len(s) {
			if row == 0 || row == numRows-1 {
				index = index + 2*(numRows-1)
				zigzag = add(zigzag, s, index)
			} else {
				index = index + 2*(numRows-row-1)
				zigzag = add(zigzag, s, index)
				index = index + 2*row
				zigzag = add(zigzag, s, index)
			}
		}
	}
	return zigzag
}

func add(zigzag, s string, index int) string {
	if index < len(s) {
		zigzag = zigzag + string(s[index])
	}
	return zigzag
}
