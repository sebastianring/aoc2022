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
	Direction     string
	Clicks        int
	FullRotations int
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

		quotient, remainder := utils.DivMod(clicks, 100)

		rotations = append(rotations, Rotation{
			Direction:     string(lineString[0]),
			Clicks:        remainder,
			FullRotations: quotient,
		})
	}

	zeroHits := 0
	pos := 50

	for _, rot := range rotations {
		startP := pos

		if rot.FullRotations > 0 {
			fmt.Println("starting pos", pos)
		}

		if rot.Direction == "R" {
			pos += rot.Clicks
		} else {
			pos -= rot.Clicks
		}

		if rot.FullRotations > 0 {
			fmt.Println(rot.Direction, rot.Clicks, rot.FullRotations, zeroHits)
			fmt.Println("pos before fixing", pos)
		}

		zeroHits += rot.FullRotations

		if pos > 99 {
			pos -= 100
			zeroHits++
		} else if pos < 0 {
			pos = 100 + pos
			if startP != 0 {
				zeroHits++
			}
		}

		// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
		//        |
		//  pos = 3
		//  L5
		//
		// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
		//                       |
		//
		//   2 - 5 = -2

		if rot.FullRotations > 0 {
			fmt.Println("pos after fixing", pos)
			fmt.Println("zero hits after", zeroHits)
			fmt.Println("---------")
		}
	}

	return zeroHits, nil
}
