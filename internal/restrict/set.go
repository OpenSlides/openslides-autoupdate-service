package restrict

// set a datastructure of unique items.
type set struct {
	m map[int]struct{}
}

func newSet() *set {
	return &set{m: make(map[int]struct{})}
}

func (s set) Add(e int) {
	(s.m)[e] = struct{}{}
}

func (s set) Has(e int) bool {
	_, ok := (s.m)[e]
	return ok
}

func (s set) List() []int {
	out := make([]int, 0, len(s.m))
	for k := range s.m {
		out = append(out, k)
	}
	return out
}

func (s set) Remove(es ...int) {
	for _, e := range es {
		delete(s.m, e)
	}
}
