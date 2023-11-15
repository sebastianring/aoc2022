package daytwo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "strings"
)

// column 1: what oppnent plays -- column 2: what you play
// points:
// lose = 0
// draw = 3
// win = 6
//
// rock = 1
// paper = 2
// scissor = 3
//
// lose with rock = 0
// win with rock = 7
// win with scissor = 9
// draw with paper = 5

type forcedResult byte

const (
	win  forcedResult = 0
	draw forcedResult = 1
	lose forcedResult = 2
)

func RockPaperScissorPartTwo(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Println("Error opening file: ", err)
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	points := 0

	enemyMap := map[byte]hand{
		'A': rock,
		'B': paper,
		'C': scissor,
	}

	resultMap := map[byte]forcedResult{
		'X': lose,
		'Y': draw,
		'Z': win,
	}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Println("Error reading line: ", err)
			return 0, err
		}

		enemy := enemyMap[lineString[0]]
		// player := playerMap[lineString[2]]
		result := resultMap[lineString[2]]

		roundpoints := PlayRockPaperScissorPartTwo(enemy, result)

		points += roundpoints
	}

	return points, nil
}

func PlayRockPaperScissorPartTwo(enemy hand, result forcedResult) int {
	// Need to force the result
	// forcedResult = draw
	// enemy = rock
	// then player should chose rock
	sum := 0
	var player hand

	winCondition := map[hand]hand{
		paper:   rock,
		rock:    scissor,
		scissor: paper,
	}

	if result == draw {
		player = enemy
		sum += 3
	} else if result == win {
		player = winCondition[enemy]
		sum += 6
	} else {
		for k, v := range winCondition {
			if v == enemy {
				player = k
				break
			}
		}
	}

	pointMap := map[hand]int{
		rock:    1,
		paper:   2,
		scissor: 3,
	}

	sum += pointMap[player]
	fmt.Println("Result: ", result, "enemy: ", enemy, "player: ", player, "sum: ", sum)

	return sum
}
