package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Days[16] = Day16
}

type Note struct {
	name      string
	range1Min int
	range1Max int
	range2Min int
	range2Max int
}

func (n *Note) prettyPrint() {
	fmt.Printf(
		"%s: %d-%d or %d-%d\n",
		n.name,
		n.range1Min, n.range1Max, n.range2Min, n.range2Max,
	)
}

func newNote(name string, ranges string) *Note {
	rs := strings.Split(ranges, " or ")

	numbs1 := strings.Split(rs[0], "-")
	mn1, _ := strconv.Atoi(numbs1[0])
	mx1, _ := strconv.Atoi(numbs1[1])

	numbs2 := strings.Split(rs[1], "-")
	mn2, _ := strconv.Atoi(numbs2[0])
	mx2, _ := strconv.Atoi(numbs2[1])

	n := Note{name: name, range1Min: mn1, range1Max: mx1, range2Min: mn2, range2Max: mx2}
	return &n
}

func convertToIntArr(numbStr string) []int {
	ns := strings.Split(numbStr, ",")
	numbs := make([]int, len(ns))

	for idx, val := range ns {
		n, _ := strconv.Atoi(val)
		numbs[idx] = n
	}

	return numbs
}

func findInvalidFields(notes []Note, tickets [][]int) (int, []int) {
	invalidSum := 0
	valid := map[int]int{}
	invalid := map[int]int{}
	var invalidTickets []int

	for t_idx, ticket := range tickets {
		invalidTicket := false
		for _, field := range ticket {
			if _, ok := valid[field]; ok {
				valid[field] += 1
				continue
			}
			if _, ok := invalid[field]; ok {
				invalid[field] += 1
				invalidTicket = true
				continue
			}
			isValid := false
			for _, note := range notes {
				if field >= note.range1Min && field <= note.range1Max {
					isValid = true
					break
				}
				if field >= note.range2Min && field <= note.range2Max {

					isValid = true
					break
				}
			}
			if isValid {
				valid[field] = 1
			} else {
				invalidTicket = true
				invalid[field] = 1
			}
		}
		if invalidTicket {
			invalidTickets = append(invalidTickets, t_idx)
		}
	}

	for k, v := range invalid {
		invalidSum += k * v
	}

	return invalidSum, invalidTickets
}

func containsItem(arr []int, item int) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}
	return false
}

func purgeInvalidTickets(tickets [][]int, invalid []int) [][]int {
	var goodTickets [][]int
	for idx, ticket := range tickets {
		if containsItem(invalid, idx) {
			continue
		}
		goodTickets = append(goodTickets, ticket)
	}
	return goodTickets
}

func remove(notes []Note, key string) []Note {
	var ns []Note

	for _, note := range notes {
		if note.name == key {
			continue
		}
		ns = append(ns, note)
	}
	return ns
}

func getDepartureProduct(ticket []int, fieldNames map[int]string) int {
	product := 1
	for pos, name := range fieldNames {
		if strings.Contains(name, "departure") {
			product *= ticket[pos]
		}
	}
	return product
}

func makeNoteMap(notes []Note) map[string][]int {
	nm := map[string][]int{}

	for _, n := range notes {
		nm[n.name] = []int{}
	}

	return nm
}

func determineFields(tickets [][]int, notes []Note) map[int]string {
	fieldNames := map[int]string{}

	for len(notes) > 0 {
		matches := makeNoteMap(notes)

		for _, note := range notes {
			for i := 0; i < len(tickets[0]); i++ {
				colValid := true
				if _, ok := fieldNames[i]; ok {
					continue
				}
				for j := 0; j < len(tickets); j++ {
					field := tickets[j][i]
					if field >= note.range1Min && field <= note.range1Max {
						continue
					}

					if field >= note.range2Min && field <= note.range2Max {
						continue
					} else {
						colValid = false
						break
					}
				}
				if colValid {
					matches[note.name] = append(matches[note.name], i)
				}
			}
		}
		for k, v := range matches {
			if len(v) == 1 {
				fieldNames[v[0]] = k
				notes = remove(notes, k)
			}
		}

	}
	return fieldNames

}

func readTicketNotes(path string) ([]Note, [][]int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var notes []Note
	var tickets [][]int

	scanner := bufio.NewScanner(file)

	dataType := "n"

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			dataType = "t"
			continue
		}

		switch dataType {
		case "t":
			tickets = append(tickets, convertToIntArr(line))
		case "n":
			ls := strings.Split(line, ": ")
			notes = append(notes, *newNote(ls[0], ls[1]))
		}
	}

	return notes, tickets
}

func Day16() {
	inputs := []string{"test01", "test02", "puzzle"} //

	for _, f := range inputs {
		path := fmt.Sprintf("input/day16/%s.input", f)

		notes, tickets := readTicketNotes(path)
		invalidSum, invalidTickets := findInvalidFields(notes, tickets)
		tickets = purgeInvalidTickets(tickets, invalidTickets)

		names := determineFields(tickets, notes)

		ansP1 := invalidSum
		ansP2 := getDepartureProduct(tickets[0], names)
		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}

}
