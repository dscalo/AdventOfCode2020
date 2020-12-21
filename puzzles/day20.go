package puzzles

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func init() {
	Days[20] = Day20
}

type Puzzle = [][]int

func readTiles(path string) Tiles {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var tiles Tiles
	id := 0
	grid := make(Grid, 10)
	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "Tile") {
			// Tile 1783:
			if id > 0 {
				tiles = append(tiles, *NewTile(id, grid))
				grid = make(Grid, 10)
				idx = 0
			}
			t := strings.ReplaceAll(line, "Tile ", "")
			t = strings.ReplaceAll(t, ":", "")
			n, _ := strconv.Atoi(t)
			id = n
			continue
		}

		grid[idx] = strings.Split(line, "")
		idx++
	}
	tiles = append(tiles, *NewTile(id, grid))
	return tiles
}

func findCorners(tiles Tiles) []int {
	var corners []int

	for t := 0; t < len(tiles); t++ {
		cur := tiles[t]
		matches := 0
		for j := 0; j < len(tiles); j++ {
			if j == t {
				continue
			}
			if findMatch2(&tiles[j], cur.Border["top"]) {
				matches++

			}
			if findMatch2(&tiles[j], cur.Border["right"]) {
				matches++

			}
			if findMatch2(&tiles[j], cur.Border["bottom"]) {
				matches++

			}
			if matches > 2 {
				break
			}
			if findMatch2(&tiles[j], cur.Border["left"]) {
				matches++

			}

		}
		if matches == 2 {
			corners = append(corners, cur.Id)
		}
	}

	return corners
}

func isTileMatch(tile *Tile, target string) bool {
	for _, border := range Borders {
		if tile.Border[border] == target {
			return true
		}
	}

	return false
}

func findMatch2(tile *Tile, toMatch string) bool {
	if isTileMatch(tile, toMatch) {
		return true
	}

	tile.FlipX()
	if isTileMatch(tile, toMatch) {
		return true
	}

	tile.FlipY()
	if isTileMatch(tile, toMatch) {
		return true
	}

	tile.FlipX()
	if isTileMatch(tile, toMatch) {
		return true
	}

	tile.RotateRight()
	if isTileMatch(tile, toMatch) {
		return true
	}

	tile.FlipX()
	if isTileMatch(tile, toMatch) {
		return true
	}

	tile.FlipY()
	if isTileMatch(tile, toMatch) {
		return true
	}

	tile.FlipX()
	if isTileMatch(tile, toMatch) {
		return true
	}

	return false
}

func isTopLeftUnMatched(tiles Tiles, list []int, id int) bool {
	indexOfTile := getIndexOfId(tiles, id)
	for _, l := range list {
		if findMatch2(&tiles[l], tiles[indexOfTile].Border["top"]) {
			return false
		}

	}
	for _, l := range list {
		if findMatch2(&tiles[l], tiles[indexOfTile].Border["left"]) {
			return false
		}

	}

	return true
}

func getIndexOfId(tiles Tiles, id int) int {
	for idx, t := range tiles {
		if t.Id == id {
			return idx
		}
	}

	return -1
}

func findTopLeftCorner(tiles Tiles, corners []int) int {
	var list []int

	for idx, t := range tiles {
		if t.Id == corners[0] || t.Id == corners[1] || t.Id == corners[2] || t.Id == corners[3] {
			continue
		}
		list = append(list, idx)
	}

	for idx, c := range corners {
		indexOfCorner := getIndexOfId(tiles, c)
		if idx == 3 {
			return corners[3]
		}

		if isTopLeftUnMatched(tiles, list, c) {
			return c
		}

		tiles[indexOfCorner].FlipX()
		if isTopLeftUnMatched(tiles, list, c) {
			return c
		}

		tiles[indexOfCorner].FlipY()
		if isTopLeftUnMatched(tiles, list, c) {
			return c
		}

		tiles[indexOfCorner].FlipX()
		if isTopLeftUnMatched(tiles, list, c) {
			return c
		}

		tiles[indexOfCorner].RotateRight()
		if isTopLeftUnMatched(tiles, list, c) {
			return c
		}

		tiles[indexOfCorner].FlipX()
		if isTopLeftUnMatched(tiles, list, c) {
			return c
		}

		tiles[indexOfCorner].FlipY()
		if isTopLeftUnMatched(tiles, list, c) {
			return c
		}

		tiles[indexOfCorner].FlipX()
		if isTopLeftUnMatched(tiles, list, c) {
			return c
		}
	}

	return -1
}

