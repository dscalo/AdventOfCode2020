package puzzles

import (
	"bufio"
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/geometry"
	"os"
	"strings"
)

func init() {
	Days[17] = Day17
}

func readCubes(path string) (geometry.Cubes, geometry.HyperCubes) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	cubes := geometry.NewCubes()
	hyperCubes := geometry.NewHyperCubes()

	scanner := bufio.NewScanner(file)
	y := 0
	x := 0
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		x = 0
		for _, state := range row {
			if state == "#" {
				cube1 := geometry.NewCube(x, y, 0)
				hyperCube := geometry.NewHyperCube(x, y, 0, 0)
				cubes.Add(cube1, "#")
				hyperCubes.Add(hyperCube, "#")
			}
			x++
		}
		y++
	}
	return *cubes, *hyperCubes
}

func bootUp(cubes *geometry.Cubes) {
	var active geometry.CubeArray
	var inactive geometry.CubeArray

	for z := cubes.LowZ - 1; z <= cubes.HighZ+1; z++ {
		for y := cubes.LowY - 1; y <= cubes.HighY+1; y++ {
			for x := cubes.LowX - 1; x <= cubes.HighX+1; x++ {
				cube := geometry.NewCube(x, y, z)
				cubeState := cubes.GetState(*cube)
				neighborList := cube.NeighborList()
				activeCount := 0
				for _, n := range neighborList {
					if cubes.GetState(n) == "#" {
						activeCount++
					}
				}
				switch cubeState {
				case "#":
					if activeCount < 2 || activeCount > 3 {
						inactive = append(inactive, *cube)
					}
				case ".":
					if activeCount == 3 {
						active = append(active, *cube)
					}
				}
			}
		}
	}

	for _, c := range inactive {
		cubes.Add(&c, ".")
	}

	for _, c := range active {
		cubes.Add(&c, "#")
	}
}

func bootUpHC(cubes *geometry.HyperCubes) {
	var active geometry.HyperCubeArray
	var inactive geometry.HyperCubeArray
	for w := cubes.LowW - 1; w <= cubes.HighW+1; w++ {
		for z := cubes.LowZ - 1; z <= cubes.HighZ+1; z++ {
			for y := cubes.LowY - 1; y <= cubes.HighY+1; y++ {
				for x := cubes.LowX - 1; x <= cubes.HighX+1; x++ {
					cube := geometry.NewHyperCube(x, y, z, w)
					cubeState := cubes.GetState(*cube)
					neighborList := cube.NeighborList()
					activeCount := 0
					for _, n := range neighborList {
						if cubes.GetState(n) == "#" {
							activeCount++
						}
					}
					switch cubeState {
					case "#":
						if activeCount < 2 || activeCount > 3 {
							inactive = append(inactive, *cube)
						}
					case ".":
						if activeCount == 3 {
							active = append(active, *cube)
						}
					}
				}
			}
		}
	}

	for _, c := range inactive {
		cubes.Add(&c, ".")
	}

	for _, c := range active {
		cubes.Add(&c, "#")
	}
}

func Day17() {
	inputs := []string{"test01", "puzzle"} //

	for _, f := range inputs {
		path := fmt.Sprintf("input/day17/%s.input", f)

		cubes, hyperCubes := readCubes(path)
		for i := 0; i < 6; i++ {
			//fmt.Printf("================= %d ========================\n", i+1)
			bootUp(&cubes)
			bootUpHC(&hyperCubes)
			//pretty.Print3dString(geometry.Build3dArray(cubes))
			//fmt.Println("=========================================")
		}
		//pretty.Print3dString(geometry.Build3dArray(cubes))
		ansP1 := cubes.CountState("#")
		ansP2 := hyperCubes.CountState("#")
		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
