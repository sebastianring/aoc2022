package fourone

import "fmt"

func FourOneSolution(lines []string) int {
	sum := 0
	h := len(lines)
	w := len(lines[0])

	scoreBoard := FormatData(lines)

	for _, col := range scoreBoard {
		for _, val := range col {
			fmt.Printf("%d", val)
		}
		fmt.Printf("\n")
	}

	for y, line := range lines {
		for x, val := range line {
			if val == rune('@') {
				AddAdjacentScores(x, y, h, w, scoreBoard)
			}
		}
	}

	for y, line := range lines {
		for x, val := range line {
			if val == rune('@') {
				if scoreBoard[y][x] < 4 {
					sum++
				}
			}
		}
	}

	return sum
}

func AddAdjacentScores(x, y, h, w int, scoreBoard [][]int) {
	for xOffset := -1; xOffset <= 1; xOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			if yOffset == 0 && xOffset == 0 {
				continue
			}

			if ValidPos(x+xOffset, y+yOffset, h, w) {
				scoreBoard[y+yOffset][x+xOffset]++
			}
		}
	}
}

func FormatData(lines []string) [][]int {
	scoreBoard := make([][]int, len(lines))

	for row, data := range lines {
		scoreBoard[row] = make([]int, len(data))
	}

	return scoreBoard
}

func ValidPos(x, y, h, w int) bool {
	if x < 0 || y < 0 || y > h-1 || x > w-1 {
		return false
	}

	return true
}
