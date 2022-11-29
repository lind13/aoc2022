package set

type Set[K comparable] struct {
	items map[K]struct{}
}

func (s *Set[K]) Has(v K) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[K]) Add(v K) {
	s.items[v] = struct{}{}
}

func (s *Set[K]) Remove(v K) {
	delete(s.items, v)
}

func (s *Set[K]) Clear() {
	s.items = make(map[K]struct{})
}

func (s *Set[K]) Len() int {
	return len(s.items)
}

func (s *Set[K]) GetMap() map[K]struct{} {
	return s.items
}

func New[K comparable]() *Set[K] {
	s := &Set[K]{}
	s.items = make(map[K]struct{})
	return s
}
