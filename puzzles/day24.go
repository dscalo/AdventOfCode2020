package puzzles

import (
	"bufio"
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/fs"
	"os"
	"strings"
)

func init() {
	Days[24] = Day24
}

type Directions [][]string

type Hex struct {
	Q int
	R int
}

func (h *Hex) PrettyPrint() {
	fmt.Printf("Q: %d, R: %d\n", h.Q, h.R)
}

func NewHex(q, r int) *Hex {
	h := Hex{Q: q, R: r}
	return &h
}

type Hexes map[Hex]bool

func prettyPrintHexes(hexes Hexes) {
	color := "white"
	for hex, isBlack := range hexes {
		if isBlack {
			color = "black"
		}
		fmt.Printf("Q: %d, R: %d %s\n", hex.Q, hex.R, color)
	}
}

func prettyPrintDirections(dirs Directions) {
	fmt.Println("---- Directions ----")
	for _, dir := range dirs {
		for _, d := range dir {
			fmt.Printf("%s->", d)
		}
		fmt.Println()
	}
}

func readTileDirections(path string) Directions {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	size, err := fs.LineCount(path)
	if err != nil {
		panic(err)
	}

	dirs := make(Directions, size)

	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		var dir []string
		for i := 0; i < len(line); i++ {
			c := line[i]
			switch c {
			case "e":
				fallthrough
			case "w":
				dir = append(dir, c)
			case "s":
				fallthrough
			case "n":
				dir = append(dir, c+line[i+1])
				i++
			}
		}
		dirs[idx] = dir
		idx++
	}

	return dirs
}

func followDirection(dirs []string) *Hex {
	hex := NewHex(0, 0)
	for _, dir := range dirs {
		switch dir {
		case "e":
			hex.Q += 1
		case "se":
			hex.R += 1
		case "sw":
			hex.Q += -1
			hex.R += 1
		case "w":
			hex.Q += -1
		case "nw":
			hex.R += -1
		case "ne":
			hex.Q += 1
			hex.R += -1
		}
	}
	return hex
}

func flipTiles(dirs Directions) *Hexes {
	hexes := Hexes{}

	for _, dir := range dirs {
		hex := followDirection(dir)
		//hex.PrettyPrint()
		if val, ok := hexes[*hex]; ok {

			hexes[*hex] = !val
		} else {
			hexes[*hex] = true
		}
	}

	return &hexes
}

func countBlackHexes(hexes *Hexes) int {
	count := 0

	for _, v := range *hexes {
		if v {
			count++
		}
	}
	return count
}

func getHexAdjacencyList(hex *Hex) []Hex {
	hexes := make([]Hex, 6)

	// ne
	hexes[0] = *NewHex(hex.Q+1, hex.R-1)
	// e
	hexes[1] = *NewHex(hex.Q+1, hex.R)
	// se
	hexes[2] = *NewHex(hex.Q, hex.R+1)
	// sw
	hexes[3] = *NewHex(hex.Q-1, hex.R+1)
	// w
	hexes[4] = *NewHex(hex.Q-1, hex.R)
	// nw
	hexes[5] = *NewHex(hex.Q, hex.R-1)

	return hexes
}

func artFlip(hexes *Hexes, days int) {
	for hx, _ := range *hexes {
		adj := getHexAdjacencyList(&hx)
		for _, h := range adj {
			if _, ok := (*hexes)[h]; !ok {
				(*hexes)[h] = false
			}
		}
	}
	for i := 0; i < days; i++ {
		var flipToBlack []Hex
		var flipToWhite []Hex

		for h, _ := range *hexes {
			adj := getHexAdjacencyList(&h)

			bct := 0
			for _, h1 := range adj {
				if isBlack, ok := (*hexes)[h1]; ok {
					if isBlack {
						bct++
					}
				} else {
					(*hexes)[h1] = false
				}
			}

			if (*hexes)[h] && (bct == 0 || bct > 2) {
				flipToWhite = append(flipToWhite, h)
			}

			if !(*hexes)[h] && bct == 2 {
				flipToBlack = append(flipToBlack, h)
			}
		}

		for _, wh := range flipToWhite {
			(*hexes)[wh] = false
		}

		for _, bh := range flipToBlack {
			(*hexes)[bh] = true
		}
	}
}

func Day24() {
	inputs := []string{"test01", "puzzle"} //

	for _, f := range inputs {
		path := fmt.Sprintf("input/day24/%s.input", f)

		dirs := readTileDirections(path)
		hexes := flipTiles(dirs)

		//prettyPrintDirections(dirs)
		//prettyPrintHexes(*hexes)
		ansP1 := countBlackHexes(hexes)
		artFlip(hexes, 100)

		ansP2 := countBlackHexes(hexes)

		fmt.Printf("%s part 1 : %d | part 2: %d \n", f, ansP1, ansP2)
	}
}
