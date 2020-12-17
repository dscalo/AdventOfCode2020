package geometry

import (
	"fmt"
	"sort"
)

type Cube struct {
	X int
	Y int
	Z int
}

type CubeArray []Cube

type Cubes struct {
	MAP   map[Cube]string
	HighZ int
	LowZ  int
	HighY int
	LowY  int
	HighX int
	LowX  int
}

func (cs *Cubes) GetArray() *CubeArray {
	arr := make(CubeArray, len(cs.MAP))
	idx := 0
	for cube, _ := range cs.MAP {
		arr[idx] = cube
		idx++
	}

	return &arr
}

func (cs *Cubes) Add(cube *Cube, state string) {
	if cube.X > cs.HighX {
		cs.HighX = cube.X
	}
	if cube.Y > cs.HighY {
		cs.HighY = cube.Y
	}
	if cube.Z > cs.HighZ {
		cs.HighZ = cube.Z
	}

	if cube.X < cs.LowX {
		cs.LowX = cube.X
	}
	if cube.Y < cs.LowY {
		cs.LowY = cube.Y
	}
	if cube.Z < cs.LowZ {
		cs.LowZ = cube.Z
	}

	cs.MAP[*cube] = state
}

func (cs *Cubes) GetState(cube Cube) string {

	state := "."

	if s, ok := cs.MAP[cube]; ok {
		state = s
	}

	return state
}

func (cs *Cubes) CountState(target string) int {
	count := 0

	for _, state := range cs.MAP {
		if state == target {
			count++
		}
	}

	return count
}

func NewCubes() *Cubes {
	c := Cubes{
		MAP:   map[Cube]string{},
		HighX: 0,
		LowX:  0,
		HighY: 0,
		LowY:  0,
		HighZ: 0,
		LowZ:  0,
	}
	return &c
}

func (cs CubeArray) Len() int {
	return len(cs)
}

func (cs CubeArray) Less(i, j int) bool {
	return cs[i].Z > cs[j].Z
}

func (cs CubeArray) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs CubeArray) Find(target *Cube) int {
	for idx, cube := range cs {
		if cube.Equals(target) {
			return idx
		}
	}
	return -1
}

func (cs CubeArray) FindXYZ(x int, y int, z int) int {
	for idx, cube := range cs {
		if cube.X == x && cube.Y == y && cube.Z == z {
			return idx
		}
	}
	return -1
}

func (cs CubeArray) Normalize() *Cube {
	lowZ := 0
	lowY := 0
	lowX := 0

	for _, cube := range cs {
		if cube.Z < lowZ {
			lowZ = cube.Z
		}
		if cube.Y < lowY {
			lowY = cube.Y
		}
		if cube.X < lowX {
			lowX = cube.X
		}
	}

	lowZ *= -1
	lowY *= -1
	lowX *= -1
	n := NewCube(lowX, lowY, lowZ)

	for i := 0; i < len(cs); i++ {
		cs[i].Move(n)
	}

	return n
}

func (cs CubeArray) PrettyPrint() {
	sort.Sort(cs)

	for _, cube := range cs {
		fmt.Printf("Z: %d, Y: %d, X: %d, \n", cube.Z, cube.Y, cube.X)
	}
}

func (c *Cube) MoveX(to int) {
	c.X += to
}

func (c *Cube) MoveY(to int) {
	c.Y += to
}

func (c *Cube) MoveZ(to int) {
	c.Z += to
}

func (c *Cube) Move(m *Cube) {
	c.X += m.X
	c.Y += m.Y
	c.Z += m.Z
}

func (c *Cube) Add(c2 *Cube) *Cube {
	x := c.X + c2.X
	y := c.Y + c2.Y
	z := c.Z + c2.Z

	return NewCube(x, y, z)
}

func (c *Cube) Equals(m *Cube) bool {
	if c.Z != m.Z {
		return false
	}

	if c.Y != m.Y {
		return false
	}

	if c.X != m.X {
		return false
	}

	return true
}

func (c *Cube) NeighborList() CubeArray {
	var neighbors CubeArray

	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				c1 := NewCube(c.X+x, c.Y+y, c.Z+z)
				if c.Equals(c1) {
					continue
				}
				neighbors = append(neighbors, *c1)
			}
		}
	}
	return neighbors
}

func NewCube(x int, y int, z int) *Cube {
	c := Cube{X: x, Y: y, Z: z}
	return &c
}

func Build3dArray(cs Cubes) [][][]string {
	orig := *cs.GetArray()
	cubes := make(CubeArray, len(orig))
	copy(cubes, orig)
	sort.Sort(cubes)
	normalizer := cubes.Normalize()
	sizeZ := 0
	sizeY := 0
	sizeX := 0

	for _, c := range cubes {
		if c.Z > sizeZ {
			sizeZ = c.Z
		}
		if c.Y > sizeY {
			sizeY = c.Y
		}
		if c.X > sizeX {
			sizeX = c.X
		}
	}

	arr := make([][][]string, sizeZ+1)

	for z := 0; z < len(arr); z++ {
		grid := make([][]string, sizeY+1)
		for y := 0; y < len(grid); y++ {
			row := make([]string, sizeX+1)
			for idx, _ := range row {
				row[idx] = "."
			}
			grid[y] = row
		}
		arr[z] = grid
	}

	for cube, state := range cs.MAP {
		cb := cube.Add(normalizer)
		arr[cb.Z][cb.Y][cb.X] = state
	}

	return arr
}
