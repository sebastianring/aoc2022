package dayeight2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	Name  string
	Left  *Direction
	Right *Direction
}

type NewDirectionResponse struct {
	Direction *Direction
	Left      string
	Right     string
	Err       error
}

func HauntedWasteland(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	row := 0
	// instructions := ""

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		memoMap := map[string]NewDirectionResponse

		if row > 1 {
			ndr := NewDirectionFromInput(lineString)

			if ndr.Err != nil {
				return 0, err
			}

			memoMap[ndr.Direction.Name]

		} else if row == 0 {
			// instructions = lineString
			row++
		} else {
			row++
			continue
		}
	}

	return sum, nil
}

func NewDirectionFromInput(input []string) NewDirectionResponse {

	segments := strings.Split(input[], " ")
	name := segments[0]

	left := segments[2]
	strings.TrimLeft(left, "(")
	strings.TrimRight(left, ",")

	right := segments[3]
	strings.TrimRight(right, ")")

	d := NewDirection(name)

	ndr := NewDirectionResponse{
		Direction: d,
		Left:      left,
		Right:     right,
		Err:       nil,
	}

	return ndr
}

func NewDirection(name string) *Direction {
	d := Direction{
		Name: name,
	}

	return &d
}
