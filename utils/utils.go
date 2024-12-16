package utils

import (
	"fmt"
)

func RemoveAtIndex[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}
	result := []T{}
	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)

	return result
}

func ReplaceAtIndex[T any](s []T, index int, val T) []T {
	if index < 0 || index >= len(s) {
		return s
	}
	result := []T{}
	result = append(result, s[:index]...)
	result = append(result, val)
	result = append(result, s[index+1:]...)

	return result
}

func AddAtIndex[T any](s []T, index int, val T) []T {
	// fmt.Printf("adding %v at index %d  of %v\n", val, index, s)
	if index < 0 || index > len(s) {
		return s
	} else if index == len(s) {
		s = append(s, val)
		return s
	}

	result := []T{}
	result = append(result, s[:index]...)
	// fmt.Printf("first part: %v\n", s[:index])
	result = append(result, val)
	result = append(result, s[index:]...)
	// fmt.Printf("second part: %v\n", s[index:])
	// fmt.Printf("full result: %v\n", result)

	return result
}

func IsEven(i int) bool {
	if i%2 == 0 {
		return true
	}

	return false
}

func OutOfBounds[T any](board [][]T, x, y int) bool {
	if x < 0 || y < 0 || y > len(board)-1 || x > len(board[y])-1 {
		return true
	}

	return false
}

func PrintBoard[T any](board [][]T, x, y int, char T) {
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
