package thirteen

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

func OneOne(filename string) (int, error) {
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

		_, remainder := utils.DivMod(clicks, 100)

		rotations = append(rotations, Rotation{
			Direction: string(lineString[0]),
			Clicks:    remainder,
			// Clicks:    clicks,
		})
	}

	zeroHits := 0
	pos := 50

	for _, rot := range rotations {
		if rot.Direction == "R" {
			pos += rot.Clicks
		} else {
			pos -= rot.Clicks
		}

		fmt.Println(rot.Direction, rot.Clicks)
		fmt.Println("pos before fixing", pos)

		if pos > 99 {
			pos -= 100
		} else if pos < 0 {
			pos = 100 + pos
		}

		if pos == 0 {
			zeroHits++
		}

		fmt.Println("pos after fixing", pos)
		fmt.Println("---------")
	}

	return zeroHits, nil
}
