package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	// Check rows
	for i := range 9 {
		seen := make(map[byte]bool)
		for j := range 9 {
			if board[i][j] != '.' {
				if seen[board[i][j]] {
					return false
				}
				seen[board[i][j]] = true
			}
		}
	}

	// Check columns
	for i := range 9 {
		seen := make(map[byte]bool)
		for j := range 9 {
			if board[j][i] != '.' {
				if seen[board[j][i]] {
					return false
				}
				seen[board[j][i]] = true
			}
		}
	}

	// Check 3x3 sub-boxes
	for boxRow := range 3 {
		for boxCol := range 3 {
			seen := make(map[byte]bool)
			for row := range 3 {
				for col := range 3 {
					cell := board[boxRow*3+row][boxCol*3+col]
					if cell != '.' {
						if seen[cell] {
							return false
						}
						seen[cell] = true
					}
				}
			}
		}
	}

	return true
}

func main() {
	board1 := [][]byte{
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
	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Printf("Board 1: %v\n", isValidSudoku(board1))
	fmt.Printf("Board 2: %v\n", isValidSudoku(board2))
}
