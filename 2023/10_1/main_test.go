package dayten2023_test

import (
	"fmt"
	dayten2023 "github.com/sebastianring/aoc2022/2023/10_1"
	"testing"
)

func TestDayTenDataOne(t *testing.T) {
	result, err := dayten2023.DayTen("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
