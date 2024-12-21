package twelveone

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x int
	y int
}

func TwelveOne(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	board := [][]string{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		currentLine := []string{}
		for _, char := range lineString {
			currentLine = append(currentLine, string(char))
		}

		board = append(board, currentLine)
	}

	charMap := make(map[string]map[pos]bool, 0)

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board); x++ {
			if board[y][x] == "." {
				continue
			}

			if charMap[board[y][x]] == nil {
				charMap[board[y][x]] = map[pos]bool{}
			}

			charMap[board[y][x]][pos{x: x, y: y}] = true
		}
	}

	fmt.Println(charMap)

	for char, pos := range charMap {
		result := findPerimeter(pos)
		fmt.Printf("Checking for char: %s, result: %d \n", char, result)

		sum += result
	}

	return sum, nil
}

func findPerimeter(posMap map[pos]bool) int {
	sum := 0

	for p := range posMap {
		for i := 0; i < 4; i++ {
			var xOffset int
			var yOffset int

			switch i {
			case 0:
				xOffset = 1
				yOffset = 0
			case 1:
				xOffset = -1
				yOffset = 0
			case 2:
				xOffset = 0
				yOffset = 1
			case 3:
				xOffset = 0
				yOffset = -1
			}

			checkPos := pos{
				x: p.x + xOffset,
				y: p.y + yOffset,
			}

			_, exists := posMap[checkPos]
			if !exists {
				sum++
			}
		}
	}

	fmt.Printf("perimeter: %d, area: %d\n", sum, len(posMap))

	return sum * len(posMap)
}
