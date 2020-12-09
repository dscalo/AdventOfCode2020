package puzzles

import (
	"errors"
	"fmt"
	"sort"
)

func init() {
	Days[9] = Day09
}

func findNonValidNumber(numbs []int, preamble int) int {
	for i := preamble; i < len(numbs); i++ {
		a := numbs[i-preamble : i]
		_, _, err := twoSum(a, numbs[i], -1)

		if err != nil {
			return numbs[i]
		}
	}
	return -1
}

func findContiguousSet(numbs []int, target int) (int, int, error) {
	s := 0
	e := 0
	val := 0
	for e < len(numbs) {
		if val == target {
			return s, e - 1, nil
		}
		if val > target {
			val -= numbs[s]
			s++
			continue
		}
		val += numbs[e]
		e++
	}
	return 0, 0, errors.New("no set found")
}

func Day09() {
	inputs := []string{"test01", "puzzle"} //
	preamble := map[string]int{"test01": 5, "puzzle": 25}
	for _, f := range inputs {
		path := fmt.Sprintf("input/day09/%s.input", f)
		numbs := readIntArray(path)
		ansP1 := findNonValidNumber(numbs, preamble[f])

		s, e, err := findContiguousSet(numbs, ansP1)

		if err != nil {
			panic(err)
		}

		set := numbs[s : e+1]
		sort.Ints(set)
		ansP2 := set[0] + set[len(set)-1]

		fmt.Printf("%s part 1 : %d | part 2: %d \n", f, ansP1, ansP2)
	}

}
