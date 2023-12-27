package dayten2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

func DayTen(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	pipeMap := []string{}
	startPos := Pos{
		x: -1,
		y: -1,
	}
	y := 0

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		pipeMap = append(pipeMap, lineString)

		if startPos.x == -1 {
			for x, s := range lineString {
				if string(s) == "S" {
					startPos.x = x
					startPos.y = y
				}
			}
		}

		y++
	}

	for _, row := range pipeMap {
		fmt.Printf("%v\n", row)
	}

	fmt.Printf("Pos x: %v y: %v\n", startPos.x, startPos.y)

	return sum, nil
}
