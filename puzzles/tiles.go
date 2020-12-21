package puzzles

import (
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/pretty"
	"sort"
)

type Tiles = []Tile

type Tile struct {
	Id     int
	Grid   Grid
	Border map[string]string
	Lock   bool
}

var Borders = []string{"top", "right", "bottom", "lft"}

func (t *Tile) GenerateBorderStrings() {
	top := ""
	bottom := ""
	right := ""
	left := ""
	lenY := len(t.Grid)
	lenX := len(t.Grid[0])
	for i, c := range t.Grid[0] {
		top += c
		bottom += t.Grid[lenY-1][i]
	}
	t.Border["top"] = top
	t.Border["bottom"] = bottom

	for i, c := range t.Grid {
		right += c[lenX-1]
		left += t.Grid[i][0]
	}
	t.Border["right"] = right
	t.Border["left"] = left
}

func (t *Tile) MergeLR(t2 *Tile) *Tile {
	grid := make(Grid, len(t.Grid))
	for y := 0; y < len(t.Grid); y++ {
		grid[y] = append(t.Grid[y], t2.Grid[y]...)
	}
	return NewTile(-1, grid)
}

func (t *Tile) FlipY() {
	lenY := len(t.Grid)
	lenX := len(t.Grid[0])
	grid := make(Grid, lenY)
	y_ := 0

	for y := lenY - 1; y >= 0; y-- {
		row := make([]string, lenX)

		copy(row, t.Grid[y])
		grid[y_] = row
		y_++
	}
	t.Grid = grid
	t.GenerateBorderStrings()
}
func (t *Tile) FlipX() {
	lenY := len(t.Grid)
	lenX := len(t.Grid[0])
	grid := make(Grid, lenY)

	for y := 0; y < lenY; y++ {
		row := make([]string, lenX)
		x_ := 0
		for x := lenX - 1; x >= 0; x-- {
			row[x_] = t.Grid[y][x]
			x_++
		}
		grid[y] = row

	}
	t.Grid = grid
	t.GenerateBorderStrings()
}

func (t *Tile) RotateLeft() {
	length := len(t.Grid)

	grid := make(Grid, length)

	idx := 0
	for x := length - 1; x >= 0; x-- {
		row := make([]string, length)
		x_ := 0
		for y := 0; y < length; y++ {
			row[x_] = t.Grid[y][x]
			x_++
		}
		grid[idx] = row
		idx++

	}
	t.Grid = grid
	t.GenerateBorderStrings()
}

func (t *Tile) RotateRight() {
	length := len(t.Grid)

	grid := make(Grid, length)

	idx := 0
	for x := 0; x < length; x++ {
		row := make([]string, length)
		x_ := 0
		for y := length - 1; y >= 0; y-- {
			row[x_] = t.Grid[y][x]
			x_++
		}
		grid[idx] = row
		idx++

	}
	t.Grid = grid
	t.GenerateBorderStrings()
}

func NewTile(id int, grid Grid) *Tile {
	m := map[string]string{"top": "", "right": "", "botom": "", "left": ""}
	t := Tile{Id: id, Grid: grid, Border: m, Lock: false}
	t.GenerateBorderStrings()

	return &t
}

func PrettyPrintTiles(tiles Tiles) {
	keys := make([]int, len(tiles))

	idx := 0
	for key, _ := range tiles {
		keys[idx] = key
		idx++
	}

	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("Tile %d:\n", tiles[key].Id)
		pretty.Print2dString(tiles[key].Grid)
	}
}
