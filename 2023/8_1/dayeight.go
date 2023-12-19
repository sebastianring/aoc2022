package dayeight2023

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
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
	nexts := []*Direction{}

	for k, ndr := range ndrMap {
		dirMap[k] = ndr.Direction
		dirMap[k].Left = ndrMap[ndr.Left].Direction
		dirMap[k].Right = ndrMap[ndr.Right].Direction

		if string(ndr.Direction.Name[2]) == "A" {
			nexts = append(nexts, ndr.Direction)
			fmt.Printf("Added start next value: %v \n", ndr.Direction.Name)
		}
	}

	sum := 0

	for k, v := range dirMap {
		fmt.Printf("Name: %v, Left: %v Right: %v \n", k, v.Left.Name, v.Right.Name)
	}

	var wg sync.WaitGroup
	safeCounter := SafeCounter{
		counter: 0,
	}

	for _, next := range nexts {
		wg.Add(1)
		go func(next *Direction) {
			for string(next.Name[2]) != "Z" {
				fmt.Printf("Next name: %v last char: %v sum: %v \n", next.Name, string(next.Name[2]), sum)
				i := 0

				for i < len(instructions) {
					char := string(instructions[i])
					fmt.Printf("Checking instruction: %v \n", char)

					if char == "R" {
						next = next.Right
					} else if char == "L" {
						next = next.Left
					} else {
						fmt.Printf("INSTRUCTIONS UNCLEAR PLEASE FIX: %v \n", instructions[i])
					}

					safeCounter.Increment()
					i++
				}
			}

			wg.Done()
		}(next)
	}

	wg.Wait()
	sum = safeCounter.counter

	return sum, nil
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
