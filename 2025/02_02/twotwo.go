package twotwo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

type Interval struct {
	Start int
	End   int
}

func TwoTwo(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sums := map[int]bool{}
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
		fmt.Println("start: ", interval.Start, " end: ", interval.End)
		// fmt.Println("diff", interval.End-interval.Start)
		ctr := interval.Start

		for ctr <= interval.End {
			ctrStr := strconv.Itoa(ctr)
			strLen := len(ctrStr)

			for i := 2; i <= strLen; i++ {
				q, r := utils.DivMod(strLen, i)
				if r != 0 {
					continue
				}

				parts := utils.SplitBySize(ctrStr, q)
				partsInt := toInts(parts)

				for j := 1; j < len(partsInt); j++ {
					if partsInt[0] < partsInt[j] {
						break
					}

					partsInt[j] = partsInt[0]

					if j == len(partsInt)-1 {
						res := mergeInts(partsInt)
						if res <= interval.End {
							fmt.Println("added: ", res)
							sums[mergeInts(partsInt)] = true
						}
					}
				}
			}

			ctr++
			fmt.Println("new ctr: ", ctr)
			fmt.Println("------")
		}
	}

	fmt.Println(sums)
	for k := range sums {
		sum += k
	}

	return sum, nil
}

func mergeInts(inputs []int) int {
	string := toStr(inputs)
	res, _ := strconv.Atoi(string)

	return res
}

func toStr(inputs []int) string {
	res := ""
	for _, input := range inputs {
		res += strconv.Itoa(input)
	}

	return res
}

func toInts(inputs []string) []int {
	res := []int{}
	for _, inp := range inputs {
		newInt, _ := strconv.Atoi(inp)
		res = append(res, newInt)
	}

	return res
}
