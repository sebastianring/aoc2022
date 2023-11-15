package daytwo

import (
	"bufio"
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

type hand byte

const (
	rock    hand = 0
	paper   hand = 1
	scissor hand = 2
)

func RockPaperScissor(filename string) (int, error) {
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

	playerMap := map[byte]hand{
		'X': rock,
		'Y': paper,
		'Z': scissor,
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
		player := playerMap[lineString[2]]

		roundpoints := PlayRockPaperScissor(enemy, player)

		points += roundpoints
	}

	return points, nil
}

func PlayRockPaperScissor(enemy hand, player hand) int {
	sum := 0

	if enemy == player {
		sum += 3
	} else {
		winCondition := map[hand]hand{
			paper:   rock,
			rock:    scissor,
			scissor: paper,
		}

		if enemy == winCondition[player] {
			sum += 6
		}
	}

	pointMap := map[hand]int{
		rock:    1,
		paper:   2,
		scissor: 3,
	}

	sum += pointMap[player]

	return sum
}
