package dayX2023_test

import (
	"fmt"
	dayX2023 "github.com/sebastianring/aoc2022/2023/3_1"
	"testing"
)

func TestDayX(t *testing.T) {
	result, err := dayX2023.DayX("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
