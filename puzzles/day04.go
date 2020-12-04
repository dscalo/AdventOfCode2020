package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readPasswordFile(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	validPasswords := 0
	pp := NewPassport()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			//pp.prettyPrint()

			if pp.IsValid() {
				validPasswords++
			}
			pp = NewPassport()
			continue
		}
		kvs := strings.Split(line, " ")

		for _, kv := range kvs {
			pp.AddField(kv)
		}
	}
	if pp.IsValid() {
		validPasswords++
	}

	return validPasswords
}

func Day04() {
	/*
		test02: all valid (4)
		test03: all INVALID (0)
	*/
	inputs := []string{"test01", "test02", "test03", "puzzle"} //
	for _, f := range inputs {
		path := fmt.Sprintf("input/day04/%s.input", f)
		ans := readPasswordFile(path)
		fmt.Printf("%s: valid: %d\n", f, ans)
	}
}
