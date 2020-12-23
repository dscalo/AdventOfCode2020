package puzzles

import (
	"bufio"
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/lists"
	"os"
	"strconv"
	"strings"
)

func init() {
	Days[23] = Day23
}

func indexOf(target int, arr []int) int {

	for idx, item := range arr {
		if item == target {
			return idx
		}
	}

	return -1
}

func readCups(path string) *lists.LinkedList {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	cups := lists.NewLinkedList()

	sArr := strings.Split(line, "")

	for _, s := range sArr {
		n, _ := strconv.Atoi(s)
		cups.InsertEnd(n)
	}

	return cups
}

func playCrapCups(cups *lists.LinkedList, moves int, max int) {
	currentCup := cups.Start

	for i := 0; i < moves; i++ {
		//if i % 1000 == 0 {
		//	fmt.Println(i)
		//}
		//fmt.Printf("---- MOVE %d ----\n", i+1)
		//cups.PrettyPrint()

		toDel1 := currentCup.Next.Value
		toDel2 := currentCup.Next.Next.Value
		toDel3 := currentCup.Next.Next.Next.Value

		//fmt.Printf("deleting %d\n", toDel1)
		cups.DeleteAfter(currentCup, 2)

		//cups.PrettyPrint()
		toFind := currentCup.Value - 1
		dest := false

		for !dest {
			if toFind < 1 {
				toFind = max
			}
			if toFind == toDel1 || toFind == toDel2 || toFind == toDel3 {
				toFind--
			} else {
				dest = true
				break
			}

		}

		destNode := cups.Find(toFind)

		if destNode == nil {
			panic("this should not happen!")
		}

		cups.InsertAfter(destNode, toDel1, toDel2, toDel3)

		currentCup = currentCup.Next
		//fmt.Printf("current cup val : %d\n", currentCup.Value)

		// cups.PrettyPrint()

	}

}

func rearrange(start int, arr []int) string {
	startIdx := indexOf(start, arr) + 1
	str := ""

	for i := 0; i < len(arr)-1; i++ {
		idx := startIdx + i

		if idx >= len(arr) {
			idx = idx - len(arr)
		}
		str = fmt.Sprintf("%s%d", str, arr[idx])
	}

	return str
}

func padTo(cups *lists.LinkedList, max, size int) {
	i := max
	for ; i <= size; i++ {
		cups.InsertEnd(i)
	}
}

func Day23() {
	inputs := []string{"test01", "puzzle"} //  , "puzzle"
	for _, f := range inputs {
		path := fmt.Sprintf("input/day23/%s.input", f)

		cups := readCups(path)
		//cups.PrettyPrint()
		//fmt.Println("cups", cups)

		//playCrapCups(cups, 100,  9)
		//fmt.Println("AFTER 100 Moves", after100Moves)
		ansP1 := ""
		//ansP1 = rearrange(1, cups.ToArray())

		//cups := readCups(path)
		padTo(cups, 10, 1000000)
		playCrapCups(cups, 10000000, 1000000)

		fmt.Printf("FINISHED PLAYING\n")

		one := cups.Find(1)

		ansP2 := one.Next.Value * one.Next.Next.Value
		fmt.Printf("%s: part 1: %s | part 2: %d\n", f, ansP1, ansP2)

	}
}
