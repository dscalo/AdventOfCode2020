package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Days[14] = Day14
}

func revArr(arr []string) []string {
	rev := make([]string, len(arr))

	idx := 0

	for i := len(arr) - 1; i >= 0; i-- {
		rev[idx] = arr[i]
		idx++
	}

	return rev
}

func createMask(maskString string) map[int]int {
	mask := map[int]int{}
	arr := revArr(strings.Split(maskString, ""))

	for idx, val := range arr {
		switch val {
		case "1":
			mask[idx] = 1
		case "0":
			mask[idx] = 0
		case "X":
			mask[idx] = -1
		}

	}
	return mask
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func clearBit(n int, pos uint) int {
	m := ^(1 << pos)
	n &= m
	return n
}

func applyMask(num int, mask map[int]int, clearZeros bool) int {
	for k, v := range mask {
		switch v {
		case 0:
			if clearZeros {
				num = clearBit(num, uint(k))
			}
		case 1:
			num = setBit(num, uint(k))
		}
	}
	return num
}

func sumMap(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}

	return sum
}

func applyValToMems(addr int, val int, mem map[int]int, mask map[int]int) {
	addresses := map[int]int{applyMask(addr, mask, false): val}

	for pos, bit := range mask {
		if bit == -1 {
			var toAdd []int
			for address, _ := range addresses {
				toAdd = append(toAdd, setBit(address, uint(pos)))
				toAdd = append(toAdd, clearBit(address, uint(pos)))
			}

			for _, ta := range toAdd {
				addresses[ta] = val
			}
		}
	}

	for addr, _ := range addresses {
		mem[addr] = val
	}
}

func getMemAddress(s string) int {
	s = strings.ReplaceAll(s, "mem[", "")
	s = strings.ReplaceAll(s, "]", "")
	n, _ := strconv.Atoi(s)
	return n
}

func readProgram(path string, part int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	mem := map[int]int{}

	var mask map[int]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		val := strings.TrimSpace(line[1])
		if strings.TrimSpace(line[0]) == "mask" {
			mask = createMask(val)
		} else {
			n, _ := strconv.Atoi(val)
			key := getMemAddress(strings.TrimSpace(line[0]))
			if part == 1 {
				mem[key] = applyMask(n, mask, true)
			} else {
				applyValToMems(key, n, mem, mask)
			}

		}
	}
	//fmt.Println(mem)
	return sumMap(mem)
}

func Day14() {
	inputs := []string{"test01", "test02", "puzzle"} // }

	for _, f := range inputs {
		path := fmt.Sprintf("input/day14/%s.input", f)

		ansP1 := readProgram(path, 1)
		ansP2 := -1
		if f != "test01" {
			ansP2 = readProgram(path, 2)
		}
		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
