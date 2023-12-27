package dayfour2023_test

import (
	"fmt"
	"testing"

	dayfour2023 "github.com/sebastianring/aoc2022/2023/04_1"
)

func TestScratchCards(t *testing.T) {
	result, err := dayfour2023.ScratchCards("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestScratchCardsPartTwoDataOne(t *testing.T) {
	result, err := dayfour2023.ScratchCardsPartTwo("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestScratchCardsPartTwoDataTwo(t *testing.T) {
	result, err := dayfour2023.ScratchCardsPartTwo("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
