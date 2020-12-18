package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	Days[18] = Day18
}

type Precedence = map[string]int

func getType(char string) string {
	switch char {
	case "(":
		return "LP"
	case ")":
		return "RP"
	case "+":
		return "A"
	case "*":
		return "M"
	case " ":
		return "S"
	default:
		matched, err := regexp.MatchString(`\d`, char)
		if err != nil || matched == false {
			return ""
		}
		return "N"
	}
}

func parseTokens(tokens *Tokens, prec Precedence) *Queue {
	stack := NewStack()
	queue := NewQueue()

	//PrettyPrintTokens(*tokens)
	for _, token := range *tokens {
		switch token.Type {
		case "N":
			queue.Enqueue(&token)
		case "A":
			fallthrough
		case "M":
			for stack.HasNext() && (stack.Peek().Type == "A" || stack.Peek().Type == "M") && (prec[stack.Peek().Type] >= prec[token.Type]) {
				queue.Enqueue(stack.Pop())
			}
			stack.Push(&token)
		case "LP":
			stack.Push(&token)
		case "RP":
			for stack.HasNext() && stack.Peek().Type != "LP" {
				queue.Enqueue(stack.Pop())
			}
			stack.Pop()
		}
	}

	//queue.PrettyPrint()
	//queue.Reverse()
	for stack.HasNext() {
		queue.Enqueue(stack.Pop())
	}

	return queue
}

func solve(RPN Queue) int64 {
	stack := NewStack()
	for RPN.HasNext() {
		token := RPN.Dequeue()
		//fmt.Printf("%s: %d\n",token.Type, token.Value)
		switch token.Type {
		case "N":
			stack.Push(token)
		case "A":
			t1 := stack.Pop()
			t2 := stack.Pop()
			t3 := NewToken("N", t1.Value+t2.Value)
			//fmt.Printf(" %d + %d == %d\n",t1.Value,t2.Value,t3.Value)
			stack.Push(t3)
		case "M":
			t1 := stack.Pop()
			t2 := stack.Pop()

			t3 := NewToken("N", t1.Value*t2.Value)
			//fmt.Printf(" %d * %d == %d\n",t1.Value,t2.Value,t3.Value)
			stack.Push(t3)
		}
	}

	result := stack.Pop()
	return result.Value
}

func tokenizeString(str string) Tokens {
	var tkns Tokens
	chars := strings.Split(str, "")

	for _, c := range chars {
		tpe := getType(c)
		switch tpe {
		case "S":
			continue
		case "N":
			numb, _ := strconv.Atoi(c)
			token := NewToken(tpe, int64(numb))
			tkns = append(tkns, *token)
		default:
			tkns = append(tkns, *NewToken(tpe, 0))
		}
	}

	return tkns
}

func readExpressions(path string) []Tokens {
	file, err := os.Open(path)
	var tokens []Tokens
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		tokens = append(tokens, tokenizeString(scanner.Text()))
	}

	return tokens
}

func Day18() {
	inputs := []string{"Test01", "test02", "puzzle"} // ,

	part1Prec := Precedence{"A": 1, "M": 1}
	part2Prec := Precedence{"A": 2, "M": 1}
	for _, f := range inputs {
		path := fmt.Sprintf("input/day18/%s.input", f)

		tokens := readExpressions(path)

		var ansP1 int64 = 0
		var ansP2 int64 = 0
		for _, tks := range tokens {
			RPNP1 := parseTokens(&tks, part1Prec)
			RPNP2 := parseTokens(&tks, part2Prec)

			ansP1 += solve(*RPNP1)
			ansP2 += solve(*RPNP2)
		}

		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
