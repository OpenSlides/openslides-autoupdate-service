package set

// Set a datastructure of unique items.
type Set struct {
	m map[int]struct{}
}

// New initializes a new set.
func New(init ...int) *Set {
	s := &Set{m: make(map[int]struct{})}
	s.Add(init...)
	return s
}

// Add adds elements to the set.
func (s Set) Add(es ...int) {
	for _, e := range es {
		s.m[e] = struct{}{}
	}
}

// Has tells, if the element exists in the set.
func (s Set) Has(e int) bool {
	_, ok := s.m[e]
	return ok
}

// List returns all elements of the set.
func (s Set) List() []int {
	out := make([]int, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	return out
}

// Remove removes the elements from the set.
func (s Set) Remove(es ...int) {
	for _, e := range es {
		delete(s.m, e)
	}
}

// Len returns the amout of elements in the set.
func (s Set) Len() int {
	return len(s.m)
}
