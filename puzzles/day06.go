package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func init() {
	Days[6] = Day06
}

func getCommonAnswers(answers map[string]int, groupSize int) int {
	ct := 0
	for _, v := range answers {
		if v == groupSize {
			ct++
		}
	}
	return ct
}

func readAnswers(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	answers := make(map[string]int)

	scanner := bufio.NewScanner(file)

	ansCount := 0
	commonAnsCT := 0
	groupSize := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			ansCount += len(answers)
			commonAnsCT += getCommonAnswers(answers, groupSize)
			answers = make(map[string]int)
			groupSize = 0
			continue
		}
		groupSize++
		ans := strings.Split(line, "")

		for _, a := range ans {
			if val, ok := answers[a]; ok {
				answers[a] = val + 1
			} else {
				answers[a] = 1
			}
		}
	}
	ansCount += len(answers)
	commonAnsCT += getCommonAnswers(answers, groupSize)
	return ansCount, commonAnsCT
}

func Day06() {
	inputs := []string{"test01", "puzzle"} //

	for _, f := range inputs {
		path := fmt.Sprintf("input/day06/%s.input", f)
		ansP1, ansP2 := readAnswers(path)

		fmt.Printf("%s: part1: %d | part2: %d\n", f, ansP1, ansP2)
	}
}
