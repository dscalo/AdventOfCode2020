package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	Days[22] = Day22
}

func readPlayerHands(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var player1 []int
	var player2 []int

	scanner := bufio.NewScanner(file)

	tpe := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Player") {
			tpe = line
			continue
		}

		if line == "" {
			continue
		}

		n, _ := strconv.Atoi(line)
		switch tpe {
		case "Player 1:":
			player1 = append(player1, n)
		case "Player 2:":
			player2 = append(player2, n)
		default:
			panic("check your input")
		}
	}

	return player1, player2
}

func scoreHand(hand []int) int64 {
	var score int64 = 0

	multiplier := len(hand)
	for _, c := range hand {
		score += int64(c * multiplier)
		multiplier--
	}
	return score
}

func playCombat(player1, player2 []int) (int, int64) {
	winner := 0
	var score int64 = 0

	round := 1
	for winner == 0 {
		p1Card := player1[0]
		p2Card := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]

		if p1Card > p2Card {
			player1 = append(player1, p1Card, p2Card)
		} else {
			player2 = append(player2, p2Card, p1Card)
		}

		round++
		if len(player1) == 0 {
			winner = 2
			score = scoreHand(player2)
		}

		if len(player2) == 0 {
			winner = 1
			score = scoreHand(player1)
		}
	}
	return winner, score
}

func intArrToString(arr []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", "", -1), "[]")
}

func playRecursiveCombat(player1, player2 []int) (int, int64) {
	winner := 0
	var score int64 = 0

	round := 1
	tracker := map[string]int{}
	for winner == 0 {

		key := fmt.Sprintf("%s-%s", intArrToString(player1), intArrToString(player2))
		if _, ok := tracker[key]; ok {
			return 1, scoreHand(player1)
		} else {
			tracker[key] = round
		}

		p1Card := player1[0]
		p2Card := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]

		if len(player1) >= p1Card && len(player2) >= p2Card {
			p1 := make([]int, p1Card)
			copy(p1, player1[:p1Card])
			p2 := make([]int, p2Card)
			copy(p2, player2[:p2Card])

			w, _ := playRecursiveCombat(p1, p2)
			switch w {
			case 1:
				player1 = append(player1, p1Card, p2Card)
			case 2:
				player2 = append(player2, p2Card, p1Card)
			}
			continue
		}

		if p1Card > p2Card {
			player1 = append(player1, p1Card, p2Card)
		} else {
			player2 = append(player2, p2Card, p1Card)
		}

		round++
		if len(player1) == 0 {
			winner = 2
			score = scoreHand(player2)
		}

		if len(player2) == 0 {
			winner = 1
			score = scoreHand(player1)
		}
	}

	return winner, score
}

func Day22() {
	inputs := []string{"test01", "puzzle"} // , "puzzle"
	for _, f := range inputs {
		path := fmt.Sprintf("input/day22/%s.input", f)

		player1, player2 := readPlayerHands(path)

		wc, sc := playCombat(player1, player2)
		wRC, sRC := playRecursiveCombat(player1, player2)

		fmt.Printf("%s: Combat Winner is Player %d with a score of %d\n", f, wc, sc)
		fmt.Printf("%s: Recursive Combat Winner is Player %d with a score of %d\n", f, wRC, sRC)
	}

}
