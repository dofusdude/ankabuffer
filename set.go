package ankabuffer

// based on string only https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80
// rewritten to be generic

type Set[T comparable] struct {
	list map[T]struct{} //empty structs occupy 0 memory
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set[T]) Add(v T) {
	s.list[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.list, v)
}

func (s *Set[T]) Clear() {
	s.list = make(map[T]struct{})
}

func (s *Set[T]) Size() int {
	return len(s.list)
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.list = make(map[T]struct{})
	return s
}

func (s *Set[T]) Slice() []T {
	var res = make([]T, 0, len(s.list))
	for k := range s.list {
		res = append(res, k)
	}
	return res
}

func (s *Set[T]) AddMulti(list ...T) {
	for _, v := range list {
		s.Add(v)
	}
}

// fp

func Map[T, U any](data []T, f func(T) U) []U {

	res := make([]U, 0, len(data))

	for _, e := range data {
		res = append(res, f(e))
	}

	return res
}
