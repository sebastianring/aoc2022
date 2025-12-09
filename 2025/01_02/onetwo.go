package onetwo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

type Rotation struct {
	Direction string
	Clicks    int
}

func OneTwo(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	rotations := []Rotation{}
	for {
		lineString, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}
		lineString = strings.TrimSuffix(lineString, "\n")

		clicks, err := strconv.Atoi(string(lineString[1:]))
		if err != nil {
			log.Fatalf("issue converting string %s", err.Error())
		}

		rotations = append(rotations, Rotation{
			Direction: string(lineString[0]),
			Clicks:    clicks,
		})
	}

	zeroHits := 0
	pos := 50

	for _, rot := range rotations {
		fmt.Println("start pos", pos)
		fmt.Println("start zero", zeroHits)
		newPos := pos
		q, r := utils.DivMod(rot.Clicks, 100)
		fmt.Println(rot.Direction, rot.Clicks, q, r)
		zeroHits += q

		if rot.Direction == "R" {
			newPos += r
			if newPos > 99 {
				zeroHits++
				newPos = newPos - 100
			}
		} else {
			newPos -= r
			fmt.Println("negative new pos", newPos)
			if newPos < 0 {
				if pos != 0 {
					zeroHits++
				}
				newPos = 100 + newPos
			} else if newPos == 0 {
				zeroHits++
			}
		}

		pos = newPos
		fmt.Println("end pos", pos)
		fmt.Println("end zero", zeroHits)
	}

	return zeroHits, nil
}
