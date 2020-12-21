package puzzles

import "testing"

func getTestTiles() Tiles {
	tiles := readTiles("../input/day20/test01.input")
	return tiles

}

func testBorders(tile Tile, expect map[string]string, t *testing.T) {
	if tile.Border["top"] != expect["top"] {
		t.Errorf("Top Border: expected %s, recieved: %s\n", expect["top"], tile.Border["top"])
	}

	if tile.Border["right"] != expect["right"] {
		t.Errorf("Right Border: expected %s, recieved: %s\n", expect["right"], tile.Border["right"])
	}

	if tile.Border["bottom"] != expect["bottom"] {
		t.Errorf("Bottom Border: expected %s, recieved: %s\n", expect["bottom"], tile.Border["bottom"])
	}
	if tile.Border["left"] != expect["left"] {
		t.Errorf("Left Border: expected %s, recieved: %s\n", expect["left"], tile.Border["left"])
	}
}

func Test_generate_borders(t *testing.T) {
	tiles := getTestTiles()
	//1951
	expect := map[string]string{
		"top":    "#.##...##.",
		"right":  ".#####..#.",
		"bottom": "#...##.#..",
		"left":   "##.#..#..#",
	}
	testTile := Tile{}

	for _, t := range tiles {
		if t.Id == 1951 {
			testTile = t
			break
		}
	}
	testBorders(testTile, expect, t)

}

func Test_FlipY(t *testing.T) {
	tiles := getTestTiles()
	testTile := Tile{}
	for _, t := range tiles {
		if t.Id == 1951 {
			testTile = t
			break
		}
	}

	expect := map[string]string{
		"top":    "#...##.#..",
		"right":  ".#..#####.",
		"bottom": "#.##...##.",
		"left":   "#..#..#.##",
	}
	testTile.FlipY()

	testBorders(testTile, expect, t)

}

func Test_FlipX(t *testing.T) {
	tiles := getTestTiles()
	testTile := Tile{}
	for _, t := range tiles {
		if t.Id == 1171 {
			testTile = t
			break
		}
	}
	expect := map[string]string{

		"top":    ".##...####",
		"right":  "###....##.",
		"bottom": "...##.....",
		"left":   ".#..#.....",
	}
	testTile.FlipX()
	testBorders(testTile, expect, t)

}

func Test_Rotate_right(t *testing.T) {
	tiles := getTestTiles()
	testTile := Tile{}
	for _, t := range tiles {
		if t.Id == 2311 {
			testTile = t
			break
		}
	}
	expect := map[string]string{

		"top":    ".#..#####.",
		"right":  "..##.#..#.",
		"bottom": "#..##.#...",
		"left":   "..###..###",
	}
	testTile.RotateRight()
	testBorders(testTile, expect, t)
}

func Test_Rotate_left(t *testing.T) {
	tiles := getTestTiles()
	testTile := Tile{}
	for _, t := range tiles {
		if t.Id == 2311 {
			testTile = t
			break
		}
	}
	expect := map[string]string{

		"top":    "...#.##..#",
		"right":  "###..###..",
		"bottom": ".#####..#.",
		"left":   ".#..#.##..",
	}
	testTile.RotateLeft()
	testBorders(testTile, expect, t)
}

func Test_Rotate_right_FlipY(t *testing.T) {
	tiles := getTestTiles()
	testTile := Tile{}
	for _, t := range tiles {
		if t.Id == 2473 {
			testTile = t
			break
		}
	}
	expect := map[string]string{

		"top":    "..#.###...",
		"right":  ".####....#",
		"bottom": ".##...####",
		"left":   ".#.#.###..",
	}
	testTile.RotateRight()
	testTile.FlipY()
	testBorders(testTile, expect, t)
}
