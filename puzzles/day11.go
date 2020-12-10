package puzzles

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	Days[11] = Day11
}

func readFile11(path string) int {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}

	return -1
}

func Day11() {
	inputs := []string{"test01"} // ,"puzzle"
	// part 2-> test01: 8, test02: 19208
	for _, f := range inputs {
		path := fmt.Sprintf("input/day11/%s.input", f)

		//fmt.Println(adapters)
		ansP1 := readFile11(path)
		ansP2 := -1

		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
