package geometry

type HyperCube struct {
	C Cube
	W int
}

type HyperCubes struct {
	MAP   map[HyperCube]string
	HighZ int
	LowZ  int
	HighY int
	LowY  int
	HighX int
	LowX  int
	HighW int
	LowW  int
}
type HyperCubeArray []HyperCube

func (cs *HyperCubes) Add(cube *HyperCube, state string) {
	if cube.C.X > cs.HighX {
		cs.HighX = cube.C.X
	}
	if cube.C.Y > cs.HighY {
		cs.HighY = cube.C.Y
	}
	if cube.C.Z > cs.HighZ {
		cs.HighZ = cube.C.Z
	}

	if cube.W > cs.HighW {
		cs.HighW = cube.W
	}

	if cube.C.X < cs.LowX {
		cs.LowX = cube.C.X
	}
	if cube.C.Y < cs.LowY {
		cs.LowY = cube.C.Y
	}
	if cube.C.Z < cs.LowZ {
		cs.LowZ = cube.C.Z
	}

	if cube.W < cs.LowW {
		cs.LowW = cube.W
	}

	cs.MAP[*cube] = state
}

func (cs *HyperCubes) GetState(cube HyperCube) string {
	state := "."

	if s, ok := cs.MAP[cube]; ok {
		state = s
	}

	return state
}

func (cs *HyperCubes) CountState(target string) int {
	count := 0

	for _, state := range cs.MAP {
		if state == target {
			count++
		}
	}

	return count
}

func NewHyperCubes() *HyperCubes {
	c := HyperCubes{
		MAP:   map[HyperCube]string{},
		HighX: 0,
		LowX:  0,
		HighY: 0,
		LowY:  0,
		HighZ: 0,
		LowZ:  0,
		HighW: 0,
		LowW:  0,
	}
	return &c
}

func (c *HyperCube) Add(c2 *HyperCube) *HyperCube {
	x := c.C.X + c2.C.X
	y := c.C.Y + c2.C.Y
	z := c.C.Z + c2.C.Z
	w := c.W + c2.W
	return NewHyperCube(x, y, z, w)
}

func (c *HyperCube) Equals(m *HyperCube) bool {
	if c.C.Z != m.C.Z {
		return false
	}

	if c.C.Y != m.C.Y {
		return false
	}

	if c.C.X != m.C.X {
		return false
	}

	if c.W != m.W {
		return false
	}

	return true
}

func (c *HyperCube) NeighborList() HyperCubeArray {
	var neighbors HyperCubeArray

	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					c1 := NewHyperCube(c.C.X+x, c.C.Y+y, c.C.Z+z, c.W+w)
					if c.Equals(c1) {
						continue
					}
					neighbors = append(neighbors, *c1)
				}
			}
		}
	}
	return neighbors
}

func NewHyperCube(x, y, z, w int) *HyperCube {
	c := Cube{X: x, Y: y, Z: z}
	h := HyperCube{C: c, W: w}
	return &h
}
