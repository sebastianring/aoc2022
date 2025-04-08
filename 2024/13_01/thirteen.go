package thirteen

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type scenario struct {
	buttonA pos
	buttonB pos
	prize   pos
}

type pos struct {
	x int
	y int
}

func Thirteen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	i := 0
	scenarios := []*scenario{}
	currentScenario := &scenario{}

	for {
		lineString, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			return 0, err
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		parts := strings.Split(lineString, ":")
		// for ctr, part := range parts {
		// 	fmt.Printf("ctr #%d: %v\n", ctr, part)
		// }

		var p pos

		if len(parts) > 1 {
			if strings.Contains("Prize", parts[0]) {
				xy := strings.Split(parts[1], ",")
				xstr := strings.TrimPrefix(xy[0], " X=")
				ystr := strings.TrimPrefix(xy[1], " Y=")

				x, _ := strconv.Atoi(xstr)
				y, _ := strconv.Atoi(ystr)

				p.x = x
				p.y = y
			} else {
				xy := strings.Split(parts[1], ",")
				xstr := strings.TrimPrefix(xy[0], " X+")
				ystr := strings.TrimPrefix(xy[1], " Y+")

				x, _ := strconv.Atoi(xstr)
				y, _ := strconv.Atoi(ystr)

				p.x = x
				p.y = y
			}
		}

		switch parts[0] {
		case "Button A":
			currentScenario.buttonA = p
		case "Button B":
			currentScenario.buttonB = p
		case "Prize":
			currentScenario.prize = p
		}

		if strings.Contains("\n", lineString) {
			scenarios = append(scenarios, currentScenario)
			currentScenario = &scenario{}
			i = 0
		} else {
			i++
		}
	}

	for _, scenario := range scenarios {
		fmt.Println("--------")
		fmt.Println(scenario)

		ok, val := scenario.converge()
		fmt.Println(ok, val)
		if ok {
			sum += val
		}

		fmt.Println("--------")
	}

	return sum, nil
}

func permutate(s *scenario) int {
	if s == nil {
		fmt.Printf("nil fail\n")
		return 0
	}

	x := 0
	y := 0

	btns := sortedButtons(s.buttonA, s.buttonB)

	buttonA := 0
	buttonB := 0

	for x < s.prize.x && y < s.prize.y {
		leftX := s.prize.x - x
		leftY := s.prize.y - y

		if leftX%btns[1].x == 0 && leftY%btns[1].y == 0 && leftX/btns[1].x == leftY/btns[1].y {
			multiplier := leftX / btns[1].x
			x += btns[1].x * multiplier
			y += btns[1].y * multiplier
			buttonB += leftX / btns[1].x
		} else {
			x += btns[0].x
			y += btns[0].y
			buttonA += 1
		}
	}

	return buttonA*3 + buttonB
}

func sortedButtons(a, b pos) []pos {
	l := []pos{a, b}
	slices.SortFunc(l, func(a, b pos) int {
		one := a.x + a.y
		two := b.x + b.y
		if one > two {
			return 1
		} else if one == two {
			return 0
		}

		return -1
	})

	return l
}

func (s *scenario) converge() (bool, int) {
	x := 0
	y := 0
	i := 0

	// do i need to consider value? most likely
	valueRatio := float32((s.buttonB.x + s.buttonB.y) / (s.buttonA.x + s.buttonB.y))
	fmt.Printf("value ratio: %v - sum b: %v a: %v\n", valueRatio, (s.buttonB.x + s.buttonB.y), (s.buttonA.x + s.buttonA.y))

	var primary pos
	primaryCost := 1
	var secondary pos
	secondaryCost := 3

	if valueRatio >= 3.0 {
		fmt.Printf("high value ratio %v\n", valueRatio)
		fmt.Printf("=========================\n")
		primary = s.buttonA
		primaryCost = 3
		secondary = s.buttonB
		secondaryCost = 1
	} else {
		fmt.Printf("low value ratio %v\n", valueRatio)
		primary = s.buttonB
		secondary = s.buttonA
	}

	for x < s.prize.x && y < s.prize.y {
		xmod, xrem := DivMod((s.prize.x - x), primary.x)
		ymod, yrem := DivMod((s.prize.y - y), primary.y)

		if xmod == ymod && xrem == 0 && yrem == 0 {
			fmt.Printf("convergence found at %d %d %d\n", x, y, i)
			return true, i + (primaryCost * xmod)
		}

		// if x%primary.x == 0 && y%primary.y == 0 {
		// 	fmt.Printf("convergence found at %d %d %d\n", x, y, i)
		// 	return true, i
		// }

		x += secondary.x
		y += secondary.y
		i += secondaryCost
	}

	if x == s.prize.x && y == s.prize.y {
		return true, i
	}

	return false, i
}

func DivMod(a, b int) (quotient, remainder int) {
	return a / b, a % b
}
