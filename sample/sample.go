package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func lineCount(path string) (error, int) {
	file, err := os.Open("sample/test01.txt")
	check(err)

	lines := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines++
	}

	err = file.Close()

	return err, lines

}

func getSize(line []string) (int, int) {
	x, _ := strconv.Atoi(line[0])
	y, _ := strconv.Atoi(line[1])
	return x, y
}

func createMatrix(scanner *bufio.Scanner, x int, y int) [][]int {
	source := make([][]int, x)

	for i := 0; i < x; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		values := make([]int, y)
		for idx, val := range line {
			val, _ := strconv.Atoi(val)
			values[idx] = val
		}
		source[i] = values
	}

	return source
}

func getOutput(path string) int {
	file, err := os.Open(path)
	check(err)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	output := scanner.Text()

	num, _ := strconv.Atoi(output)

	return num

}

func readFile(path string) ([][]int, [][]int) {
	//err, size  := lineCount(path)
	//check(err)

	file, err := os.Open(path)
	check(err)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	sourceLine := strings.Split(scanner.Text(), " ")
	sx, sy := getSize(sourceLine)
	source := createMatrix(scanner, sx, sy)

	scanner.Scan()
	patternLine := strings.Split(scanner.Text(), " ")
	px, py := getSize(patternLine)
	pattern := createMatrix(scanner, px, py)

	return source, pattern
}

func getMatrixSize(M [][]int) (int, int) {
	return len(M), len(M[0])
}

func findPatterns(source [][]int, pattern [][]int) int {
	matches := 0
	px, py := getMatrixSize(pattern)

	for x := 0; x <= len(source)-px; x++ {
		for y := 0; y <= len(source[x])-py; y++ {

			// searching for match
			match := 0
			for j := 0; j < len(pattern); j++ {
				for k := 0; k < len(pattern[j]); k++ {
					val := source[x+j][y+k]
					if val != pattern[j][k] {
						match = -1
						break
					}
				}
				if match == -1 {
					break
				}
				match += 1
			}
			if match == px {
				matches += 1
			}
		}
	}

	return matches
}

func main() {

	for i := 0; i <= 10; i++ {
		input := fmt.Sprintf("sample/IO/test%d.input", i)
		output := fmt.Sprintf("sample/IO/test%d.output", i)

		source, pattern := readFile(input)
		answer := getOutput(output)
		matches := findPatterns(source, pattern)

		fmt.Printf("Test %d expected: %d actual: %d \n", i, answer, matches)
	}
}
