package stack

type Stack[T any] []T

func NewStack[T any]() Stack[T] {
	return make(Stack[T], 0)
}

func (s Stack[T]) Size() int {
	return len(s)
}

func (s *Stack[T]) Push(value T) {
	*s = append((*s), value)
}

func (s *Stack[T]) Pop() (T, bool) {
	l := s.Size()
	if l == 0 {
		return *new(T), false
	}
	value := (*s)[l-1]
	*s = (*s)[:l-1]
	return value, true
}
