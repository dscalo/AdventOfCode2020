package puzzles

import (
	"fmt"
	"sort"
)

func init() {
	Days[10] = Day10
}

func chainAdapters(adapters []int) int {
	diffs := map[int]int{0: 0, 1: 0, 2: 0, 3: 0}
	// 1-jolt differences * 3-jolt differences?

	diff := adapters[0]

	if diff > 3 {
		fmt.Printf("first adapter > 3 %d\n", adapters[0])
	} else {
		diffs[diff] += 1
	}

	for i := 0; i < len(adapters)-1; i++ {
		diff = adapters[i+1] - adapters[i]
		if diff > 3 {
			fmt.Printf("idx %d:%d - idx %d:%d > 3\n", i+1, adapters[i+1], i, adapters[i])
		} else {
			diffs[diff] += 1
		}
	}
	return diffs[1] * diffs[3]
}

func makeMap(arr []int) map[int]int {
	m := map[int]int{}

	for _, v := range arr {
		m[v] = 0
	}
	m[0] = 1
	return m
}

func arrangements(adapters []int) int {
	ln := len(adapters)
	paths := makeMap(adapters)

	for i := 0; i < ln-1; i++ {
		for j := 1; j <= 3; j++ {
			if j+i >= ln {
				break
			}
			if adapters[i+j]-adapters[i] <= 3 {
				paths[adapters[i+j]] += paths[adapters[i]]
			}
		}
	}

	return paths[adapters[len(adapters)-1]]
}

func Day10() {
	inputs := []string{"test01", "test02", "puzzle"} //
	// part 2-> test01: 8, test02: 19208
	for _, f := range inputs {
		path := fmt.Sprintf("input/day10/%s.input", f)

		adapters := readIntArray(path)
		sort.Ints(adapters)

		adapters = append([]int{0}, adapters...)
		adapters = append(adapters, adapters[len(adapters)-1]+3)

		//fmt.Println(adapters)
		ansP1 := chainAdapters(adapters)
		ansP2 := arrangements(adapters)

		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}

}
