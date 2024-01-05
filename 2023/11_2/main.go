package dayeleventwo2023

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func DayEleven(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

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

		y++
	}

	universe.Print()
	voids := universe.CheckVoids()
	fmt.Printf("Voids rows: %v cols: %v \n", voids.Rows, voids.Cols)

	//universe.ExpandUniverse(voids)
	//universe.Print()
	galaxies := universe.GetAllGalaxies()
	fmt.Printf("Number of galaxies: %v\n", len(galaxies))

	distances := WalkAllDistances(galaxies, voids)

	sum = Sum(distances)

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

func (u *Universe) CheckVoids() Voids {
	voids := Voids{}
	columnHit := map[int]bool{}

	for y, row := range *u {
		rowHit := false
		for x, char := range row {
			if string(char) == "#" {
				rowHit = true
				// fmt.Printf("Non void row: %v \n", y)

				if _, ok := columnHit[x]; !ok {
					columnHit[x] = true
				}

				//break
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
	rowCtr := 0
	colCtr := 0

	counted := 0

	for i := 0; i < len(voids.Rows)+len(voids.Cols); i++ {
		//fmt.Printf("ROW: rowctr: %v, colctr: %v, rows: %v, cols: %v \n",
		//	rowCtr, colCtr, len((*u)[0]), len(*u))

		//fmt.Printf("COL: rowctr: %v, colctr: %v, rows: %v, cols: %v \n",
		//	rowCtr, colCtr, len((*u)[0]), len(*u))
		//
		//fmt.Printf("ROW: total: %v, rows: %v rowctr: %v, colctr: %v, rows: %v, cols: %v \n",
		//	voids.Rows[rowCtr]+rowCtr, voids.Rows[rowCtr], rowCtr, colCtr, len((*u)[0]), len(*u))
		//
		//fmt.Printf("COL: total: %v, cols: %v rowctr: %v, colctr: %v, rows: %v, cols: %v \n",
		//	voids.Cols[colCtr]+colCtr, voids.Cols[colCtr], rowCtr, colCtr, len((*u)[0]), len(*u))

		if rowCtr < len(voids.Rows) && colCtr < len(voids.Cols) && voids.Rows[rowCtr] <= voids.Cols[colCtr] {
			u.ExpandRow(voids.Rows[rowCtr] + rowCtr)
			rowCtr++

		} else if colCtr < len(voids.Cols) {
			u.ExpandCol(voids.Cols[colCtr] + colCtr)
			colCtr++

		} else {
			fmt.Printf("Something is wrong: rowctr: %v, colctr: %v, rows: %v, cols: %v counted: %v\n",
				rowCtr, colCtr, len((*u)[0]), len(*u), counted)
		}

		counted++
	}

	//fmt.Printf("Counted this many new expansions: %v\n", counted)
}

func (u *Universe) ExpandRow(row int) {
	tempSlice := *u
	tempSlice = append(tempSlice, "")
	newRow := u.GetRowWithDots()

	copy(tempSlice[row+1:], tempSlice[row:])
	tempSlice[row] = newRow

	*u = tempSlice
}

func (u *Universe) ExpandCol(col int) {
	// fmt.Printf("Trying to expand col: %v, current cols: %v \n", col, len((*u)[0]))

	for i := 0; i < len(*u); i++ {
		prefix := (*u)[i][:col]
		suffix := (*u)[i][col:]

		(*u)[i] = prefix + string("C") + suffix

	}
}

func (u *Universe) GetRowWithDots() string {
	s := ""

	for i := 0; i < len((*u)[0]); i++ {
		s += "R"
	}

	return s
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

func (u *Universe) GetAllGalaxies() []*Galaxy {
	var galaxies []*Galaxy

	for y, row := range *u {
		for i := 0; i < len(row); i++ {
			if row[i] == '#' {
				g := NewGalaxy(Pos{
					x: i,
					y: y,
				},
					len(galaxies)+1)

				galaxies = append(galaxies, g)
				fmt.Printf("Galaxy #%v at x: %v y: %v \n", g.NumberString(), g.x, g.y)
			}
		}
	}

	//yCount := 0
	//for y, row := range *u {
	//	xCount := 0
	//	for x, char := range row {
	//		if string(char) == "#" {
	//			g := NewGalaxy(Pos{x: x, y: y}, len(galaxies)+1)
	//
	//			galaxies = append(galaxies, g)
	//
	//			fmt.Printf("Galaxy #%v at x: %v y: %v \n", g.NumberString(), g.x, g.y)
	//		}
	//		xCount++
	//	}
	//	fmt.Printf("Counter xchars: %v\n", xCount)
	//	yCount++
	//}
	//fmt.Printf("Counter ychars: %v\n", yCount)

	fmt.Printf("Number of galaxies: %v\n", len(galaxies))
	return galaxies
}

func (u *Universe) Print() {
	for y := 0; y < len(*u); y++ {
		for x := 0; x < len((*u)[y]); x++ {
			fmt.Printf("%v", string((*u)[y][x]))
		}
		fmt.Printf("\n")
	}
	//for i, row := range *u {
	//	fmt.Printf("#%v - %v\n", i+1, row)
	//}
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

func WalkAllDistances(galaxies []*Galaxy, voids Voids) []int {
	var results []int
	sum := 0

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			//fmt.Printf("I: %v, J: %v Length: %v\n", i, j, len(results))
			distance := WalkDistance(galaxies[i], galaxies[j], voids)
			sum += distance
			results = append(results, distance)
		}
	}

	fmt.Printf("Sum: %v Length of results %v\n", sum, len(results))
	return results
}

func WalkDistance(galaxyOne *Galaxy, galaxyTwo *Galaxy, voids Voids) int {
	fmt.Printf("Walking the distance between: #%v, #%v\n", galaxyOne.Number, galaxyTwo.Number)
	voidCtr := 0

	for _, row := range voids.Rows {
		if galaxyOne.y < row && galaxyTwo.y > row ||
			galaxyTwo.y < row && galaxyOne.y > row {
			voidCtr++
		}
	}

	for _, row := range voids.Cols {
		if galaxyOne.x < row && galaxyTwo.x > row ||
			galaxyTwo.x < row && galaxyOne.x > row {
			voidCtr++
		}
	}

	diffX := Diff(galaxyOne.x, galaxyTwo.x)
	diffY := Diff(galaxyOne.y, galaxyTwo.y)

	result := diffX + diffY + (voidCtr * 1000000) - voidCtr

	fmt.Printf("There were %v voids in between, thus the final result is: %v\n", voidCtr, result)
	return result
}

func Diff(a int, b int) int {
	sum := a - b

	if sum < 0 {
		sum *= -1
	}

	return sum
}

func MeasureAllGalaxyDistances(galaxies []Galaxy) []int {
	var results []int

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := MeasureDistance(galaxies[i], galaxies[j])
			results = append(results, distance)
		}
	}

	return results
}

func MeasureDistance(galaxyOne Galaxy, galaxyTwo Galaxy) int {
	sum := math.Pow(float64(galaxyTwo.x-galaxyOne.x), 2) +
		math.Pow(float64(galaxyTwo.y-galaxyOne.y), 2)

	distance := math.Sqrt(sum)

	fmt.Printf("Measuring distance between: #%v, %v, %v, and #%v, %v, %v\n",
		galaxyOne.Number, galaxyOne.x, galaxyOne.y, galaxyTwo.Number, galaxyTwo.x, galaxyTwo.y)
	fmt.Printf("Distance: %v\n", distance)
	return int(distance)
}

func Sum(values []int) int {
	sum := 0

	for _, val := range values {
		sum += val
	}

	return sum
}

func PrintDummyUniverse(u *Universe, galaxies []*Galaxy) {
	dummyUni := *u

	for _, galaxy := range galaxies {
		row := dummyUni[galaxy.y]
		modified := ChangeCharAtIndex(row, galaxy.x, '@')
		dummyUni[galaxy.y] = modified
	}

	dummyUni.Print()
}

func ChangeCharAtIndex(str string, index int, newChar rune) string {
	runes := []rune(str)

	if index >= 0 && index < len(runes) {
		runes[index] = newChar
		return string(runes)
	}

	fmt.Printf("Out of bounds\n")
	return str
}
