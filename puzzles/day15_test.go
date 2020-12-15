package puzzles

import "testing"

func Test_playGameTo2020(t *testing.T) {
	tests := map[int][]int{
		436: []int{0, 3, 6},
		1:   []int{1, 3, 2},
		10:  []int{2, 1, 3},
		27:  []int{1, 2, 3},
		78:  []int{2, 3, 1},
		438: []int{3, 2, 1},
	}
	for expected, test := range tests {
		result := playGame(test, 2020)
		if result != expected {
			t.Errorf("expected %d, received: %d", expected, result)
		}
	}

}

func Test_playGameTo30000000(t *testing.T) {
	tests := map[int][]int{
		175594:  []int{0, 3, 6},
		2578:    []int{1, 3, 2},
		3544142: []int{2, 1, 3},
		261214:  []int{1, 2, 3},
		6895259: []int{2, 3, 1},
		18:      []int{3, 2, 1},
		362:     []int{3, 1, 2},
	}
	for expected, test := range tests {
		result := playGame(test, 30000000)
		if result != expected {
			t.Errorf("expected %d, received: %d", expected, result)
		}
	}

}
