package daytwopartone

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2Part1(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0

	for {
		var lineInts []int
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		ints := strings.Split(lineString, " ")

		for _, no := range ints {
			noInt, _ := strconv.Atoi(no)
			lineInts = append(lineInts, noInt)
		}

		safe := isSafe(lineInts)
		if safe {
			fmt.Print("GOT A SCORE?\n")
			sum++
		}

		fmt.Printf("-------------------\n")
	}

	return sum, nil
}

func isSafe(lineInts []int) bool {
	// rule #1 - all need to increase or decrease
	// rule #2 minmum 1 and maximum of 3 difference between two values

	var maxDiff int
	var minDiff int
	var lastVal int

	for i, no := range lineInts {
		if i == 0 {
			if no < lineInts[i+1] {
				maxDiff = 3
				minDiff = 1
			} else {
				maxDiff = -1
				minDiff = -3
			}

			lastVal = no
			continue
		}

		safe := isWithinLimit(lastVal, no, minDiff, maxDiff)
		if !safe {
			return false
		}

		lastVal = no
	}

	return true
}

func isWithinLimit(a int, b int, minDiff int, maxDiff int) bool {
	aMin := a + minDiff
	aMax := a + maxDiff

	fmt.Printf("a: %d b: %d aMin: %d, aMax: %d \n", a, b, aMin, aMax)

	if b >= aMin && b <= aMax {
		fmt.Printf("return TRUE\n")
		return true
	}

	fmt.Printf("return FALSE\n")
	return false
}
