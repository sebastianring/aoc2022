package dayfourpartone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const XMAS = "XMAS"

var saveboard []string
var resultsTracker map[string]int

func DayFourPartOne(filename string) (int, error) {
	resultsTracker = map[string]int{}
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
		saveboard = append(saveboard, strings.Repeat("O", len(board[y])))
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x:x+1] == "X" {
				sum += CheckDirections(board, x, y)
			}
		}
	}

	// for _, yRow := range saveboard {
	// 	fmt.Println(yRow)
	// }

	for y := 0; y < len(saveboard); y++ {
		for x := 0; x < len(saveboard[y]); x++ {
			fmt.Print(saveboard[y][x : x+1])
		}
		fmt.Print("\n")
	}

	for k, v := range resultsTracker {
		fmt.Printf("offsets: %s - pts: %d\n", k, v)
	}

	return sum, nil
}

func CheckDirections(board []string, x, y int) int {
	result := 0

	for yOffset := -1; yOffset < 2; yOffset++ {
		for xOffset := -1; xOffset < 2; xOffset++ {
			// || currentX < 0 || currentX > len(board[currentY])-1 || currentY < 0 || currentY > len(board)-1
			maxY := y + (yOffset * 3)
			maxX := x + (xOffset * 3)
			if maxY < 0 || maxY > len(board)-1 || maxX < 0 || maxX > len(board[y])-1 {
				continue
			}

			if yOffset == 0 && xOffset == 0 {
				// fmt.Printf("skipping 0, 0")
				continue
			}

			for i := 1; i < 4; i++ {
				currentX := x + (xOffset * i)
				currentY := y + (yOffset * i)

				// fmt.Printf("currentX: %d - currentY: %d\n", currentX, currentY)

				if board[currentY][currentX:currentX+1] != XMAS[i:i+1] {
					break
				}

				if i == 3 {
					result++
					resultsTracker[TrackerString(xOffset, yOffset)] += 1
					AddValue(currentX, currentY, xOffset, yOffset, board[currentY][currentX:currentX+1])

					// fmt.Printf("found at: x: %d, y: %d w xOffset: %d, yOffset: %d\n", currentX+1, currentY+1, xOffset, yOffset)
				}
			}

			// if word == "MAS" {
			// 	result++
			// }
		}
	}

	return result
}

func TrackerString(a, b int) string {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	return aStr + " " + bStr
}

func AddValue(x, y, xOffset, yOffset int, currChar string) {
	// fmt.Printf("xOffset: %d, yOffset: %d\n", xOffset, yOffset)
	// fmt.Printf("current char: %s\n", currChar)
	// fmt.Printf("x:%d y:%d \n", x, y)
	word := []rune("XMAS")
	xOffset *= -1
	yOffset *= -1

	for i := 0; i < 4; i++ {
		currentX := x + (xOffset * i)
		currentY := y + (yOffset * i)
		// fmt.Printf("currentx: %d, currenty: %d\n", currentX+1, currentY+1)
		newY := replaceAtIndex(saveboard[currentY], currentX, word[4-i-1])
		// fmt.Printf("newy: %s\n", newY)
		saveboard[currentY] = newY
	}
}

func replaceAtIndex(s string, index int, replacement rune) string {
	if index < 0 || index >= len(s) {
		return s // Return the original string if the index is out of range
	}
	// Convert the replacement rune to a string and concatenate slices
	return s[:index] + string(replacement) + s[index+1:]
}
