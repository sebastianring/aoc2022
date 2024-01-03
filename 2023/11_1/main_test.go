package dayeleven2023_test

import (
	"fmt"
	dayelevent2023 "github.com/sebastianring/aoc2022/2023/11_1"
	"testing"
)

func TestDayElevenDataOne(t *testing.T) {
	result, err := dayelevent2023.DayEleven("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
