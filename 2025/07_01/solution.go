package sevenone

import "fmt"

type laser struct {
	x int
	y int
}

func (l *laser) move(dx, dy int) {
	l.x += dx
	l.y += dy
}

func (l *laser) split() []*laser {
	l1 := laser{
		x: l.x - 1,
		y: l.y,
	}

	l2 := laser{
		x: l.x + 1,
		y: l.y,
	}

	return []*laser{&l1, &l2}
}

func Template(lines []string) int {
	sum := 0
	lasers := []*laser{FormatData(lines)}

	for range len(lines) - 1 {
		fmt.Println("-----")
		for _, l := range lasers {
			l.move(0, 1)
			if string(lines[l.y][l.x]) == "^" {
				split := l.split(line[y])
			}

			lines[l.y] = lines[l.y][:l.x] + "|" + lines[l.y][l.x+1:]
		}

		for _, line := range lines {
			fmt.Println(line)
		}
	}

	return sum
}

func FormatData(lines []string) *laser {
	for x, v := range lines[0] {
		if string(v) == "S" {
			return &laser{
				x: x,
				y: 0,
			}
		}
	}

	return nil
}
