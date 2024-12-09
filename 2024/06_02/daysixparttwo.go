package daysixparttwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "time"
	// "time"
)

type rotation struct {
	xSpeed int
	ySpeed int
	x      int
	y      int
}

type Guard struct {
	x               int
	y               int
	xSpeed          int
	ySpeed          int
	steps           int
	potantialBlocks int
	permutation     bool
	rotatations     []rotation
	traveledSteps   [][]int
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
		for x := 0; x < len(board[y])-1; x++ {
			if board[y][x] == "^" {
				guard.x = x
				guard.y = y
				guard.xSpeed = 0
				guard.ySpeed = -1
				guard.steps = 1
				guard.rotatations = append(guard.rotatations, rotation{
					xSpeed: 0,
					ySpeed: -1,
					x:      x,
					y:      y,
				})
			}
		}
	}

	sum = guard.Act(board)

	return sum, nil
}

func (g *Guard) Act(board [][]string) int {
	for {
		if g.permutation {
			if g.potentialBlock() && g.steps > 0 {
				return 1
			}
		}

		var x, y int
		for {
			x, y = g.NextPos()
			if outOfBounds(board, x, y) {
				if !g.permutation {
					return g.potantialBlocks
				}

				return 0
			} else if checkObject(board, x, y) {
				g.Rotate()
			} else {
				break
			}
		}

		if !g.permutation && board[y][x] != "Z" {
			permutation := Guard{
				x:               g.x,
				y:               g.y,
				xSpeed:          g.xSpeed,
				ySpeed:          g.ySpeed,
				steps:           0,
				potantialBlocks: 0,
				permutation:     true,
			}

			permutation.Rotate()
			// fmt.Printf("---------------\n")
			// fmt.Printf("rotations before start: %v\n", permutation.rotatations)
			boardCopy := make([][]string, len(board))
			for ctr := range board {
				rowCopy := make([]string, len(board[ctr]))
				copy(rowCopy, board[ctr])
				boardCopy[ctr] = rowCopy
			}

			boardCopy[y][x] = "#"

			// fmt.Printf("BOARD COPY: %v\n", boardCopy)

			result := permutation.Act(boardCopy)
			if result > 0 {
				// fmt.Printf("result from perm:%d\n", result)
				// fmt.Printf("rotations after end: %v\n", permutation.rotatations)
				// fmt.Printf("x: %d\n", g.x)
				// fmt.Printf("y: %d\n", g.y)
				// fmt.Printf("xSpeed: %d \n", g.xSpeed)
				// fmt.Printf("ySpeed: %d \n", g.ySpeed)
				// time.Sleep(1 * time.Second)
			}
			g.potantialBlocks += result
		}

		board[g.y][g.x] = "Z"
		g.MoveForward()

		// char := "X"
		// speed := 30 * time.Millisecond
		// if g.permutation {
		// 	speed = 25 * time.Millisecond
		// 	char = "O"
		// }

		// printBoard(board, g.x, g.y, char)
		// time.Sleep(speed)
	}
}

func (g *Guard) potentialBlock() bool {
	// if g.x == g.startingRotation.x && g.y == g.startingRotation.y && g.xSpeed == g.startingRotation.xSpeed && g.ySpeed == g.startingRotation.ySpeed {
	// 	return true
	// }

	for _, rot := range g.rotatations {
		if g.x == rot.x && g.y == rot.y && g.xSpeed == rot.xSpeed && g.ySpeed == rot.ySpeed {
			// if i != 0 {
			// 	fmt.Printf("rotations: %v\n", g.rotatations)
			// 	fmt.Printf("steps: %d\n", g.steps)
			// 	fmt.Printf("rotation #%d\n", i)
			// 	fmt.Printf("x: %d rot x: %d\n", g.x, rot.x)
			// 	fmt.Printf("y: %d rot y: %d\n", g.y, rot.y)
			// 	fmt.Printf("xSpeed: %d rot xSpeed: %d\n", g.xSpeed, rot.xSpeed)
			// 	fmt.Printf("ySpeed: %d rot ySpeed: %d\n", g.ySpeed, rot.ySpeed)
			// }
			return true
		}
	}

	return false
}

func (g *Guard) MoveForward() {
	g.traveledSteps = append(g.traveledSteps, []int{g.x, g.y})
	g.x += g.xSpeed
	g.y += g.ySpeed
	g.steps++
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
	g.rotatations = append(g.rotatations, rotation{
		xSpeed: g.xSpeed,
		ySpeed: g.ySpeed,
		x:      g.x,
		y:      g.y,
	})

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

func printBoard(board [][]string, x, y int, char string) {
	fmt.Println("--------")
	for boardY, row := range board {
		if y == boardY {
			temprow := make([]string, len(row))
			copy(temprow, row)
			temprow[x] = char
			fmt.Println(temprow)
		} else {
			fmt.Println(row)
		}
	}
}
