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

type area struct {
	poses     map[pos]int
	perims    map[pos]int
	perimsCtr int
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

	charMap := make(map[string]map[pos]int, 0)

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board); x++ {
			if board[y][x] == "." {
				continue
			}

			if charMap[board[y][x]] == nil {
				charMap[board[y][x]] = map[pos]int{}
			}

			charMap[board[y][x]][pos{x: x, y: y}] = 0
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

func findPerimeter(posMap map[pos]int) int {
	sum := 0
	var curPos pos
	var areas []area

	for len(posMap) > 0 {
		for k := range posMap {
			curPos = k
			break
		}

		curArea := area{
			poses: map[pos]int{
				curPos: 1,
			},
			perims: map[pos]int{},
		}

		curArea = recursiveAreaCheck(curPos, curArea, posMap)
		for p := range curArea.poses {
			_, exist := posMap[p]
			if exist {
				delete(posMap, p)
				// fmt.Printf("deleted from posMap: %v\n", posMap)
			} else {
				fmt.Printf("something fail\n")
			}
		}

		areas = append(areas, curArea)
	}

	// fmt.Printf("perimeter: %d, area: %d\n", sum, len(posMap))
	for _, a := range areas {
		fmt.Printf("area poses len: %d, perims len: %d\n", len(a.poses), len(a.perims))
		sum += a.perimsCtr * len(a.poses)
		fmt.Printf("area poses: %v, perims: %v\n", a.poses, a.perims)
	}

	return sum
}

func recursiveAreaCheck(curPos pos, a area, posMap map[pos]int) area {
	// fmt.Printf("starting to recurse check for: %v\n", curPos)
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
			x: curPos.x + xOffset,
			y: curPos.y + yOffset,
		}

		_, exists := posMap[checkPos]
		if !exists {
			// fmt.Printf("added a perimter\n")
			a.perims[checkPos] += 1
			a.perimsCtr += 1
		} else {
			_, checked := a.poses[checkPos]
			if !checked {
				a.poses[checkPos] += 1
				a = recursiveAreaCheck(checkPos, a, posMap)
			}
		}
	}

	// fmt.Printf("finished with recursive check")

	return a
}

// 	curArea.poses[curPos] = -1
// 	for i := 0; i < 4; i++ {
// 		var xOffset int
// 		var yOffset int
//
// 		switch i {
// 		case 0:
// 			xOffset = 1
// 			yOffset = 0
// 		case 1:
// 			xOffset = -1
// 			yOffset = 0
// 		case 2:
// 			xOffset = 0
// 			yOffset = 1
// 		case 3:
// 			xOffset = 0
// 			yOffset = -1
// 		}
//
// 		checkPos := pos{
// 			x: curPos.x + xOffset,
// 			y: curPos.y + yOffset,
// 		}
//
// 		_, exists := posMap[checkPos]
// 		if !exists {
// 			curArea.perims[checkPos] += 1
// 		} else {
// 			curArea.poses[checkPos] += 1
// 		}
// 	}
//
