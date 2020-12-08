package puzzles

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/fs"
	"os"
	"strconv"
	"strings"
)

func init() {
	Days[8] = Day08
}

type instruction struct {
	op  string
	arg int
	ran int
}

func (instr *instruction) clear() {
	instr.ran = 0
}

func (instr *instruction) prettyPrint() {
	fmt.Printf("%s: %d | ran: %d\n", instr.op, instr.arg, instr.ran)
}

func (instr *instruction) flip() {
	if instr.op == "acc" {
		return
	}
	if instr.op == "jmp" {
		instr.op = "nop"
	} else {
		instr.op = "jmp"
	}

}

func clearInstrs(instrs []instruction) {
	for idx := range instrs {
		instrs[idx].clear()
	}
}

func printInstrs(instrs []instruction) {
	fmt.Printf("-------------------------\n")
	for idx, instr := range instrs {
		fmt.Printf("%d:", idx)
		instr.prettyPrint()
	}
	fmt.Printf("-------------------------\n")
}

func newInstr(op string, arg int) *instruction {
	instr := instruction{op: op, arg: arg, ran: 0}

	return &instr
}

func readInstructions(path string) []instruction {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	size, err := fs.LineCount(path)

	instructions := make([]instruction, size)
	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(line[1])
		instructions[idx] = *newInstr(line[0], num)
		idx++
	}

	return instructions
}

func runInstructions(instructions []instruction) (int, error) {
	acc := 0
	line := 0

	for {
		//printInstrs(instructions)
		if line >= len(instructions) {
			break
		}
		switch instructions[line].op {
		case "acc":
			instructions[line].ran += 1
			if instructions[line].ran > 1 {
				return acc, errors.New("infinite loop")
			}
			acc += instructions[line].arg
			line++
		case "jmp":
			instructions[line].ran += 1
			if instructions[line].ran > 1 {
				return acc, errors.New("infinite loop")
			}
			line += instructions[line].arg
		case "nop":
			instructions[line].ran += 1
			if instructions[line].ran > 1 {
				return acc, errors.New("infinite loop")
			}
			line++
		}
	}

	return acc, nil
}

func flipNext(instructions []instruction, changeIdx int) int {
	validIdx := false
	for validIdx != true {
		if instructions[changeIdx].op == "nop" || instructions[changeIdx].op == "jmp" {
			instructions[changeIdx].flip()
			validIdx = true
			continue
		}
		changeIdx++

	}
	return changeIdx
}

func fixProgram(instructions []instruction) int {
	changeIdx := -1

	for {
		clearInstrs(instructions)
		idx := flipNext(instructions, changeIdx+1)
		changeIdx = idx
		acc, err := runInstructions(instructions)

		if err == nil {
			return acc
		}

		instructions[changeIdx].flip()

	}

}

func Day08() {
	inputs := []string{"test01", "puzzle"} //
	for _, f := range inputs {
		path := fmt.Sprintf("input/day08/%s.input", f)
		instructions := readInstructions(path)

		ansP1, _ := runInstructions(instructions)

		ansP2 := fixProgram(instructions)

		fmt.Printf("%s: part 1: %d | part 2: %d\n", f, ansP1, ansP2)
	}
}
