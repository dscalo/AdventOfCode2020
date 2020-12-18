package puzzles

import "testing"

func Test_solve_part1(t *testing.T) {
	prec := Precedence{"A": 1, "M": 1}
	tokenStrings := map[int64]string{
		26:     "2 * 3 + (4 * 5)",
		437:    "5 + (8 * 3 + 9 + 3 * 4 * 3)",
		12240:  "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		13632:  "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
		160707: "3 + (6 * 9 * (4 * 4) * (7 + 2 + 9 * 3 + 5 + 8)) + (3 * (6 + 7 * 4 + 9 * 8) + 3 * 5 + 9 * (5 + 9))",
	}

	for ans, str := range tokenStrings {
		tokens := tokenizeString(str)
		RPN := parseTokens(&tokens, prec)
		res := solve(*RPN)

		if res != ans {
			t.Errorf("PART 1: expected : %d recieved: %d\n", ans, res)
		}
	}
}

func Test_solve_part2(t *testing.T) {
	prec := Precedence{"A": 2, "M": 1}
	tokenStrings := map[int64]string{
		51:     "1 + (2 * 3) + (4 * (5 + 6))",
		46:     "2 * 3 + (4 * 5)",
		1445:   "5 + (8 * 3 + 9 + 3 * 4 * 3)",
		669060: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		23340:  "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
	}

	for ans, str := range tokenStrings {
		tokens := tokenizeString(str)
		RPN := parseTokens(&tokens, prec)
		res := solve(*RPN)

		if res != ans {
			t.Errorf("PART 2: expected : %d recieved: %d\n", ans, res)
		}
	}
}