func removeFromIntArr(arr []int, item int) []int {
	newArr := make([]int, len(arr)-1)

	idx := 0
	for _, i := range arr {
		if i == item {
			continue
		}

		newArr[idx] = i
		idx++
	}

	return newArr
}

func findMatch3(tiles Tiles, m1, m2 int, b1, b2 string) bool {
	if tiles[m1].Border[b1] == tiles[m2].Border[b2] {
		return true
	}

	tiles[m2].FlipX()
	if tiles[m1].Border[b1] == tiles[m2].Border[b2] {
		return true
	}

	tiles[m2].FlipY()
	if tiles[m1].Border[b1] == tiles[m2].Border[b2] {
		return true
	}

	tiles[m2].FlipX()
	if tiles[m1].Border[b1] == tiles[m2].Border[b2] {
		return true
	}

	tiles[m2].RotateRight()
	if tiles[m1].Border[b1] == tiles[m2].Border[b2] {
		return true
	}

	tiles[m2].FlipX()
	if tiles[m1].Border[b1] == tiles[m2].Border[b2] {
		return true
	}

	tiles[m2].FlipY()
	if tiles[m1].Border[b1] == tiles[m2].Border[b2] {
		return true
	}

	tiles[m2].FlipX()
	if tiles[m1].Border[b1] == tiles[m2].Border[b2] {
		return true
	}

	return false
}

func assemblePuzzle(tiles Tiles, topLeftId int) Puzzle {
	puzzleGrid := int(math.Sqrt(float64(len(tiles))))
	puzzle := make(Puzzle, puzzleGrid)

	puzzle[0] = []int{topLeftId}

	var list []int

	for _, t := range tiles {
		if t.Id == topLeftId {
			continue
		}

		list = append(list, t.Id)
	}
	X := 0
	for len(list) > 0 {

		for y := 0; y < puzzleGrid-1; y++ {
			toMatch := getIndexOfId(tiles, puzzle[y][X])
			matchedListIndex := -1
			for l := 0; l < len(list); l++ {
				if findMatch3(tiles, toMatch, getIndexOfId(tiles, list[l]), "bottom", "top") {
					matchedListIndex = list[l]
					break
				}

			}
			puzzle[y+1] = append(puzzle[y+1], tiles[getIndexOfId(tiles, matchedListIndex)].Id)

			list = removeFromIntArr(list, matchedListIndex)
			if len(list) == 0 {
				return puzzle
			}
		}

		toMatch := getIndexOfId(tiles, puzzle[0][X])
		X++
		matchX := -1
		for l := 0; l < len(list); l++ {
			if findMatch3(tiles, toMatch, getIndexOfId(tiles, list[l]), "right", "left") {
				matchX = list[l]
				break
			}

		}
		puzzle[0] = append(puzzle[0], tiles[getIndexOfId(tiles, matchX)].Id)

		list = removeFromIntArr(list, matchX)

	}

	return puzzle
}

//
func formMegaTile(tiles Tiles, puzzle Puzzle) *Tile {
	var grid Grid
	for z := 0; z < len(puzzle); z++ {
		var rowTiles Tiles

		for x := 0; x < len(puzzle[z]); x++ {
			rowTiles = append(rowTiles, tiles[getIndexOfId(tiles, puzzle[z][x])])
		}

		for i := 1; i < 9; i++ {
			var row []string
			for _, t := range rowTiles {
				row = append(row, t.Grid[i][1:len(t.Grid[i])-1]...)
			}
			grid = append(grid, row)

		}

	}

	return NewTile(-1, grid)
}

