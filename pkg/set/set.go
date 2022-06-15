package set

// Set a datastructure of unique items.
type Set[T comparable] struct {
	m map[T]struct{}
}

// New initializes a new set.
func New[T comparable](init ...T) *Set[T] {
	s := &Set[T]{m: make(map[T]struct{})}
	s.Add(init...)
	return s
}

// Add adds elements to the set.
func (s Set[T]) Add(es ...T) {
	for _, e := range es {
		s.m[e] = struct{}{}
	}
}

// Merge adds all elements from the other set to this set.
func (s Set[T]) Merge(other Set[T]) {
	for e := range other.m {
		s.m[e] = struct{}{}
	}
}

// Has tells, if the element exists in the set.
func (s Set[T]) Has(e T) bool {
	_, ok := s.m[e]
	return ok
}

// List returns all elements of the set.
func (s Set[T]) List() []T {
	out := make([]T, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	return out
}

// Remove removes the elements from the set.
func (s Set[T]) Remove(es ...T) {
	for _, e := range es {
		delete(s.m, e)
	}
}

// Len returns the amout of elements in the set.
func (s Set[T]) Len() int {
	return len(s.m)
}

// Equal returns true if both sets have the same values
func Equal[T comparable](s1, s2 *Set[T]) bool {
	if s1.Len() != s2.Len() {
		return false
	}

	for k := range s1.m {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}
