package puzzles

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	Days[21] = Day21
}

func readFile21(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
	}

	return -1
}

func Day21() {
	inputs := []string{"test01"} // , "puzzle"
	for _, f := range inputs {
		path := fmt.Sprintf("input/day21/%s.input", f)

		ansP1 := readFile21(path)
		ansP2 := -1

		fmt.Printf("%s: part 1: %d | part 2: %d\n", f, ansP1, ansP2)
	}
}
