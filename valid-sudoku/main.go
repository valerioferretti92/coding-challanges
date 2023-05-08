package main

import (
	"fmt"
	"math"
)

func main() {
	sudoku := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(sudoku))
}

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]bool)
	cols := make(map[int]map[byte]bool)
	sqrs := make(map[int]map[byte]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n := board[i][j]
			if string(n) == "." {
				continue
			}

			if rows[i] == nil {
				rows[i] = make(map[byte]bool)
			}
			if _, ok := rows[i][n]; ok {
				return false
			} else {
				rows[i][n] = true
			}

			if cols[j] == nil {
				cols[j] = make(map[byte]bool)
			}
			if _, ok := cols[j][n]; ok {
				return false
			} else {
				cols[j][n] = true
			}

			sqridx := int(math.Floor(float64(i)/3) + 3*math.Floor(float64(j)/3))
			if sqrs[sqridx] == nil {
				sqrs[sqridx] = make(map[byte]bool)
			}
			if _, ok := sqrs[sqridx][n]; ok {
				return false
			} else {
				sqrs[sqridx][n] = true
			}
		}
	}
	return true
}
