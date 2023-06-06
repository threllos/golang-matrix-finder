package task

type Group struct {
	Group []Point
	Color int
	Size  int
}

func NewGroup() Group {
	return Group{
		Group: make([]Point, 0),
		Size:  0,
	}
}

func (g *Group) Append(p Point) {
	g.Group = append(g.Group, p)
	g.Size += 1
}
