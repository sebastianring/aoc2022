package daysix2023_test

import (
	"fmt"
	"testing"

	daysix2023 "github.com/sebastianring/aoc2022/2023/6_1"
)

func TestButtonToWinDataOne(t *testing.T) {
	result, err := daysix2023.ButtonToWin("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}

func TestButtonToWinDataTwo(t *testing.T) {
	result, err := daysix2023.ButtonToWin("data2.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
