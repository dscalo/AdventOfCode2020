package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type password struct {
	min int
	max int
	c   string
	p   string
}

func (pwd *password) isValid() bool {
	count := 0
	for _, char := range pwd.p {
		if string(char) == pwd.c {
			count++
		}
	}
	if count >= pwd.min && count <= pwd.max {
		return true
	}
	return false
}

func (pwd *password) isValidPart2() bool {
	// Be careful; Toboggan Corporate Policies have no concept of "index zero"!
	pos1 := pwd.min - 1
	pos2 := pwd.max - 1
	count := 0

	if string(pwd.p[pos1]) == pwd.c {
		count++
	}

	if string(pwd.p[pos2]) == pwd.c {
		count++
	}

	return count == 1
}

func newPassword(min int, max int, c string, p string) *password {
	pwd := password{min: min, max: max, c: c, p: p}
	return &pwd
}

func readFiled2(path string) (int, int) {
	file, err := os.Open(path)
	check(err)

	part1Valid := 0
	part2Valid := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// ["1-2", "a:", "aaasds" ]
		line := strings.Split(scanner.Text(), " ")
		// ["1", "2"]
		minMax := strings.Split(line[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		pwd := newPassword(min, max, string(line[1][0]), line[2])
		if pwd.isValid() {
			part1Valid++
		}

		if pwd.isValidPart2() {
			part2Valid++
		}
	}

	return part1Valid, part2Valid
}

func Day02() {
	// test 1 : 2 valid
	inputs := []string{"test01", "puzzle"}
	for _, f := range inputs {
		path := fmt.Sprintf("input/day02/%s.input", f)
		part1, part2 := readFiled2(path)
		fmt.Printf("%s: part1 valid: %d | part2 valid: %d \n", f, part1, part2)

	}
}
