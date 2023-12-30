package daytentwo2023

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func DayTen(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	pipeMap := []string{}
	startPos := Pos{
		x: -1,
		y: -1,
	}
	y := 0

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		pipeMap = append(pipeMap, lineString)

		if startPos.x == -1 {
			for x, s := range lineString {
				if string(s) == "S" {
					startPos.x = x
					startPos.y = y
				}
			}
		}

		y++
	}

	for _, row := range pipeMap {
		fmt.Printf("%v\n", row)
	}

	fmt.Printf("Pos x: %v y: %v\n", startPos.x, startPos.y)
	directions := GetDirectionsFromStart(startPos, pipeMap)

	// for _, dirction := range directions {
	// 	fmt.Println(dirction.String())
	// }

	points2D := []Pos{}

	if IsTurn(directions) {
		fmt.Printf("Start pos is a turn!!\n")
		points2D = append(points2D, startPos)
	}

	steps := 0

	for i, direction := range directions {
		if i > 0 {
			break
		}

		fmt.Printf("New starting direction: %v \n", direction.String())

		p := Player{
			Pos:       startPos,
			Direction: direction,
		}

		for !IsPlayerAtStartPos(startPos, p.Pos) || steps == 0 {
			p.Move(pipeMap, &points2D)
			steps++
		}

		fmt.Printf("Steps taken: %v \n", steps)
	}

	points2D = append(points2D, points2D[0])

	// for i, pos := range points2D {
	// 	fmt.Printf("Point #%v x: %v y: %v \n", i, pos.x, pos.y)
	// }

	sum = Shoelace(points2D, steps)

	return sum, nil
}

func IsPlayerAtStartPos(startPos Pos, playerPos Pos) bool {
	// fmt.Printf("Startpos: %v %v playerPos: %v %v \n", startPos.x, startPos.y, playerPos.x, playerPos.y)
	if startPos.x == playerPos.x && startPos.y == playerPos.y {
		return true
	}

	return false
}

func (p *Player) Move(pipeMap []string, points *[]Pos) {
	newPos := GetNextPos(p.Pos, p.Direction)
	pipe := GetPipeFromPos(newPos, pipeMap)
	// fmt.Printf("Player moved from: %v %v to %v %v \n", p.x, p.y, newPos.x, newPos.y)

	if pipe.Turn {
		*points = append(*points, newPos)
	}

	newDir, err := pipe.GetOutFromIn(p.Direction)

	if err != nil {
		panic(err)
	}

	p.Direction = newDir
	p.Pos = newPos
}

func GetPipeFromLetter(letter byte) Pipe {
	// fmt.Printf("Trying to fine pipe from letter: %v \n", letter)

	letterConvertMap := map[byte]Pipe{
		'F': Pipe{
			AllowedInOut: []InOut{{
				In:  Left,
				Out: Down,
			}, {
				In:  Up,
				Out: Right,
			}},
			Turn: true,
		},
		'J': Pipe{
			AllowedInOut: []InOut{{
				In:  Right,
				Out: Up,
			}, {
				In:  Down,
				Out: Left,
			}},
			Turn: true,
		},
		'7': Pipe{
			AllowedInOut: []InOut{{
				In:  Right,
				Out: Down,
			}, {
				In:  Up,
				Out: Left,
			}},
			Turn: true,
		},
		'L': Pipe{
			AllowedInOut: []InOut{{
				In:  Left,
				Out: Up,
			}, {
				In:  Down,
				Out: Right,
			}},
			Turn: true,
		},
		'-': Pipe{
			AllowedInOut: []InOut{{
				In:  Right,
				Out: Right,
			}, {
				In:  Left,
				Out: Left,
			}},
			Turn: false,
		},
		'|': Pipe{
			AllowedInOut: []InOut{{
				In:  Up,
				Out: Up,
			}, {
				In:  Down,
				Out: Down,
			}},
			Turn: false,
		},
		'S': Pipe{
			AllowedInOut: []InOut{{
				In:  Up,
				Out: Up,
			}, {
				In:  Down,
				Out: Down,
			}, {
				In:  Right,
				Out: Right,
			}, {
				In:  Left,
				Out: Left,
			}},
			Turn: false,
		},
	}

	return letterConvertMap[letter]
}

func GetDirectionsFromStart(pos Pos, pipeMap []string) []Direction {
	directions := []Direction{}

	for direction := Direction(0); direction < 4; direction++ {
		newPos := GetNextPos(pos, direction)

		if newPos.y < 0 || newPos.y > len(pipeMap)-1 || newPos.x < 0 || newPos.x > len(pipeMap[newPos.y])-1 {
			continue
		}
		// fmt.Printf("Checking direction: %v New pos x: %v y: %v \n", direction.String(), newPos.x, newPos.y)
		pipe := GetPipeFromPos(newPos, pipeMap)

		_, err := pipe.GetOutFromIn(direction)

		if err != nil {
			continue
		}

		// fmt.Printf("Found a valid direction: %v \n", direction.String())
		directions = append(directions, direction)
	}

	return directions
}

func IsTurn(directions []Direction) bool {
	if directions[0] == Up && directions[1] == Down ||
		directions[0] == Down && directions[1] == Up ||
		directions[0] == Right && directions[1] == Left ||
		directions[0] == Left && directions[1] == Right {
		return false
	}

	return true
}

func GetPipeFromPos(pos Pos, pipeMap []string) Pipe {
	char := pipeMap[pos.y][pos.x]

	return GetPipeFromLetter(char)
}

func GetNextPos(pos Pos, dir Direction) Pos {
	switch dir {
	case Up:
		pos.y--
	case Down:
		pos.y++
	case Left:
		pos.x--
	case Right:
		pos.x++
	}

	return pos
}

type Player struct {
	Pos
	Direction
}

type Pos struct {
	x int
	y int
}

type Direction byte

const (
	Up    Direction = 0
	Down  Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

type InOut struct {
	In  Direction
	Out Direction
}

type Pipe struct {
	AllowedInOut []InOut
	Value        byte
	Turn         bool
}

func (p Pipe) GetOutFromIn(in Direction) (Direction, error) {
	// fmt.Printf("Current In direction: %v\n", in.String())
	for _, InOut := range p.AllowedInOut {
		// fmt.Printf("Trying to compare to this InOut: %v \n", InOut.In.String())
		if InOut.In == in {
			return InOut.Out, nil
		}
	}

	return 0, errors.New("In direction does not exist.")
}

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"

	default:
		return ""
	}
}

func Shoelace(input []Pos, steps int) int {
	// positive := 0
	// negative := 0
	n := len(input)

	area := 0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += input[i].x * input[j].y
		area -= input[j].x * input[i].y
	}
	//
	// for i, pos := range input {
	// 	fmt.Printf("I: #%v \n", i)
	// 	if i < len(input)-1 {
	// 		fmt.Printf("Triggered positive.\n")
	// 		positive += pos.x * input[i+1].y
	// 	}
	//
	// 	if i > 0 {
	// 		fmt.Printf("Triggered negative.\n")
	// 		negative += pos.x * input[i-1].y
	// 	}
	// }
	//
	// sum := positive - negative

	fmt.Printf("Sum: %v Steps: %v \n", area, steps)

	if area < 0 {
		area *= -1
	}

	area -= steps
	area /= 2
	area += 1

	fmt.Printf("Sum: %v Steps: %v FloatSum:  \n", area, steps)

	return area
}
