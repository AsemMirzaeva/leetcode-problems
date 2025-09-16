package main

import "fmt"

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

func isValidSudoku(board [][]byte) bool {
	for i := range 9 {
		if !checkNineCells(0, i, 9, board) {
			return false
		}
		if !checkNineCells(i, 0, 1, board) {
			return false
		}
		if !checkNineCells((i%3)*3, (i/3)*3, 3, board) {
			return false
		}
	}
	return true
}

func checkNineCells(x, y, width int, board [][]byte) bool {
	var exist struct{}
	set := make(map[byte]struct{}, 9)
	for i := range 9 {
		dx, dy := i%width, i/width
		if v := board[y+dy][x+dx]; '1' <= v && v <= '9' {
			_, ok := set[v]
			if ok {
				return false
			}
			set[v] = exist
		}
	}
	return true
}

func main() {
	b1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(b1))
}
