package dayeleven2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DayEleven(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	y := 0
	universe := Universe{}

	for {
		lineString, err := reader.ReadString('\n')

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}

		lineString = strings.TrimSuffix(lineString, "\n")
		universe = append(universe, lineString)

		// for x, char := range lineString {
		// 	if string(char) == "#" {
		// 		nonVoidColumns = append(nonVoidColumns, x)
		// 		nonVoidRows = append(nonVoidRows, y)
		//
		// 		g := NewGalaxy(Pos{x: x, y: y}, len(galaxies)+1)
		//
		// 		galaxies = append(galaxies, g)
		// 	}
		// }

		y++
	}

	galaxies := universe.GetAllGalaxies()
	voids := universe.CheckVoids(galaxies)

	universe.ExpandUniverse(voids)

	fmt.Printf("Voids rows: %v cols: %v \n", voids.Rows, voids.Cols)

	return sum, nil
}

type Pos struct {
	x int
	y int
}

type Galaxy struct {
	Number int
	Pos
}

type Voids struct {
	Rows []int
	Cols []int
}

type Universe []string

func (u *Universe) CheckVoids(galaxies []Galaxy) Voids {
	voids := Voids{}
	columnHit := map[int]bool{}

	for y, row := range *u {
		rowHit := false
		for x, char := range row {
			if string(char) == "#" {
				rowHit = true
				fmt.Printf("Non void row: %v \n", y)

				_, ok := columnHit[x]

				if !ok {
					columnHit[x] = true
				}

				break
			}
		}

		if rowHit == false {
			voids.Rows = append(voids.Rows, y)
		}
	}

	for x := range (*u)[0] {
		if _, ok := columnHit[x]; !ok {
			voids.Cols = append(voids.Cols, x)
		}
	}

	return voids
}

func (u *Universe) ExpandUniverse(voids Voids) {

}

func GetStandardVoids(u Universe) Voids {
	ctrRows := len(u)
	ctrCols := len((u)[0])

	v := Voids{}

	for i := 0; i < max(ctrRows, ctrCols); i++ {
		if i < ctrCols {
			v.Cols = append(v.Cols, i)
		}

		if i < ctrRows {
			v.Rows = append(v.Rows, i)
		}
	}

	return v
}

func (u *Universe) GetAllGalaxies() []Galaxy {
	galaxies := []Galaxy{}

	for y, row := range *u {
		for x, char := range row {
			if string(char) == "#" {
				g := Galaxy{
					Number: len(galaxies) + 1,
					Pos:    Pos{x, y},
				}

				galaxies = append(galaxies, g)

				fmt.Printf("Galaxy #%v at x: %v y: %v \n", g.NumberString(), g.x, g.y)
			}
		}
	}

	return galaxies
}

func NewGalaxy(pos Pos, number int) *Galaxy {
	g := Galaxy{
		Pos:    pos,
		Number: number,
	}

	return &g
}

func (g *Galaxy) NumberString() string {
	return strconv.Itoa(g.Number)
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
