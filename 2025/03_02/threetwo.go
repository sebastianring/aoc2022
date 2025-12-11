package threetwo

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

func ThreeTwo(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	batteries := [][]int{}

	for {
		lineString, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		batteryCell := make([]int, len(lineString))

		for i := range len(lineString) {
			n, _ := strconv.Atoi(string(lineString[i]))
			batteryCell[i] = n
		}

		batteries = append(batteries, batteryCell)
	}

	maxBatteries := 12
	for _, batteryCell := range batteries {
		// fmt.Println(batteryCell)
		res := []int{}
		i := 0

		for len(res) < maxBatteries {
			limit := len(batteryCell) - maxBatteries + len(res) + 1
			latestHighest := 0
			latestHighestIndex := 0

			// fmt.Println("looking at:", batteryCell[i:limit])

			for i < limit {
				if batteryCell[i] > latestHighest {
					latestHighest = batteryCell[i]
					latestHighestIndex = i
				}

				i++
			}

			res = append(res, latestHighest)

			i = latestHighestIndex + 1
			// limit = len(batteryCell) - maxBatteries + len(res) + 1

			// fmt.Println("i: ", i, "limit:", limit, "latest highest", latestHighest)
		}

		// fmt.Println(res)

		sum += utils.MergeInts(res)
	}

	return sum, nil
}
