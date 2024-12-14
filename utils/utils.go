package utils

import (
	"fmt"
)

func removeAtIndex[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}
	result := []T{}
	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)

	return result
}

func outOfBounds[T any](board [][]T, x, y int) bool {
	if x < 0 || y < 0 || y > len(board)-1 || x > len(board[y])-1 {
		return true
	}

	return false
}

func printBoard[T any](board [][]T, x, y int, char T) {
	fmt.Println("--------")
	for boardY, row := range board {
		if y == boardY {
			temprow := make([]T, len(row))
			copy(temprow, row)
			temprow[x] = char
			fmt.Println(temprow)
		} else {
			fmt.Println(row)
		}
	}
}
