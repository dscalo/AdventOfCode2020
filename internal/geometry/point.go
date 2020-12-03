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
