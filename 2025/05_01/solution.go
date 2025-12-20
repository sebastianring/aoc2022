package template

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sebastianring/aoc2022/utils"
)

func Template(lines []string) int {
	sum := 0
	ranges, ids := FormatData(lines)

	for _, r := range ranges {
		fmt.Println(r)
	}

	for _, id := range ids {
		fmt.Println(id)
	}

	ranges = MergeOverlaps(ranges)
	for _, r := range ranges {
		fmt.Println(r)
	}

	for _, id := range ids {
		for _, r := range ranges {
			if InRange(id, r) {
				sum++
			}
		}
	}

	return sum
}

func MergeOverlaps(ranges []Range) []Range {
	merged := map[int]bool{}
	newRanges := []Range{}

	for i, r := range ranges {
		if _, ok := merged[i]; ok {
			continue
		}

		currentRange := r
		for j := i + 1; j < len(ranges); j++ {
			if ok, newRange := RangeInRange(currentRange, ranges[j]); ok {
				currentRange = newRange
				fmt.Println("found overlaps in original: ", currentRange, ranges[j])
				merged[j] = true
				merged[i] = true
			}
		}

		hits := []int{}
		if currentRange != r {
			for j := range newRanges {
				if ok, newRange := RangeInRange(currentRange, newRanges[j]); ok {
					currentRange = newRange
					fmt.Println("found overlaps in merged: ", currentRange, newRanges[j])
					hits = append(hits, j)
					merged[i] = true
				}
			}
		}

		for _, idx := range hits {
			newRanges = utils.RemoveAtIndex(newRanges, idx)
		}

		fmt.Println("added current range: ", currentRange)
		newRanges = append(newRanges, currentRange)
	}

	return newRanges
}

func InRange(x int, r Range) bool {
	if x >= r.Min && x <= r.Max {
		return true
	}

	return false
}

func RangeInRange(r1 Range, r2 Range) (bool, Range) {
	res := false
	if r1.Max >= r2.Min && r1.Min <= r2.Max {
		res = true
	} else if r1.Max >= r2.Max && r1.Min <= r2.Max {
		res = true
	} else if r1.Min <= r2.Min && r1.Max >= r2.Max {
		res = true
	} else if r2.Min <= r1.Min && r2.Max >= r1.Max {
		res = true
	}

	return res, Range{
		Min: min(r1.Min, r2.Min),
		Max: max(r1.Max, r2.Max),
	}
}

func FormatData(lines []string) ([]Range, []int) {
	ranges := []Range{}
	ids := []int{}
	idx := 0

	for i, line := range lines {
		if len(line) == 0 {
			idx = i
			break
		}

		parts := strings.Split(line, "-")

		minI, _ := strconv.Atoi(parts[0])
		maxI, _ := strconv.Atoi(parts[1])
		r := Range{
			Min: minI,
			Max: maxI,
		}

		ranges = append(ranges, r)
	}

	for _, line := range lines[idx+1:] {
		i, _ := strconv.Atoi(line)

		ids = append(ids, i)
	}

	return ranges, ids
}

type Range struct {
	Min int
	Max int
}
