package dayeight2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type DirectionHead struct {
	Current     *Direction
	OriginalDir *Direction
	Ctr         int
	Done        bool
	Multiplier  int
}

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

type SafeCounter struct {
	mutex   sync.Mutex
	counter int
}

func (sc *SafeCounter) Increment() {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	sc.counter++
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
	instructions := ""
	ndrMap := map[string]*NewDirectionResponse{}
	var next *Direction

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		if row > 1 {
			ndr := NewDirectionFromInput(lineString)

			if ndr.Err != nil {
				return 0, err
			}

			ndrMap[ndr.Direction.Name] = ndr
			fmt.Printf("Added ndr with name: %v with left: %v and right: %v \n",
				ndr.Direction.Name, ndr.Left, ndr.Right)

		} else if row == 0 {
			instructions = lineString
			row++

			fmt.Printf("Instructions: %v\n", instructions)
		} else {
			row++
			continue
		}
	}

	dirMap := map[string]*Direction{}

	for k, ndr := range ndrMap {
		dirMap[k] = ndr.Direction
		dirMap[k].Left = ndrMap[ndr.Left].Direction
		dirMap[k].Right = ndrMap[ndr.Right].Direction
	}

	next = dirMap["AAA"]

	for k, v := range dirMap {
		fmt.Printf("Name: %v, Left: %v Right: %v \n", k, v.Left.Name, v.Right.Name)
	}

	for next.Name != "ZZZ" {
		// fmt.Printf("Next name: %v\n", next.Name)
		i := 0
		for i < len(instructions) {
			char := string(instructions[i])
			// fmt.Printf("Checking instruction: %v \n", char)
			if char == "R" {
				next = next.Right
			} else if char == "L" {
				next = next.Left
			} else {
				fmt.Printf("INSTRUCTIONS UNCLEAR PLEASE FIX: %v \n", instructions[i])
			}

			sum++
			i++
		}
	}

	return sum, nil
}

func HauntedWastelandPartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	row := 0
	instructions := ""
	ndrMap := map[string]*NewDirectionResponse{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")

		if row > 1 {
			ndr := NewDirectionFromInput(lineString)

			if ndr.Err != nil {
				return 0, err
			}

			ndrMap[ndr.Direction.Name] = ndr
			// fmt.Printf("Added ndr with name: %v with left: %v and right: %v \n",
			// ndr.Direction.Name, ndr.Left, ndr.Right)

		} else if row == 0 {
			instructions = lineString
			row++

			fmt.Printf("Instructions: %v\n", instructions)
		} else {
			row++
			continue
		}
	}

	dirMap := map[string]*Direction{}
	// nexts := []*Direction{}
	headers := []*DirectionHead{}

	for k, ndr := range ndrMap {
		dirMap[k] = ndr.Direction
		dirMap[k].Left = ndrMap[ndr.Left].Direction
		dirMap[k].Right = ndrMap[ndr.Right].Direction

		if string(ndr.Direction.Name[2]) == "A" {
			// nexts = append(nexts, ndr.Direction)

			head := DirectionHead{
				Current:     ndr.Direction,
				OriginalDir: ndr.Direction,
				Ctr:         0,
			}

			headers = append(headers, &head)
		}
	}

	sum := 0

	fmt.Printf("Start values: \n")
	for i, head := range headers {
		fmt.Printf("Ctr: %v, Value: %v Right: %v Left: %v \n", i, head.Current.Name, head.Current.Right.Name, head.Current.Left.Name)
	}

	for !AllDone(headers) {
		i := 0

		for i < len(instructions) {
			char := string(instructions[i])

			switch char {
			case "R":
				for j := 0; j < len(headers); j++ {
					if !headers[j].Done {
						if string(headers[j].Current.Name[2]) == "Z" {
							headers[j].Done = true
						} else {
							headers[j].Current = headers[j].Current.Right
							headers[j].Ctr++
						}
					}
				}

			case "L":
				for j := 0; j < len(headers); j++ {
					if !headers[j].Done {
						if string(headers[j].Current.Name[2]) == "Z" {
							headers[j].Done = true
						} else {
							headers[j].Current = headers[j].Current.Left
							headers[j].Ctr++
						}
					}
				}

			default:
				fmt.Printf("INSTRUCTIONS UNCLEAR PLEASE FIX: %v \n", instructions[i])
			}

			i++
		}
	}

	for i, head := range headers {
		fmt.Printf("Head #%v - Original %v - Ctr: %v - Multiplier: %v - Total: %v \n",
			i, head.OriginalDir.Name, head.Ctr, head.Multiplier, head.Ctr*head.Multiplier)
	}

	LadderItUp(headers)

	fmt.Println("-------")
	for i, head := range headers {
		fmt.Printf("Head #%v - Original %v - Ctr: %v - Multiplier: %v - Total: %v \n",
			i, head.OriginalDir.Name, head.Ctr, head.Multiplier, head.Ctr*head.Multiplier)
	}

	sum = headers[0].Ctr * headers[0].Multiplier

	return sum, nil
}

func AllDone(input []*DirectionHead) bool {
	for _, head := range input {
		if !head.Done {
			return false
		}
	}

	return true
}

func IfAllEndsInZ(nexts []*Direction) bool {
	for i, next := range nexts {
		if string(next.Name[2]) != "Z" {
			if i > 3 {
				fmt.Printf("This many did end in Z: %v This one did not end in Z: %v \n", i, next.Name)
			}
			return false
		}
	}

	return true
}

func NewDirectionFromInput(input string) *NewDirectionResponse {
	segments := strings.Split(input, " ")
	name := segments[0]

	left := segments[2]
	left = strings.TrimLeft(left, "(")
	left = strings.TrimRight(left, ",")

	right := segments[3]
	right = strings.TrimRight(right, ")")

	d := NewDirection(name)

	ndr := NewDirectionResponse{
		Direction: d,
		Left:      left,
		Right:     right,
		Err:       nil,
	}

	return &ndr
}

func NewDirection(name string) *Direction {
	d := Direction{
		Name: name,
	}

	return &d
}

func LadderItUp(headers []*DirectionHead) {
	// HIGHEST
	// Ctr: 22301
	// Mlt: 1
	//
	// First head
	// Value: 16631
	// Mlt: 1

	highest := &DirectionHead{}

	for _, head := range headers {
		if head.Ctr > highest.Ctr {
			highest = head
		}
	}

	done := false

	for !done {
		done = true
		highestValue := highest.Ctr + (highest.Ctr * highest.Multiplier)
		for i := 0; i < len(headers); i++ {
			if highest == headers[i] {
				continue
			}

			headers[i].Multiplier = highestValue / headers[i].Ctr
			remainder := highestValue % headers[i].Ctr

			// fmt.Printf("highest ctr: %v, mult: %v, total: %v \n", highest.Ctr, highest.Multiplier, highestValue)

			// fmt.Printf("For value: %v, the multiplier is: %v and the remainder is: %v \n",
			// 	headers[i].Current.Name, headers[i].Multiplier, remainder)

			if remainder > 0 {
				done = false
			}
		}

		highest.Multiplier++
	}
}

func DeleteMultipleIndexInSlice(slice []*Direction, indexes []int) []*Direction {
	for _, index := range indexes {
		slice = DeleteIndexInSlice(slice, index)
	}

	return slice
}

func DeleteIndexInSlice(slice []*Direction, index int) []*Direction {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
