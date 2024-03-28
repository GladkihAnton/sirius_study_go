package tests

// IntSet1 implements a mathematical set
// of integers, elements are unique.
type IntSet1 struct {
	elems map[int]struct{}
}

// MakeIntSet1 creates an empty set.
func MakeIntSet1() IntSet1 {
	elems := make(map[int]struct{})
	return IntSet1{elems}
}

// Contains reports whether an element is within the set.
func (s IntSet1) Contains(elem int) bool {
	_, ok := s.elems[elem]
	return ok
}

// Add adds an element to the set.
func (s IntSet1) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	s.elems[elem] = struct{}{}
	return true
}