// row 1: 20
// row 2: 2,7,8, 13, 14,19,20, 21
// row 3: 1
func readSeaMonster() Grid {
	file, err := os.Open("input/day20/seaMonster.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var grid Grid

	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	return grid
}

func findSeaMonsters(tile Tile) int {
	total := 0
	seaMonster := readSeaMonster()
	indexes := make([][]int, len(seaMonster))

	for sy := 0; sy < len(seaMonster); sy++ {
		var row []int
		for sx := 0; sx < len(seaMonster[sy]); sx++ {
			if seaMonster[sy][sx] == "#" {
				row = append(row, sx)
			}
		}
		indexes[sy] = row
	}

	//pretty.Print2dInt(indexes)

	//pretty.Print2dString(seaMonster)
	for y := 0; y < len(tile.Grid)-len(indexes); y++ {
		rowLen := len(tile.Grid)
		ii := indexes[0][0]
		for x := 0; x < rowLen; x++ {
			if x+ii >= rowLen {
				break
			}
			if tile.Grid[y][x+ii] == "#" {
				rowMatch := true
				for _, ix := range indexes[1] {
					if x+ix >= rowLen {
						break
					}
					if tile.Grid[y+1][x+ix] != "#" {
						rowMatch = false
						break
					}
				}

				if rowMatch {
					for _, iix := range indexes[2] {
						if x+iix >= rowLen {
							break
						}
						if tile.Grid[y+2][x+iix] != "#" {
							rowMatch = false
							break
						}
					}

					if rowMatch {
						total++
					}
				}
			}
		}
	}

	return total
}

func searchGridForMonster(tile Tile) int {
	total := findSeaMonsters(tile)
	if total != 0 {
		return total
	}
	tile.FlipX()
	total = findSeaMonsters(tile)
	if total != 0 {
		return total
	}

	tile.FlipY()
	total = findSeaMonsters(tile)
	if total != 0 {
		return total
	}

	tile.FlipX()
	total = findSeaMonsters(tile)
	if total != 0 {
		return total
	}

	tile.RotateRight()
	total = findSeaMonsters(tile)
	if total != 0 {
		return total
	}

	tile.FlipX()
	total = findSeaMonsters(tile)
	if total != 0 {
		return total
	}

	tile.FlipY()
	total = findSeaMonsters(tile)
	if total != 0 {
		return total
	}

	tile.FlipX()
	total = findSeaMonsters(tile)
	if total != 0 {
		return total
	}

	return -1
}

func countTileGrid(grid Grid) int {
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

func Day20() {
	inputs := []string{"test01", "puzzle"} // , "test02", "puzzle"

	for _, f := range inputs {
		path := fmt.Sprintf("input/day20/%s.input", f)

		tiles := readTiles(path)
		//viewTile(&tiles[3])
		//break
		fmt.Printf("totalTiles: %d\n", len(tiles))
		corners := findCorners(tiles)
		topLeft := findTopLeftCorner(tiles, corners)
		puzzle := assemblePuzzle(tiles, topLeft)
		//fmt.Println("PRINTING PUZZLE")
		//pretty.Print2dInt(puzzle)

		//fmt.Println("CORNERS", corners)
		//fmt.Println("TOP LEFT", topLeft)
		//idx := getIndexOfId(tiles, topLeft)
		//
		//pretty.Print2dString(tiles[idx].Grid)
		//PrettyPrintTiles(tiles)
		//puzzle := arrangeTiles(tiles)
		//fmt.Println("MEGA TIlE")
		megaTile := formMegaTile(tiles, puzzle)

		//pretty.Print2dString(megaTile.Grid)
		totalMonsters := searchGridForMonster(*megaTile)
		// 15
		totalHashMarks := countTileGrid(megaTile.Grid)

		fmt.Printf("Total monsters: %d\n", totalMonsters)

		//pretty.Print2dString(megaTile.Grid)
		//println(puzzle)
		ansP1 := corners[0] * corners[1] * corners[2] * corners[3]
		// the 15 is the amount of spaces in the sea monster
		// i'm too tired to write a function to count it
		ansP2 := totalHashMarks - (totalMonsters * 15)
		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
