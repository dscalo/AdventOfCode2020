package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func lineCount(path string) (error, int) {
	file, err := os.Open(path)
	check(err)

	lines := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines++
	}

	err = file.Close()

	return err, lines

}

func readFile(path string) []int {
	err, size := lineCount(path)
	check(err)

	file, err := os.Open(path)
	check(err)

	numbs := make([]int, size)

	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbs[idx] = n
		idx += 1
	}

	return numbs
}

func twoSum(numbs []int, target int, skip int) (int, int, error) {
	m := make(map[int]int)
	for i := 0; i < len(numbs); i++ {
		if i == skip {
			continue
		}
		if _, ok := m[target-numbs[i]]; ok {
			return target - numbs[i], numbs[i], nil
		}
		m[numbs[i]] = numbs[i]
	}

	return 0, 0, errors.New("no matches found")
}

func threeSum(numbs []int, target int) (int, int, int) {
	for idx, n := range numbs {
		x, y, err := twoSum(numbs, target-n, idx)
		if err == nil {
			return n, x, y
		}
	}

	return 0, 0, 0
}

func main() {
	// test01 == 514579
	part1 := []string{"test01", "puzzle"}

	for _, f := range part1 {
		path := fmt.Sprintf("day01/%s.input", f)
		numbs := readFile(path)
		x, y, err := twoSum(numbs, 2020, -1)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s: x: %d * y: %d == %d\n", f, x, y, x*y)

		a, b, c := threeSum(numbs, 2020)
		fmt.Printf("(PART 2) %s: a: %d * b: %d * c: %d == %d\n", f, a, b, c, (a * b * c))
	}

}
