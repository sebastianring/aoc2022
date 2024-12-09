package daysixparttwo

//
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )
//
// type pos struct {
// 	x int
// 	y int
// }
//
// type Beam struct {
// 	x      int
// 	y      int
// 	xSpeed int
// 	ySpeed int
// }
//
// type hits struct {
// }
//
// func DaySixPartTwo(filename string) (int, error) {
// 	file, err := os.Open(filename)
//
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	defer file.Close()
//
// 	reader := bufio.NewReader(file)
// 	sum := 0
// 	_ = sum
// 	board := [][]string{}
//
// 	for {
// 		lineString, err := reader.ReadString('\n')
//
// 		if err != nil {
// 			if err.Error() == "EOF" {
// 				break
// 			}
//
// 			return 0, err
// 		}
//
// 		lineString = strings.TrimSuffix(lineString, "\n")
// 		currentLine := []string{}
// 		for _, char := range lineString {
// 			currentLine = append(currentLine, string(char))
// 		}
//
// 		board = append(board, currentLine)
// 	}
//
// 	blocks := []pos{}
//
// 	for y := 0; y < len(board); y++ {
// 		for x := 0; x < len(board[y]); x++ {
// 			if board[y][x] == "#" {
// 				blocks = append(blocks, pos{x: x, y: y})
// 			}
// 		}
// 	}
//
// 	fmt.Println(blocks)
//
// 	return 0, nil
// }
//
// func (b *Beam) Rotate() {
// 	if b.xSpeed == 0 && b.ySpeed == -1 {
// 		b.xSpeed = 1
// 		b.ySpeed = 0
// 	} else if b.xSpeed == 1 && b.ySpeed == 0 {
// 		b.xSpeed = 0
// 		b.ySpeed = 1
// 	} else if b.xSpeed == 0 && b.ySpeed == 1 {
// 		b.xSpeed = -1
// 		b.ySpeed = 0
// 	} else if b.xSpeed == -1 && b.ySpeed == 0 {
// 		b.xSpeed = 0
// 		b.ySpeed = -1
// 	}
// }
