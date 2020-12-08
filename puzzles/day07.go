package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Days[7] = Day07
}

type bagRule struct {
	bags map[string]int
}

type bagRules = map[string]*bagRule

func (br *bagRule) prettyPrint() {
	if len(br.bags) == 0 {
		fmt.Println("No bags")
		return
	}

	for bag, qt := range br.bags {
		fmt.Printf("%s: %d\n", bag, qt)
	}
}

func newBagRule(r string) *bagRule {
	br := bagRule{bags: make(map[string]int)}
	rule := strings.TrimSpace(r)
	if rule != "no other" {
		bags := strings.Split(rule, ", ")
		for _, b := range bags {
			words := strings.Split(strings.TrimSpace(b), " ")
			qt, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}
			name := strings.Join(words[1:], "_")

			br.bags[name[0:len(name)]] = qt
		}
	}

	return &br
}

func printBags(bags bagRules) {
	for b, brs := range bags {
		fmt.Printf("---- %s ----\n", b)
		brs.prettyPrint()

	}
}

func contains(arr []string, target string) bool {
	for _, a := range arr {
		if a == target {
			return true
		}
	}
	return false
}

func containsTargetBag(graph map[string][]string, target string) int {
	var visited []string
	var visit []string

	visit = append(visit[0:], graph[target]...)

	for len(visit) > 0 {
		v := visit[0]
		visit = append(visit[1:], graph[v]...)
		if !contains(visited, v) {
			visited = append(visited[0:], v)
		}
	}

	return len(visited)
}

func numberOfBags(rules bagRules, target string) int {
	if len(rules[target].bags) == 0 {
		return 0
	}
	ct := 0
	for name, qt := range rules[target].bags {
		ct += qt + (qt * numberOfBags(rules, name))
	}

	return ct
}

func createGraph(bags bagRules) *map[string][]string {
	//printBags(bags)
	graph := make(map[string][]string)
	for bag, brs := range bags {
		for b := range brs.bags {
			if _, ok := graph[b]; ok {

				graph[b] = append(graph[b], bag)

			} else {
				arr := []string{bag}
				graph[b] = arr
			}
		}
	}
	return &graph
}

//shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
func readLuggageRules(path string) *bagRules {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	bags := make(map[string]*bagRule)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "bags", "")
		line = strings.ReplaceAll(line, "bag", "")
		line = strings.ReplaceAll(line, ".", "")

		rule := strings.Split(line, "contain")
		name := strings.TrimSpace(rule[0])
		name = strings.ReplaceAll(name, " ", "_")
		bags[name] = newBagRule(rule[1])

	}
	//for k,v := range graph {
	//	fmt.Print(k,": ", v, "\n")
	//}
	return &bags
}

func Day07() {
	inputs := []string{"test01", "test02", "puzzle"} //
	for _, f := range inputs {
		path := fmt.Sprintf("input/day07/%s.input", f)
		bags := readLuggageRules(path)
		graph := createGraph(*bags)

		ansP1 := containsTargetBag(*graph, "shiny_gold")
		ansP2 := numberOfBags(*bags, "shiny_gold")
		fmt.Printf("%s: part1: %d | part2: %d\n", f, ansP1, ansP2)
	}
}
