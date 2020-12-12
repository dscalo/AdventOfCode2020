package puzzles

import (
	"bufio"
	"fmt"
	"github.com/dscalo/AdventOfCode2020/internal/fs"
	"github.com/dscalo/AdventOfCode2020/internal/geometry"
	"os"
	"strconv"
)

func init() {
	Days[12] = Day12
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

type Nav struct {
	action string
	value  int
}

func newNav(action string, value int) *Nav {
	n := Nav{action: action, value: value}
	return &n
}
func (n *Nav) prettyPrint() {
	fmt.Printf("Action: \"%s\", Value: %d\n", n.action, n.value)
}

func readNavInstructions(path string) []Nav {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	size, err := fs.LineCount(path)

	if err != nil {
		panic(err)
	}

	navs := make([]Nav, size)

	scanner := bufio.NewScanner(file)

	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		action := line[0:1]
		value, _ := strconv.Atoi(line[1:])
		navs[idx] = *newNav(action, value)
		idx++
	}
	return navs
}

func move(ship *geometry.Point, dir string, value int) {
	switch dir {
	case "N":
		// up y + value
		ship.MoveY(value)
	case "S":
		// down y - value
		ship.MoveY(-1 * value)
	case "E":
		// right x + value
		ship.MoveX(value)
	case "W":
		// left x - value
		ship.MoveX(-1 * value)
	}
}
func turn(ship *geometry.Point, nav Nav) {
	dir := ship.Dir
	switch nav.action {
	case "L":
		dir -= nav.value
	case "R":
		dir += nav.value
	}
	if dir < 0 {
		dir += 360
	}

	if dir > 360 {
		dir -= 360
	}

	ship.Dir = dir
}

func rotateWaypoint(wp *geometry.Point, nav Nav) {
	x := wp.X
	y := wp.Y
	switch nav.action {
	// counter clockwise
	case "L":
		switch nav.value {
		case 90:
			wp.X = -y
			wp.Y = x
		case 180:
			wp.X = -x
			wp.Y = -y
		case 270:
			wp.X = y
			wp.Y = -x
		}
	// clock wise
	case "R":
		switch nav.value {
		case 90:
			wp.X = y
			wp.Y = -x
		case 180:
			wp.X = -x
			wp.Y = -y
		case 270:
			wp.X = -y
			wp.Y = x
		}
	}
}

func degToDir(deg int) string {
	switch deg {
	case 0:
		fallthrough
	case 360:
		return "N"
	case 90:
		return "E"
	case 180:
		return "S"
	case 270:
		return "W"
	}
	return ""
}

func moveShip(ship *geometry.Point, nav Nav) {
	switch nav.action {
	case "L":
		fallthrough
	case "R":
		turn(ship, nav)
	case "F":
		move(ship, degToDir(ship.Dir), nav.value)
	default:
		move(ship, nav.action, nav.value)
	}
}

func moveShipToWaypoint(ship *geometry.Point, wp *geometry.Point, nav Nav) {
	ship.MoveX(wp.X * nav.value)
	ship.MoveY(wp.Y * nav.value)
}

func moveWaypoint(ship *geometry.Point, wp *geometry.Point, nav Nav) {
	switch nav.action {
	case "L":
		fallthrough
	case "R":
		rotateWaypoint(wp, nav)
	case "F":
		moveShipToWaypoint(ship, wp, nav)
	default:
		move(wp, nav.action, nav.value)
	}
}

func executeInstructions(navs []Nav, part int) int {
	ship := geometry.NewPoint(0, 0)
	ship.Dir = 90
	wp := geometry.NewPoint(10, 1)

	for _, nav := range navs {
		if part == 1 {
			moveShip(ship, nav)
		} else {
			moveWaypoint(ship, wp, nav)
		}
	}
	return abs(ship.X) + abs(ship.Y)
}

func Day12() {
	inputs := []string{"test01", "test02", "test03", "puzzle"} // , "puzzle"
	for _, f := range inputs {
		path := fmt.Sprintf("input/day12/%s.input", f)

		navs := readNavInstructions(path)

		ansP1 := executeInstructions(navs, 1)
		ansP2 := executeInstructions(navs, 2)

		fmt.Printf("%s part 1 : %d | part2: %d  \n", f, ansP1, ansP2)
	}
}
