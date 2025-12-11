package threeone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

func ThreeOne(filename string) (int, error) {
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

	for _, batteryCell := range batteries {
		fmt.Println(batteryCell)
		highest := 0
		highestIndex := 0

		// find the highest
		for i := range len(batteryCell) - 1 {
			if batteryCell[i] > highest {
				highest = batteryCell[i]
				highestIndex = i
			}
		}

		fmt.Println("highest: ", highest)

		secondHighest := 0

		for i := highestIndex + 1; i < len(batteryCell); i++ {
			if batteryCell[i] > secondHighest {
				secondHighest = batteryCell[i]
			}
		}

		fmt.Println("second highest: ", highest)

		s := utils.FromIntsToStr([]int{highest, secondHighest})
		pSum, _ := strconv.Atoi(s)

		fmt.Println("psum: ", pSum)

		sum += pSum
	}

	return sum, nil
}
