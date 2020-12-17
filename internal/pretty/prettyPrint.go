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

func Print2dInt(arr [][]int) {
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[x]); y++ {
			fmt.Printf("%d ", arr[x][y])
		}
		fmt.Print("\n")
	}
}

func Print3dString(arr [][][]string) {
	for z := 0; z < len(arr); z++ {
		fmt.Printf("----- Z == %d -----\n", z)
		Print2dString(arr[z])
	}
}
