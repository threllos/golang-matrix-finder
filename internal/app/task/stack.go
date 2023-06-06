package task

type StackB []Block

func NewStackB() StackB {
	return make(StackB, 0)
}

func (s StackB) Push(b Block) StackB {
	return append(s, b)
}

func (s StackB) Pop() (StackB, Block, bool) {
	l := len(s)
	if l == 0 {
		return s, Block{}, false
	}
	return s[:l-1], s[l-1], true
}

func (s StackB) Size() int {
	return len(s)
}