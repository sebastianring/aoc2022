package twelvetwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

type pos struct {
	x       int
	y       int
	up      *pos
	down    *pos
	left    *pos
	right   *pos
	started bool
}

func TwelveTwo(filename string) (int, error) {
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

	posBoard := map[string][][]*pos{}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if board[y][x] == "." {
				continue
			}

			if posBoard[board[y][x]] == nil {
				newBoard := make([][]*pos, len(board))
				for i := 0; i < len(board); i++ {
					newRow := make([]*pos, len(board[i]))
					newBoard[i] = newRow
				}
				posBoard[board[y][x]] = newBoard
			}

			posBoard[board[y][x]][y][x] = &pos{x: x, y: y}
		}
	}

	for char, m := range posBoard {
		printPosBoard(m, char)
		areas := linkPosToAreas(m)
		for i, area := range areas {
			corners := calcCorners(area, len(m), len(m[0]))
			fmt.Printf("area #%d pos: %v corners: %d\n", i, len(area), corners)
			for ctr, p := range area {
				fmt.Printf("P#%d [%d,%d] ", ctr, p.x, p.y)
			}
			fmt.Println()
			sum += corners * len(area)
		}
	}

	return sum, nil
}

func mapPs(ps []*pos) map[pos]int {
	result := map[pos]int{}
	for _, p := range ps {
		tP := pos{
			x: p.x,
			y: p.y,
		}

		result[tP] += 1
	}

	return result
}

func calcCorners(ps []*pos, rows, cols int) int {
	res := 0
	// oldincorn := map[pos]int{}
	mappedPs := mapPs(ps)
	innerCorners := make([][]*pos, rows)

	for y := 0; y < rows; y++ {
		row := make([]*pos, cols)
		innerCorners[y] = row
	}

	for _, p := range ps {
		if p.down == nil && p.left == nil {
			tpos := pos{
				x: p.x - 1,
				y: p.y + 1,
			}
			_, ok := mappedPs[tpos]
			if !ok {
				res += 1
			}

		}

		if p.down == nil {
			if p.right == nil {
				tpos := pos{
					x: p.x + 1,
					y: p.y + 1,
				}
				_, ok := mappedPs[tpos]
				if !ok {
					res += 1
				}
			}
		}

		if p.up == nil && p.right == nil {
			tpos := pos{
				x: p.x + 1,
				y: p.y - 1,
			}
			_, ok := mappedPs[tpos]
			if !ok {
				res += 1
			}
		}

		if p.up == nil && p.left == nil {
			tpos := pos{
				x: p.x - 1,
				y: p.y - 1,
			}

			_, ok := mappedPs[tpos]
			if !ok {
				res += 1
			}
		}

		// inner corners
		if p.down == nil {
			x := p.x
			y := p.y + 1

			if !utils.OutOfBounds(innerCorners, x, y) {
				if innerCorners[y][x] == nil {
					innerCorners[y][x] = &pos{
						x: x,
						y: y,
					}
				}

				innerCorners[y][x].up = p
			}
		}

		if p.right == nil {
			x := p.x + 1
			y := p.y

			if !utils.OutOfBounds(innerCorners, x, y) {
				if innerCorners[y][x] == nil {
					innerCorners[y][x] = &pos{
						x: x,
						y: y,
					}
				}

				innerCorners[y][x].left = p
			}

		}

		if p.left == nil {
			x := p.x - 1
			y := p.y

			fmt.Println(p.x, p.y, "checking left hand side", x, y, p.left)

			if !utils.OutOfBounds(innerCorners, x, y) {
				if innerCorners[y][x] == nil {
					innerCorners[y][x] = &pos{
						x: x,
						y: y,
					}
				}

				innerCorners[y][x].right = p
			}
		}

		if p.up == nil {
			x := p.x
			y := p.y - 1

			if !utils.OutOfBounds(innerCorners, x, y) {
				if innerCorners[y][x] == nil {
					innerCorners[y][x] = &pos{
						x: x,
						y: y,
					}
				}

				innerCorners[y][x].down = p
			}
		}
	}

	for y := 0; y < len(innerCorners); y++ {
		for x := 0; x < len(innerCorners[y]); x++ {
			if innerCorners[y][x] != nil {
				fmt.Printf("inner corner: %d, %d - loop: %d %d\n", innerCorners[y][x].x, innerCorners[y][x].y, x, y)
				fmt.Printf("vals: %d %d up: %v, left: %v, right: %v, down: %v\n", x, y, innerCorners[y][x].up, innerCorners[y][x].left, innerCorners[y][x].right, innerCorners[y][x].down)
				if innerCorners[y][x].up != nil && innerCorners[y][x].left != nil {
					fmt.Printf("up left\n")
					res += 1
				}

				if innerCorners[y][x].up != nil && innerCorners[y][x].right != nil {
					fmt.Printf("up right\n")
					res += 1
				}

				if innerCorners[y][x].down != nil && innerCorners[y][x].right != nil {
					fmt.Printf("down right\n")
					res += 1
				}

				if innerCorners[y][x].down != nil && innerCorners[y][x].left != nil {
					fmt.Printf("down left\n")
					res += 1
				}
			}
		}
	}

	return res
}

func linkPosToAreas(b [][]*pos) [][]*pos {
	areas := [][]*pos{}

	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[y]); x++ {
			var area []*pos
			if b[y][x] != nil {
				fmt.Println("found a new area")
				b, area = b[y][x].recursiveLink(b, nil)
				areas = append(areas, area)
			}
		}
		fmt.Println()
	}

	return areas
}

func (cur *pos) recursiveLink(b [][]*pos, area []*pos) ([][]*pos, []*pos) {
	// fmt.Printf("starting rec link at: %v\n", cur)
	if !cur.started {
		cur.started = true
		area = append(area, cur)
	} else {
		return b, area
	}

	for i := 0; i < 4; i++ {
		var xOffset int
		var yOffset int

		switch i {
		// right
		case 0:
			if cur.right != nil {
				continue
			}
			xOffset = 1
			yOffset = 0
		// left
		case 1:
			if cur.left != nil {
				continue
			}
			xOffset = -1
			yOffset = 0
		// down
		case 2:
			if cur.down != nil {
				continue
			}
			xOffset = 0
			yOffset = 1
		// up
		case 3:
			if cur.up != nil {
				continue
			}
			xOffset = 0
			yOffset = -1
		}

		x := cur.x + xOffset
		y := cur.y + yOffset

		if utils.OutOfBounds(b, x, y) {
			continue
		}

		if b[y][x] != nil {
			switch i {
			case 0:
				cur.right = b[y][x]
				b[y][x].left = cur
			case 1:
				cur.left = b[y][x]
				b[y][x].right = cur
			case 2:
				cur.down = b[y][x]
				b[y][x].up = cur
			case 3:
				cur.up = b[y][x]
				b[y][x].down = cur
			}

			if !b[y][x].started {
				b, area = b[y][x].recursiveLink(b, area)
			}
		}
	}

	b[cur.y][cur.x] = nil

	return b, area
}

func printPosBoard(b [][]*pos, char string) {
	fmt.Println("-------")
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b[y]); x++ {
			if b[y][x] == nil {
				fmt.Print(".")
			} else {
				fmt.Print(char)
			}
		}
		fmt.Println()
	}
}
