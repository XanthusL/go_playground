package mylib

type IntSet struct {
	data map[int]struct{}
}

func (s *IntSet) Init() {
	s.data = make(map[int]struct{})
}
func (s *IntSet) Size() int {
	return len(s.data)
}
func (s *IntSet) Add(n int) bool {
	if _, ok := s.data[n]; ok {
		return false
	}
	s.data[n] = struct{}{}
	return true
}

func (s *IntSet) Remove(n int) bool {
	if _, ok := s.data[n]; ok {
		delete(s.data, n)
		return true
	}
	return false
}
func (s *IntSet) Has(n int) bool {
	_, ok := s.data[n]
	return ok
}
func (s *IntSet) IsEmpty() bool {
	return len(s.data) == 0
}
