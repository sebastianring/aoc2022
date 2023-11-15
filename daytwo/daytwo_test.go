package daytwo_test

import (
	"fmt"
	"github.com/sebastianring/aoc2022/daytwo"
	"testing"
)

func TestRockPaperScissor(t *testing.T) {
	result, err := daytwo.RockPaperScissor("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
