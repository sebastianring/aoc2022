package thirteen

import (
	"bufio"
	"fmt"
	"os"
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
			p.x += 10000000000000
			p.y += 10000000000000
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

func (s *scenario) converge() (bool, int) {
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

	// x := s.prize.x
	// y := s.prize.y
	// i := 0

	//
	// for x < s.prize.x && y < s.prize.y {
	// 	xmod, xrem := DivMod((s.prize.x - x), primary.x)
	// 	ymod, yrem := DivMod((s.prize.y - y), primary.y)
	//
	// 	if xmod == ymod && xrem == 0 && yrem == 0 {
	// 		fmt.Printf("convergence found at %d %d %d\n", x, y, i)
	// 		return true, i + (primaryCost * xmod)
	// 	}
	//
	// 	// if x%primary.x == 0 && y%primary.y == 0 {
	// 	// 	fmt.Printf("convergence found at %d %d %d\n", x, y, i)
	// 	// 	return true, i
	// 	// }
	//
	// 	x -= secondary.x
	// 	y -= secondary.y
	// 	i += secondaryCost
	// }
	//
	// if x == s.prize.x && y == s.prize.y {
	// 	return true, i
	// }
	//
	// return false, i
}

func DivMod(a, b int) (quotient, remainder int) {
	return a / b, a % b
}
