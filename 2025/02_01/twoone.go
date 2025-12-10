package twoone

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

type Interval struct {
	Start int
	End   int
}

func TwoOne(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	intervals := []Interval{}

	for {
		lineString, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		lineString = strings.TrimSuffix(lineString, ",")

		lineIDs := strings.Split(lineString, ",")

		for _, lineID := range lineIDs {
			intervalStrings := strings.Split(lineID, "-")
			start, _ := strconv.Atoi(intervalStrings[0])
			end, _ := strconv.Atoi(intervalStrings[1])

			intervals = append(intervals, Interval{
				Start: start,
				End:   end,
			})
		}
	}

	for _, interval := range intervals {
		// odd intervals no need to check
		// fmt.Println("start: ", interval.Start, " end: ", interval.End)
		// fmt.Println("diff", interval.End-interval.Start)
		ctr := interval.Start

		for ctr <= interval.End {
			ctrStr := strconv.Itoa(ctr)
			strLen := len(ctrStr)
			if !utils.IsEven(strLen) {
				ctr = newCtr(strLen)
				continue
			}

			first := ctrStr[:strLen/2]
			second := ctrStr[strLen/2:]

			if first == second {
				sum += ctr
			}

			firstInt, _ := strconv.Atoi(first)
			secondInt, _ := strconv.Atoi(second)

			if firstInt > secondInt {
				ctr, _ = strconv.Atoi(first + first)
			} else {
				firstInt += 1
				newFirst := strconv.Itoa(firstInt)
				ctr, _ = strconv.Atoi(newFirst + newFirst)
			}

			// fmt.Println("end ctr: ", ctr)
		}
	}

	return sum, nil
}

func newCtr(n int) int {
	start := "1"
	for range n {
		start += "0"
	}

	res, _ := strconv.Atoi(start)

	return res
}
