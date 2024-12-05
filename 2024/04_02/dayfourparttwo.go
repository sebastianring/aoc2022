package dayfourparttwo

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func DayFourPartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0

	board := []string{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		board = append(board, lineString)
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x:x+1] == "A" {
				sum += CheckX(board, x, y)
			}
		}
	}

	return sum, nil
}

func CheckX(board []string, x, y int) int {
	lettersOne := []string{"S", "M"}
	lettersTwo := []string{"S", "M"}

	for yOffset := -1; yOffset < 2; yOffset += 2 {
		for xOffset := -1; xOffset < 2; xOffset += 2 {
			currentY := y + yOffset
			currentX := x + xOffset
			if currentY < 0 || currentY > len(board)-1 || currentX < 0 || currentX > len(board[y])-1 {
				continue
			}

			if yOffset == 0 && xOffset == 0 {
				fmt.Printf("skipping 0, 0")
				continue
			}

			// first line: -1 -1, +1 +1
			// second line: 1 -1, -1  1
			var currentLetters *[]string
			if xOffset == -1 && yOffset == -1 || xOffset == 1 && yOffset == 1 {
				currentLetters = &lettersOne
			} else if xOffset == 1 && yOffset == -1 || xOffset == -1 && yOffset == 1 {
				currentLetters = &lettersTwo
			}

			index := slices.Index(*currentLetters, board[currentY][currentX:currentX+1])
			if index == -1 {
				return 0
			}

			*currentLetters = removeAtIndex(*currentLetters, index)
		}
	}

	if len(lettersOne) == 0 && len(lettersTwo) == 0 {
		return 1
	}

	return 0
}

func removeAtIndex(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		return s
	}
	result := []string{}
	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)

	return result
}
