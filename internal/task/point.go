package task

type Point struct {
	X int
	Y int
}

func (p Point) Add(dir Point) Point {
	return Point{
		X: p.X + dir.X,
		Y: p.Y + dir.Y,
	}
}

type Block struct {
	Point Point
	Color int
}
