package fourtwo

import (
	"fmt"
)

func FourOneSolution(lines []string) int {
	sum := 0
	h := len(lines)
	w := len(lines[0])

	scoreBoard := FormatData(lines)

	startingPoints := [][]int{}
	for y, line := range lines {
		for x, val := range line {
			if val == rune('@') {
				AdjustAdjescentScore(x, y, h, w, scoreBoard, 1)
				startingPoints = append(startingPoints, []int{x, y})
			}
		}
	}

	return Cascade(startingPoints, h, w, scoreBoard, lines, sum)
}

func Cascade(adjusted [][]int, h, w int, scoreBoard [][]int, lines []string, sum int) int {
	if len(adjusted) == 0 {
		return sum
	}

	sum = 0

	for _, adj := range adjusted {
		x := adj[0]
		y := adj[1]

		if scoreBoard[y][x] < 4 && lines[y][x] == '@' {
			lines[y] = lines[y][:x] + "r" + lines[y][x+1:]
			sum++
			sum += Cascade(AdjustAdjescentScore(x, y, h, w, scoreBoard, -1), h, w, scoreBoard, lines, sum)
		}
	}

	return sum
}

func AdjustAdjescentScore(x, y, h, w int, scoreBoard [][]int, val int) [][]int {
	adjusted := [][]int{}

	for xOffset := -1; xOffset <= 1; xOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			if yOffset == 0 && xOffset == 0 {
				continue
			}

			if ValidPos(x+xOffset, y+yOffset, h, w) {
				scoreBoard[y+yOffset][x+xOffset] += val
				adjusted = append(adjusted, []int{x + xOffset, y + yOffset})
			}
		}
	}

	return adjusted
}

func PrintBoard(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println("------------------")
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
