package main

import (
	"fmt"
	"github.com/dscalo/AdventOfCode2020/puzzles"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	if month == 12 && year == 2020 {
		if fn, ok := puzzles.Days[day]; ok {
			fmt.Printf("---- Day %d ----\n", day)
			fn()
		} else {
			fmt.Printf("Day %d not found", day)
		}

	} else {
		for _, fn := range puzzles.Days {
			fmt.Printf("---- Day %d ----\n", day)
			fn()
		}
	}

}
