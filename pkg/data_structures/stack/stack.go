package stack

type Stack[T comparable] []T

func (s *Stack[T]) Push(x T) int {
	*s = append(*s, x)
	return len(*s)
}

func (s *Stack[T]) PushN(x []T) int {
	*s = append(*s, x...)
	return len(*s)
}

func (s *Stack[T]) Pop() T {
	var x T
	x, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return x
}

func (s *Stack[T]) PopN(n int) []T {
	var x []T
	x, *s = (*s)[len(*s)-n:], (*s)[:len(*s)-n]
	return x
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
