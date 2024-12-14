package tenone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func TenOne(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	topoMap := [][]string{}
	startingPos := []pos{}
	board := [][]int{}
	nines := map[pos]int{}

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

		topoMap = append(topoMap, currentLine)
	}

	for y := 0; y < len(topoMap); y++ {
		newRow := []int{}
		for x := 0; x < len(topoMap[y]); x++ {
			if topoMap[y][x] == "0" {
				newPos := pos{
					x: x,
					y: y,
				}

				startingPos = append(startingPos, newPos)
			}

			if topoMap[y][x] == "9" {
				nine := pos{
					x: x,
					y: y,
				}

				nines[nine] += 1
			}

			val, _ := strconv.Atoi(topoMap[y][x])
			newRow = append(newRow, val)
		}

		board = append(board, newRow)
	}

	for _, row := range board {
		fmt.Println(row)
	}

	fmt.Printf("Starting pos: %v\n", startingPos)

	sum = start(startingPos, nines, board)

	return sum, nil
}

func start(startingPos []pos, nines map[pos]int, board [][]int) int {
	sum := 0
	for _, s := range startingPos {
		for k := range nines {
			nines[k] = 1
		}

		result := recursiveCheck(board, []pos{s}, nines)
		sum += result
	}

	return sum
}

func recursiveCheck(board [][]int, traveled []pos, nines map[pos]int) int {
	result := 0
	curPos := traveled[len(traveled)-1]
	if board[curPos.y][curPos.x] == 9 {
		hits, exists := nines[curPos]
		if !exists {
			fmt.Printf("cant find the mapped nine... somewthing very wrong")
		}

		if hits > 0 {
			nines[curPos] -= 1
			return 1
		}

		return 0
	}

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

		// fmt.Printf("xoffset: %d - yoffset: %d\n", xOffset, yOffset)
		if outOfBounds(board, curPos.x+xOffset, curPos.y+yOffset) {
			continue
		}

		// printBoard(board, curPos.x+xOffset, curPos.y+yOffset, 0)
		// time.Sleep(200 * time.Millisecond)

		// fmt.Printf("checking: x: %d y: %d - new value: %d, if check:%v\n", curPos.x+xOffset, curPos.y+yOffset, board[curPos.y+yOffset][curPos.x+xOffset], board[curPos.y+yOffset][curPos.x+xOffset] == board[curPos.y][curPos.x]+1)
		if board[curPos.y+yOffset][curPos.x+xOffset] == board[curPos.y][curPos.x]+1 {
			// fmt.Println("Perm -----")
			traveled = append(traveled, pos{
				x: curPos.x + xOffset,
				y: curPos.y + yOffset,
			})
			tempRes := recursiveCheck(board, traveled, nines)
			if tempRes > 0 {
				result += tempRes
			}

			traveled = traveled[:len(traveled)-1]
		}
	}

	return result
}

func outOfBounds(board [][]int, x, y int) bool {
	if x < 0 || y < 0 || y > len(board)-1 || x > len(board[y])-1 {
		return true
	}

	return false
}

func printBoard(board [][]int, x, y int, char int) {
	fmt.Println("--------")
	for boardY, row := range board {
		if y == boardY {
			temprow := make([]int, len(row))
			copy(temprow, row)
			temprow[x] = char
			fmt.Println(temprow)
		} else {
			fmt.Println(row)
		}
	}
}
