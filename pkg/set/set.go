package set

type void struct{}

var member void

type Set[T comparable] map[T]void

func NewSet[T comparable]() Set[T] {
	return make(map[T]void)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Get() (T, bool) {
	for value := range s {
		return value, true
	}
	return *new(T), false
}

func (s *Set[T]) Add(value T) {
	(*s)[value] = member
}

func (s *Set[T]) Remove(value T) {
	delete(*s, value)
}

func (s *Set[T]) Pop() (T, bool) {
	value, ok := s.Get()
	if !ok {
		return *new(T), false
	}
	s.Remove(value)
	return value, true
}

func (s Set[T]) Exist(value T) bool {
	_, ok := s[value]
	return ok
}
