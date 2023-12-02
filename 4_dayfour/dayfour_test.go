package dayfour_test

import (
	"fmt"
	"github.com/sebastianring/aoc2022/4_dayfour"
	"testing"
)

func TestOverlapSections(t *testing.T) {
	result, err := dayfour.OverlapSections("data.txt")

	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
