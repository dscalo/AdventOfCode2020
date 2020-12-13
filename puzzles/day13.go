package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Days[13] = Day13
}

func readtimestamps(path string) (int, []int, map[int]int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	ts, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	ids := strings.Split(scanner.Text(), ",")

	var times []int
	offsets := map[int]int{}

	for idx, id := range ids {
		if id == "x" {
			continue
		}
		n, _ := strconv.Atoi(id)
		offsets[n] = idx
		times = append(times, n)
	}
	return ts, times, offsets
}

func getEarliestTimeStamp(ts int, times []int) int {
	lowest := 999999999
	timeId := -1
	for _, time := range times {
		n := ts / time
		n = (n + 1) * time
		if n < lowest {
			lowest = n
			timeId = time
		}
	}
	return (lowest - ts) * timeId
}

func getEarliestConsecutiveTS(times []int, offsets map[int]int) int {
	toMatch := times[0]
	multiple := times[0]
	for idx := 1; idx < len(times); idx++ {
		match := false
		for !match {
			toMatch += multiple
			for x := idx; x > 0; x-- {
				t := times[x]
				if (toMatch+offsets[t])%t != 0 {
					match = false
					break
				}
				match = true
			}
		}
		if match && idx < len(times)-1 {
			multiple = times[0]
			for m := idx; m > 0; m-- {
				multiple *= times[m]
			}
		}
	}
	return toMatch
}

func Day13() {
	inputs := []string{"test01", "test02A", "test02", "test03", "test04", "test05", "puzzle"} //
	//
	for _, f := range inputs {
		path := fmt.Sprintf("input/day13/%s.input", f)

		ts, times, offsets := readtimestamps(path)

		ansP1 := -1
		if ts > 0 {
			ansP1 = getEarliestTimeStamp(ts, times)
		}
		ansP2 := getEarliestConsecutiveTS(times, offsets)
		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
