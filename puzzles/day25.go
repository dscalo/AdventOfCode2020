package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func init() {
	Days[25] = Day25
}

func readKeys(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	keys := make([]int, 2)
	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		keys[idx] = n
		idx++
	}

	return keys
}

func transform(subjectNumber, loopSize int) int {
	value := 1

	for i := 0; i < loopSize; i++ {
		value *= subjectNumber
		value = value % 20201227
		//fmt.Printf("value: %d\n", value)
	}

	return value
}

func getLoopSize(key int) int {
	subjectNumber := 7
	value := 1
	idx := 0

	for {
		//if idx > 100000000 {
		//	fmt.Println("SOMETHING IS WRONG")
		//	return -1
		//}
		value *= subjectNumber
		value = value % 20201227
		//fmt.Printf(" key : %d value = %d\n",key, value)
		idx++

		if value == key {
			return idx
		}
	}

	return -1
}

func Day25() {
	inputs := []string{"test01", "puzzle"} //

	for _, f := range inputs {
		path := fmt.Sprintf("input/day25/%s.input", f)

		keys := readKeys(path)
		fmt.Println(keys)

		lp1 := getLoopSize(keys[0])
		//lp2 := getLoopSize(keys[1])

		ansP1 := transform(keys[1], lp1)
		ansP2 := -1

		fmt.Printf("%s part 1 : %d | part 2: %d \n", f, ansP1, ansP2)
	}
}
