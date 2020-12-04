package puzzles

import (
	"testing"
)

func Test_isValidBYR(t *testing.T) {
	tests := map[string]bool{
		"2002": true,
		"2003": false,
		"":     false,
	}

	for test, expect := range tests {
		result := isValidBYR(test)
		if result != expect {
			t.Errorf("testing: %s, expected %t recieved: %t", test, expect, result)
		}
	}
}

func Test_isValidIYR(t *testing.T) {
	tests := map[string]bool{
		"2010": true,
		"1988": false,
		"":     false,
	}

	for test, expect := range tests {
		result := isValidIYR(test)
		if result != expect {
			t.Errorf("testing: %s, expected %t recieved: %t", test, expect, result)
		}
	}
}

func Test_isValidBEYR(t *testing.T) {
	tests := map[string]bool{
		"2025": true,
		"2003": false,
		"":     false,
	}

	for test, expect := range tests {
		result := isValidEYR(test)
		if result != expect {
			t.Errorf("testing: %s, expected %t recieved: %t", test, expect, result)
		}
	}
}

func Test_isValidHGT(t *testing.T) {
	tests := map[string]bool{
		"60in":  true,
		"190cm": true,
		"190in": false,
		"190":   false,
		"":      false,
	}

	for test, expect := range tests {
		result := isValidHGT(test)
		if result != expect {
			t.Errorf("testing: %s, expected %t recieved: %t", test, expect, result)
		}
	}
}

func Test_isValidHCL(t *testing.T) {
	tests := map[string]bool{
		"#123abc": true,
		"#123abz": false,
		"123abc":  false,
		"":        false,
	}

	for test, expect := range tests {
		result := isValidHCL(test)
		if result != expect {
			t.Errorf("testing: %s, expected %t recieved: %t", test, expect, result)
		}
	}
}

func Test_isValidECL(t *testing.T) {
	tests := map[string]bool{
		"brn":    true,
		"banana": false,
		"":       false,
	}

	for test, expect := range tests {
		result := isValidECL(test)
		if result != expect {
			t.Errorf("testing: %s, expected %t recieved: %t", test, expect, result)
		}
	}
}

func Test_isValidPID(t *testing.T) {
	tests := map[string]bool{
		"000000001":  true,
		"0123456789": false,
		"":           false,
	}

	for test, expect := range tests {
		result := isValidPID(test)
		if result != expect {
			t.Errorf("testing: %s, expected %t recieved: %t", test, expect, result)
		}
	}
}
