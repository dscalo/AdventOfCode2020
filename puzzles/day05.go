package puzzles

import (
	"bufio"
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/fs"
	"os"
	"sort"
	"strings"
)

func init() {
	Days[5] = Day05
}

func getRange(lower int, upper int, section string) (int, int) {
	l := 0
	u := 0
	switch section {
	case "F":
		fallthrough
	case "L":
		l = lower
		u = (upper + lower) / 2
	case "B":
		fallthrough
	case "R":
		l = ((upper + lower) / 2) + 1
		u = upper
	}
	return l, u
}

func getAirplaneSection(chars string, section string) int {
	upper := 127
	if section == "COL" {
		upper = 7
	}
	lower := 0
	for _, c := range strings.Split(chars, "") {
		lower, upper = getRange(lower, upper, c)
	}

	if upper != lower {
		panic("Bad results")
	}

	return upper
}

func readBoardingPasses(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	size, _ := fs.LineCount(path)
	missingId := 0
	ids := make([]int, size)

	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		rowChars := line[0:7]
		colChars := line[7:]

		row := getAirplaneSection(rowChars, "ROW")
		col := getAirplaneSection(colChars, "COL")
		ids[idx] = row*8 + col

		idx++
	}
	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		if ids[i]+1 != ids[i+1] {
			missingId = ids[i] + 1
			break
		}
	}
	return ids[len(ids)-1], missingId
}

func Day05() {
	inputs := []string{"test01", "puzzle"} //

	/*
		FBFBBFFRLR: row 44, column 5. seat ID 357
		BFFFBBFRRR: row 70, column 7, seat ID 567.
		FFFBBBFRRR: row 14, column 7, seat ID 119.
		BBFFBBFRLL: row 102, column 4, seat ID 820.

	*/
	for _, f := range inputs {
		if f == "test01" {
			continue
		}
		path := fmt.Sprintf("input/day05/%s.input", f)
		highest, missing := readBoardingPasses(path)
		fmt.Printf("%s: highest ID: %d missing ID: %d\n", f, highest, missing)
	}
}
