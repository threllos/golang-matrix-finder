package task

type Point struct {
	X int
	Y int
}

type Block struct {
	X     int
	Y     int
	Color int
}

func (p Point) Add(dir Point) Point {
	return Point{
		X: p.X + dir.X,
		Y: p.Y + dir.Y,
	}
}
