package puzzles

import (
	"bufio"
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/fs"
	"github.com/dscalo/AdventOfCode2020/internal/geometry"
	"os"
	"strings"
)

type Grid = [][]string

func readFile3(path string) [][]string {
	size, err := fs.LineCount(path)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	grid := make(Grid, size)

	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid[idx] = strings.Split(line, "")
		idx++
	}

	return grid
}

func goTobogganing(grid Grid, move *geometry.Point) int {
	trees := 0
	curPos := geometry.NewPoint(0, 0)

	gridWidth := len(grid[0]) - 1

	for curPos.Y < len(grid) {
		if curPos.X > gridWidth {
			curPos.X = curPos.X - gridWidth - 1
		}

		if grid[curPos.Y][curPos.X] == "#" {
			trees++
		}

		curPos.Add(move)
	}
	return trees
}

func Day03() {
	inputs := []string{"test01", "puzzle"} //

	part2Moves := geometry.NewPoints([][]int{{1, 1}, {5, 1}, {7, 1}, {1, 2}})

	for _, f := range inputs {
		path := fmt.Sprintf("input/day03/%s.input", f)
		grid := readFile3(path)
		//pretty.Print2dString(grid)
		moveP1 := geometry.NewPoint(3, 1)
		ansP1 := goTobogganing(grid, moveP1)

		ansP2 := ansP1
		for _, move := range *part2Moves {
			ansP2 *= goTobogganing(grid, &move)

		}

		fmt.Printf("%s trees hit: %d\n", f, ansP1)
		fmt.Printf("(PART2) %s TOTAL trees hit: %d\n", f, ansP2)
	}
}
