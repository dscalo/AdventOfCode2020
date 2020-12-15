package puzzles

import (
	"fmt"
)

func init() {
	Days[15] = Day15
}

func makeGameMap(numbs []int) map[int][]int {
	spoken := map[int][]int{}
	for i, n := range numbs {
		spoken[n] = []int{i}
	}
	return spoken
}

func addSpoken(num int, turn int, spoken map[int][]int) {
	if val, ok := spoken[num]; ok {
		spoken[num] = append([]int{turn}, val[0])
	} else {
		spoken[num] = []int{turn}
	}
}

func playGame(numbs []int, finish int) int {
	spoken := makeGameMap(numbs)
	turn := len(spoken)
	lastSpoken := numbs[len(numbs)-1]
	for t := turn; t < finish; t++ {
		if val, ok := spoken[lastSpoken]; ok {
			if len(val) == 1 {
				spoken[lastSpoken] = append(val, t)
				lastSpoken = 0
				addSpoken(0, t, spoken)
			} else {
				v := val[0] - val[1]
				lastSpoken = v
				addSpoken(lastSpoken, t, spoken)
			}
		} else {
			addSpoken(lastSpoken, t, spoken)
			lastSpoken = 0
		}
	}
	return lastSpoken
}

func Day15() {
	inputs := []string{"test01", "puzzle"} //

	for _, f := range inputs {
		path := fmt.Sprintf("input/day15/%s.input", f)
		numbs := readIntArray(path)

		ansP1 := playGame(numbs, 2020)
		ansP2 := playGame(numbs, 30000000)

		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
