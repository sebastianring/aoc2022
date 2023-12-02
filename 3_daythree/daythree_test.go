package daythree_test

import (
	"fmt"
	"github.com/sebastianring/aoc2022/3_daythree"
	"testing"
)

func TestRucksack(t *testing.T) {
	result, err := daythree.SumPrioritiesInRucksack("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestRuckSackDataTwo(t *testing.T) {
	result, err := daythree.SumPrioritiesInRucksack("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestRucksackDayTwo(t *testing.T) {
	result, err := daythree.SumRucksackPartTwo("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
