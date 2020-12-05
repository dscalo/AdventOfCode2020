package puzzles

import (
	"testing"
)

func Test_getRange1(t *testing.T) {
	lower, upper := getRange(0, 127, "F")
	if lower != 0 || upper != 63 {
		t.Errorf("expected 0, 63, received: %d %d", lower, upper)
	}
}

func Test_getRange2(t *testing.T) {
	lower, upper := getRange(0, 127, "B")
	if lower != 64 || upper != 127 {
		t.Errorf("expected 63, 127, received: %d %d", lower, upper)
	}
}

func Test_getRange3(t *testing.T) {
	lower, upper := getRange(0, 63, "B")
	if lower != 32 || upper != 63 {
		t.Errorf("expected 32, 63, received: %d %d", lower, upper)
	}
}

func Test_getRange4(t *testing.T) {
	lower, upper := getRange(32, 47, "B")
	if lower != 40 || upper != 47 {
		t.Errorf("expected 40, 47, received: %d %d", lower, upper)
	}
}

func Test_getRange5(t *testing.T) {
	lower, upper := getRange(44, 45, "F")
	if lower != 44 || upper != 44 {
		t.Errorf("expected 45, 45, received: %d %d", lower, upper)
	}
}

func Test_getRange6(t *testing.T) {
	lower, upper := getRange(4, 7, "L")
	if lower != 4 || upper != 5 {
		t.Errorf("expected 4, 5, received: %d %d", lower, upper)
	}
}

func Test_getRow(t *testing.T) {
	res := getAirplaneSection("FBFBBFF", "ROW")

	if res != 44 {
		t.Errorf("expected 44 received: %d", res)
	}
}

func Test_getCol(t *testing.T) {
	res := getAirplaneSection("RLR", "COL")

	if res != 5 {
		t.Errorf("expected 5 received: %d", res)
	}
}
