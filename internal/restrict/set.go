package restrict

// Set a datastructure of unique items.
type Set struct {
	m map[int]struct{}
}

func NewSet() *Set {
	return &Set{m: make(map[int]struct{})}
}

func (s Set) Add(e int) {
	(s.m)[e] = struct{}{}
}

func (s Set) Has(e int) bool {
	_, ok := (s.m)[e]
	return ok
}

func (s Set) List() []int {
	out := make([]int, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	return out
}

func (s Set) Remove(es ...int) {
	for _, e := range es {
		delete(s.m, e)
	}
}
