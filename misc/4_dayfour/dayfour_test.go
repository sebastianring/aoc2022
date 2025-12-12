package dayfour_test

import (
	"fmt"
	"testing"

	dayfour "github.com/sebastianring/aoc2022/misc/4_dayfour"
)

func TestOverlapSections(t *testing.T) {
	result, err := dayfour.OverlapSections("data.txt")
	if err != nil {
		t.Fatal("Error: ", err)
	}

	fmt.Println("Result from the test: ", result)
	t.Log(result)
}
