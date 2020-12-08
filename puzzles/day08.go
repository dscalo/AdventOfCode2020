package puzzles

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	Days[8] = Day08
}

func readFile08(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
	}

	return -1
}

func Day08() {
	inputs := []string{"test01"} // , "puzzle"
	for _, f := range inputs {
		fmt.Printf("DAY 08 %s under construction!\n", f)
	}
}
