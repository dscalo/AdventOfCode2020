package puzzles

import (
	"bufio"
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/fs"
	"github.com/dscalo/AdventOfCode2020/internal/geometry"
	"os"
	"strings"
)

func init() {
	Days[11] = Day11
}

var directions = map[string]*geometry.Point{
	"up":        geometry.NewPoint(0, -1),
	"right":     geometry.NewPoint(1, 0),
	"down":      geometry.NewPoint(0, 1),
	"left":      geometry.NewPoint(-1, 0),
	"upRight":   geometry.NewPoint(1, -1),
	"upLeft":    geometry.NewPoint(-1, -1),
	"downRight": geometry.NewPoint(1, 1),
	"downLeft":  geometry.NewPoint(-1, 1),
}

func onGrid(grid [][]string, pt geometry.Point) bool {
	if pt.Y < 0 || pt.Y >= len(grid) {
		return false
	}

	if pt.X < 0 || pt.X >= len(grid[pt.Y]) {
		return false
	}

	return true
}

func readSeatLayout(path string) [][]string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	size, err := fs.LineCount(path)

	grid := make([][]string, size)

	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		grid[idx] = row
		idx++
	}
	return grid
}

func los(grid [][]string, pt geometry.Point, move geometry.Point) string {
	valid := true

	for valid {
		pt.MovePoint(&move)
		if onGrid(grid, pt) {
			if grid[pt.Y][pt.X] != "." {
				return grid[pt.Y][pt.X]
			}
		} else {
			valid = false
		}
	}

	return "."
}

func visibilityCount(grid [][]string, y int, x int) int {
	ct := 0
	for _, pt := range directions {
		if los(grid, *geometry.NewPoint(x, y), *pt) == "#" {
			ct++
		}
	}
	return ct
}

func adjacentCount(grid [][]string, y int, x int) int {
	count := 0
	// top y-1 x+ 0
	if y-1 >= 0 && grid[y-1][x] == "#" {
		count++
	}
	// dia top left y-1  x-1
	if y-1 >= 0 && x-1 >= 0 && grid[y-1][x-1] == "#" {
		count++
	}

	// dia top right y -1 x + 1
	if y-1 >= 0 && x+1 < len(grid[y-1]) && grid[y-1][x+1] == "#" {
		count++
	}

	// left y+0 x - 1
	if x-1 >= 0 && grid[y][x-1] == "#" {
		count++
	}

	// right y+0 x+ 1
	if x+1 < len(grid[y]) && grid[y][x+1] == "#" {
		count++
	}

	// bottom y +1 x +0
	if y+1 < len(grid) && grid[y+1][x] == "#" {
		count++
	}
	// dia bottom left y+1  x-1
	if y+1 < len(grid) && x-1 >= 0 && grid[y+1][x-1] == "#" {
		count++
	}

	// dia bottom right y + 1 x + 1
	if y+1 < len(grid) && x+1 < len(grid[y+1]) && grid[y+1][x+1] == "#" {
		count++
	}

	return count
}

func seating(grid [][]string, part int) int {
	var occp []geometry.Point
	var empt []geometry.Point

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			switch grid[y][x] {
			case "L":
				if part == 1 {
					if adjacentCount(grid, y, x) == 0 {
						occp = append(occp, *geometry.NewPoint(x, y))
					}
				} else {
					if visibilityCount(grid, y, x) == 0 {
						occp = append(occp, *geometry.NewPoint(x, y))
					}
				}

			case "#":
				if part == 1 {
					if adjacentCount(grid, y, x) >= 4 {
						empt = append(empt, *geometry.NewPoint(x, y))
					}
				} else {
					if visibilityCount(grid, y, x) >= 5 {
						empt = append(empt, *geometry.NewPoint(x, y))
					}
				}

			case ".":
			default:
				fmt.Printf("grid y: %d / x: %d is %s\n", y, x, grid[y][x])
			}
		}
	}

	for _, pt := range occp {
		grid[pt.Y][pt.X] = "#"
	}

	for _, pt := range empt {
		grid[pt.Y][pt.X] = "L"
	}

	return len(occp) + len(empt)
}

func simulateSeating(grid [][]string, part int) int {
	for seating(grid, part) > 0 {
	}
	return countSeats(grid)
}

func countSeats(grid [][]string) int {
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "#" {
				count++
			}
		}
	}

	return count
}

func Day11() {
	inputs := []string{"test01", "puzzle"} //
	for _, f := range inputs {
		path := fmt.Sprintf("input/day11/%s.input", f)

		grid := readSeatLayout(path)
		//pretty.Print2dString(grid)
		//println("___________________________")

		ansP1 := simulateSeating(grid, 1)
		ansP2 := simulateSeating(readSeatLayout(path), 2)

		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
