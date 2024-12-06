package daysixparttwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "time"
)

type Guard struct {
	x      int
	y      int
	xSpeed int
	ySpeed int
	steps  int
}

func DayOne(filename string) (int, error) {
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

	guard := Guard{}

	for y := 0; y < len(board)-1; y++ {
		for x := 0; x < len(board)-1; x++ {
			if board[y][x] == "^" {
				guard.x = x
				guard.y = y
				guard.xSpeed = 0
				guard.ySpeed = -1
				guard.steps = 1
			}
		}
	}

	sum = guard.Act(board)

	return sum, nil
}

func (g *Guard) Act(board [][]string) int {
	for {
		x, y := g.NextPos()

		if outOfBounds(board, x, y) {
			return g.steps
		}

		if checkObject(board, x, y) {
			g.Rotate()
		}

		board[g.y][g.x] = "X"

		g.MoveForward()
		added := g.AddStep(board)
		if !added {
			board[g.y][g.x] = "S"
			// prevX, prevY := g.PrevPos()
			// board[prevY][prevX] = "S"
		} else {
			board[g.y][g.x] = "O"
		}

		// printBoard(board)

		// time.Sleep(20 * time.Millisecond)
	}
}

func (g *Guard) MoveForward() {
	g.x += g.xSpeed
	g.y += g.ySpeed
}

func (g *Guard) PrevPos() (int, int) {
	return g.x - g.xSpeed, g.y - g.ySpeed
}

func (g *Guard) AddStep(board [][]string) bool {
	if board[g.y][g.x] != "X" || board[g.y][g.x] == "S" {
		g.steps++
		return true
	}
	return false
}

func (g *Guard) NextPos() (int, int) {
	return g.x + g.xSpeed, g.y + g.ySpeed
}

func (g *Guard) Rotate() {
	if g.xSpeed == 0 && g.ySpeed == -1 {
		g.xSpeed = 1
		g.ySpeed = 0
	} else if g.xSpeed == 1 && g.ySpeed == 0 {
		g.xSpeed = 0
		g.ySpeed = 1
	} else if g.xSpeed == 0 && g.ySpeed == 1 {
		g.xSpeed = -1
		g.ySpeed = 0
	} else if g.xSpeed == -1 && g.ySpeed == 0 {
		g.xSpeed = 0
		g.ySpeed = -1
	}
}

func checkObject(board [][]string, x, y int) bool {
	return board[y][x] == "#"
}

func outOfBounds(board [][]string, x, y int) bool {
	if x < 0 || y < 0 || y > len(board)-1 || x > len(board[y])-1 {
		return true
	}

	return false
}

func printBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(row)
	}
}
