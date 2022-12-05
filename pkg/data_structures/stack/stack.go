package stack

type Stack[T comparable] struct {
	internal []T
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		internal: make([]T, 0),
	}
}

func NewFromArr[T comparable](a []T) *Stack[T] {
	return &Stack[T]{
		internal: a,
	}
}

func (s *Stack[T]) Push(x T) int {
	s.internal = append(s.internal, x)
	return len(s.internal)
}

func (s *Stack[T]) Pop() T {
	var x T
	x, s.internal = s.internal[len(s.internal)-1], s.internal[:len(s.internal)-1]
	return x
}

func (s *Stack[T]) Len() int {
	return len(s.internal)
}
