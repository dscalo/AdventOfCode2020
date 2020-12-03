package pretty

import "fmt"

func Print2dString(arr [][]string) {
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[x]); y++ {
			fmt.Printf("%s ", arr[x][y])
		}
		fmt.Print("\n")
	}
}
