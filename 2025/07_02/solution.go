package seventwo

import (
	"fmt"
	"time"
)

type laser struct {
	x int
	y int
}

func (l *laser) move(dx, dy int) {
	l.x += dx
	l.y += dy
}

func (l *laser) colllisionCheck(board []string, sum int) (bool, int) {
	if board[l.y][l.x] == '^' {
		sum = l.spawn(board, sum)
		return true, sum
	}

	return false, sum
}

func (l *laser) spawn(board []string, sum int) int {
	l1 := l.mutate(true)
	sum = l1.shoot(board, sum)

	l2 := l.mutate(false)
	return sum + l2.shoot(board, sum)
}

func (l *laser) mutate(left bool) *laser {
	newLaser := *l

	if left {
		newLaser.x += -1
	} else {
		newLaser.x += 1
	}

	return &newLaser
}

func (l *laser) shoot(lines []string, sum int) int {
	board := []string{}
	board = append(board, lines...)

	for {
		board[l.y] = board[l.y][:l.x] + "|" + board[l.y][l.x+1:]
		lines[l.y] = lines[l.y][:l.x] + "|" + lines[l.y][l.x+1:]
		for _, b := range lines {
			fmt.Println(b)
		}

		l.move(0, 1)
		time.Sleep(10 * time.Millisecond)
		if l.y == len(board)-1 {
			sum++
			return sum
		}

		col, ends := l.colllisionCheck(lines, sum)
		if col {
			return ends
		}
	}
}

func SevenTwo(lines []string) int {
	sum := 0
	originLaser := FormatData(lines)
	sum = originLaser.shoot(lines, sum)

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
