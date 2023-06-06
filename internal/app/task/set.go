package task

type void struct{}

var member void

type SetP map[Point]void

func NewSetP() SetP {
	return make(map[Point]void)
}

func (s SetP) Add(p Point) SetP {
	s[p] = member
	return s
}

func (s SetP) Delete(p Point) SetP {
	delete(s, p)
	return s
}

func (s SetP) Exist(p Point) bool {
	_, ok := s[p]
	return ok
}

func (s SetP) Size() int {
	return len(s)
}

func (s SetP) Get() (Point, bool) {
	for p := range s {
		return p, true
	}
	return Point{}, false
}

type SetB map[Block]void

func NewSetB() SetB {
	return make(map[Block]void)
}

func (s SetB) Add(b Block) SetB {
	s[b] = member
	return s
}

func (s SetB) Delete(b Block) SetB {
	delete(s, b)
	return s
}

// func (s SetP) Exist(p Point) bool {
// 	_, ok := s[p]
// 	return ok
// }

func (s SetB) Size() int {
	return len(s)
}

func (s SetB) Get() (Block, bool) {
	for b := range s {
		return b, true
	}
	return Block{}, false
}
