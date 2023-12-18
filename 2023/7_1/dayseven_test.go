package dayseven2023_test

import (
	"fmt"
	"testing"

	dayseven2023 "github.com/sebastianring/aoc2022/2023/7_1"
)

func TestCamelCardsDataOne(t *testing.T) {
	result, err := dayseven2023.CamelCards("data3.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
