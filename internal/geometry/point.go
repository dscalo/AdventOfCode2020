package geometry

type Point struct {
	X int
	Y int
}

func (p *Point) MoveX(to int) {
	p.X += to
}

func (p *Point) MoveY(to int) {
	p.Y += to
}

func (p *Point) MovePoint(m *Point) {
	p.X += m.X
	p.Y += m.Y
}

func NewPoint(x int, y int) *Point {
	p := Point{X: x, Y: y}

	return &p
}

func NewPoints(xys [][]int) *[]Point {
	ps := make([]Point, len(xys))

	for idx, xy := range xys {
		p := NewPoint(xy[0], xy[1])

		ps[idx] = *p
	}

	return &ps
}

var Directions = map[string]*Point{
	"up":        NewPoint(0, -1),
	"right":     NewPoint(1, 0),
	"down":      NewPoint(0, 1),
	"left":      NewPoint(-1, 0),
	"upRight":   NewPoint(1, -1),
	"upLeft":    NewPoint(-1, -1),
	"downRight": NewPoint(1, 1),
	"downLeft":  NewPoint(-1, 1),
}
