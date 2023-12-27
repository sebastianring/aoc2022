package daynine2023_test

import (
	"fmt"
	daynine2023 "github.com/sebastianring/aoc2022/2023/09_1"
	"testing"
)

func TestDayNineDataOne(t *testing.T) {
	result, err := daynine2023.DayNine("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestDayNineDataTwo(t *testing.T) {
	result, err := daynine2023.DayNine("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
