package eight

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

func DayEightPartOne(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	board := [][]string{}
	antennas := map[string][]pos{}

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

	cleanBoard := make([][]string, len(board))

	for y := 0; y < len(board); y++ {
		rowCopy := make([]string, len(board[y]))
		copy(rowCopy, board[y])

		for x := 0; x < len(board[y]); x++ {
			if board[y][x] != "." {
				rowCopy[x] = "."
				antennas[board[y][x]] = append(antennas[board[y][x]], pos{x: x, y: y})
			}
		}

		cleanBoard[y] = rowCopy
	}

	for a, pos := range antennas {
		// fmt.Println(a, pos)
		// superCleanBoard := copyBoard(cleanBoard)

		for i := 0; i < len(pos)-1; i++ {
			for j := i + 1; j < len(pos); j++ {
				fmt.Printf("antenna: %s - checking pos: %v and %d\n", a, pos[i], pos[j])
				sum += createAntinodes(pos[i], pos[j], cleanBoard)
			}
		}
	}

	fmt.Println("-------")
	printBoard(board)
	fmt.Println("-------")

	return sum, nil
}

func createAntinodes(a, b pos, board [][]string) int {
	result := 0
	xDiff, yDiff := getDiff(a, b)
	fmt.Printf("xdiff: %d, ydiff: %d\n", xDiff, yDiff)

	antinode1 := pos{x: a.x + xDiff, y: a.y + yDiff}
	antinode2 := pos{x: b.x + xDiff*-1, y: b.y + yDiff*-1}

	if !outOfBounds(board, antinode1.x, antinode1.y) && board[antinode1.y][antinode1.x] != "#" {
		result++
		board[antinode1.y][antinode1.x] = "#"
	}

	if !outOfBounds(board, antinode2.x, antinode2.y) && board[antinode2.y][antinode2.x] != "#" {
		result++
		board[antinode2.y][antinode2.x] = "#"
	}

	printBoard(board)

	return result
}

func createAntinodesGrid(a, b pos, board [][]string) int {
	result := 0
	xDiff, yDiff := getDiff(a, b)
	fmt.Printf("xdiff: %d, ydiff: %d\n", xDiff, yDiff)

	i := 0
	for {
		xOffset := xDiff * i
		yOffset := yDiff * i

		x := a.x + xOffset
		y := a.y + yOffset

		fmt.Printf("x: %d, y: %d\n", x, y)

		if outOfBounds(board, x, y) {
			break
		}

		if board[y][x] == "#" {
			i++
			continue
		}

		result++
		board[y][x] = "#"
		i++
	}

	i = 0
	for {
		xOffset := (xDiff * -1) * i
		yOffset := (yDiff * -1) * i

		x := a.x + xOffset
		y := a.y + yOffset

		fmt.Printf("x: %d, y: %d\n", x, y)

		if outOfBounds(board, x, y) {
			break
		}

		if board[y][x] == "#" {
			i++
			continue
		}

		result++
		board[y][x] = "#"
		i++
	}

	//
	// antinode1 := pos{x: a.x + xDiff, y: a.y + yDiff}
	// antinode2 := pos{x: b.x + xDiff*-1, y: b.y + yDiff*-1}
	//
	// if !outOfBounds(board, antinode1.x, antinode1.y) && board[antinode1.y][antinode1.x] != "#" {
	// 	result++
	// 	board[antinode1.y][antinode1.x] = "#"
	// }
	//
	// if !outOfBounds(board, antinode2.x, antinode2.y) && board[antinode2.y][antinode2.x] != "#" {
	// 	result++
	// 	board[antinode2.y][antinode2.x] = "#"
	// }

	printBoard(board)

	return result
}

func getDiff(a, b pos) (int, int) {
	x := a.x - b.x
	y := a.y - b.y

	return x, y
}

func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func copyBoard(board [][]string) [][]string {
	cleanBoard := make([][]string, len(board))

	for y := 0; y < len(board); y++ {
		rowCopy := make([]string, len(board[y]))
		copy(rowCopy, board[y])

		cleanBoard[y] = rowCopy
	}

	return cleanBoard
}

func outOfBounds(board [][]string, x, y int) bool {
	if x < 0 || y < 0 || y > len(board)-1 || x > len(board[y])-1 {
		return true
	}

	return false
}

func DayEightPartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	board := [][]string{}
	antennas := map[string][]pos{}

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

	cleanBoard := make([][]string, len(board))

	for y := 0; y < len(board); y++ {
		rowCopy := make([]string, len(board[y]))
		copy(rowCopy, board[y])

		for x := 0; x < len(board[y]); x++ {
			if board[y][x] != "." {
				rowCopy[x] = "."
				antennas[board[y][x]] = append(antennas[board[y][x]], pos{x: x, y: y})
			}
		}

		cleanBoard[y] = rowCopy
	}

	for a, pos := range antennas {
		// fmt.Println(a, pos)
		// superCleanBoard := copyBoard(cleanBoard)

		for i := 0; i < len(pos)-1; i++ {
			for j := i + 1; j < len(pos); j++ {
				fmt.Printf("antenna: %s - checking pos: %v and %d\n", a, pos[i], pos[j])
				sum += createAntinodesGrid(pos[i], pos[j], cleanBoard)
			}
		}
	}

	fmt.Println("-------")
	printBoard(board)
	fmt.Println("-------")

	return sum, nil
}
