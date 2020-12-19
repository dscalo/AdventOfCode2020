package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Days[19] = Day19
}

type Rule struct {
	sub1 []int
	sub2 []int
	char string
}

type Rules = map[int]*Rule
type Messages = [][]string

func getIntsFromString(strArr []string) []int {
	arr := make([]int, len(strArr))

	for idx, s := range strArr {
		n, err := strconv.Atoi(s)

		if err == nil {
			arr[idx] = n
		}
	}

	return arr
}

func NewRule(rStr string) *Rule {
	rStr = strings.ReplaceAll(rStr, `"`, "")
	rStr = strings.TrimSpace(rStr)
	r := Rule{sub1: make([]int, 0), sub2: make([]int, 0), char: ""}

	if len(rStr) == 1 {
		r.char = rStr
		return &r
	}

	subs := strings.Split(rStr, " | ")

	sArr := strings.Split(strings.TrimSpace(subs[0]), " ")

	r.sub1 = getIntsFromString(sArr)

	if len(subs) == 2 {
		sArr = strings.Split(strings.TrimSpace(subs[1]), " ")
		r.sub2 = getIntsFromString(sArr)
	}
	return &r
}

func PrettyPrintRules(rules Rules) {
	keys := make([]int, len(rules))
	idx := 0
	for n, _ := range rules {
		keys[idx] = n
		idx++
	}

	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%d: ", key)
		rule := rules[key]

		if rule.char != "" {
			fmt.Printf(" \"%s\"\n", rule.char)
			continue
		}

		for _, r1 := range rule.sub1 {
			fmt.Printf(" %d ", r1)
		}

		if len(rule.sub2) > 0 {
			fmt.Printf(" | ")
			for _, r2 := range rule.sub2 {
				fmt.Printf(" %d ", r2)
			}
		}
		fmt.Println("")
	}
}

func readRules(path string) (Rules, Messages) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	rules := Rules{}
	var messages = Messages{}
	scanner := bufio.NewScanner(file)
	tpe := "rules"
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			tpe = "messages"
			continue
		}

		switch tpe {
		case "rules":
			rule := strings.Split(line, ": ")
			numb, _ := strconv.Atoi(rule[0])
			rules[numb] = NewRule(rule[1])
		case "messages":
			messages = append(messages, strings.Split(line, ""))

		}
	}

	return rules, messages
}

func isEqual(strArr []string, target string, start, end int) bool {
	if end > len(strArr) {
		return false
	}
	s := strings.Join(strArr[start:end], "")

	if s != target {
		return false
	}
	return true
}

type Snapshot struct {
	toVisit  []int
	matchStr string
}

func newSnapshot(tv []int, ms string) *Snapshot {
	s := Snapshot{toVisit: tv, matchStr: ms}

	return &s
}

func isMessageValid(message []string, rules Rules) bool {
	rNum := -1
	matchStr := ""
	var branches [][]int
	var toVisit []int
	var snapshots []Snapshot
	toVisit = append(rules[0].sub1, toVisit...)

	for len(toVisit) > 0 {

		rNum = toVisit[0]
		toVisit = toVisit[1:]

		// we hit a character rule
		if rules[rNum].char != "" {
			matchStr += rules[rNum].char

			if !isEqual(message, matchStr, 0, len(matchStr)) {
				if len(branches) > 0 {

					snap := snapshots[0]
					snapshots = snapshots[1:]
					toVisit = snap.toVisit
					matchStr = snap.matchStr
					toVisit = append(branches[len(branches)-1], toVisit...)

					branches = branches[0 : len(branches)-1]
				} else {
					return false
				}
			}
		} else {
			snap := newSnapshot(toVisit, matchStr)
			toVisit = append(rules[rNum].sub1, toVisit...)
			if len(rules[rNum].sub2) > 0 {

				snapshots = append([]Snapshot{*snap}, snapshots...)
				branches = append(branches, rules[rNum].sub2)
			}
		}

	}

	return matchStr == strings.Join(message, "")
}

func updateRulesForPart2(rules Rules) {
	rule8 := NewRule("42 | 42 8")
	rule11 := NewRule("42 31 | 42 11 31")
	rules[8] = rule8
	rules[11] = rule11

}

func Day19() {
	inputs := []string{"test01", "test02", "puzzle"} //
	for _, f := range inputs {
		path := fmt.Sprintf("input/day19/%s.input", f)
		rules, messages := readRules(path)

		ansP1 := 0
		for _, ms := range messages {
			if isMessageValid(ms, rules) {
				ansP1++
			}
		}
		ansP2 := 0
		updateRulesForPart2(rules)
		if f != "test01" {
			for _, ms := range messages {
				if isMessageValid(ms, rules) {
					ansP2++
				}
			}
		}

		fmt.Printf("%s: part1: %d | part2: %d\n", f, ansP1, ansP2)
	}
}
