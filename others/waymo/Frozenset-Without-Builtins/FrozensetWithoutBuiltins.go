package waymo

type FrozenSet struct {
	elems []int
}

func NewFrozenSet(values []int) *FrozenSet {
	set := &FrozenSet{}
	for _, value := range values {
		set = set.With(value)
	}
	return set
}

func (s *FrozenSet) Size() int {
	return len(elemsOf(s))
}

func (s *FrozenSet) Contains(value int) bool {
	_, found := lowerBound(elemsOf(s), value)
	return found
}

func (s *FrozenSet) Elements() []int {
	base := elemsOf(s)
	out := make([]int, len(base))
	copy(out, base)
	return out
}

func (s *FrozenSet) With(value int) *FrozenSet {
	base := elemsOf(s)
	index, found := lowerBound(base, value)
	if found {
		return normalizeSet(s)
	}
	next := make([]int, len(base)+1)
	copy(next, base[:index])
	next[index] = value
	copy(next[index+1:], base[index:])
	return &FrozenSet{elems: next}
}

func (s *FrozenSet) Without(value int) *FrozenSet {
	base := elemsOf(s)
	index, found := lowerBound(base, value)
	if !found {
		return normalizeSet(s)
	}
	next := make([]int, len(base)-1)
	copy(next, base[:index])
	copy(next[index:], base[index+1:])
	return &FrozenSet{elems: next}
}

func (s *FrozenSet) Union(other *FrozenSet) *FrozenSet {
	left := elemsOf(s)
	right := elemsOf(other)
	out := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			out = append(out, left[i])
			i++
			continue
		}
		if left[i] > right[j] {
			out = append(out, right[j])
			j++
			continue
		}
		out = append(out, left[i])
		i++
		j++
	}
	for i < len(left) {
		out = append(out, left[i])
		i++
	}
	for j < len(right) {
		out = append(out, right[j])
		j++
	}
	return &FrozenSet{elems: out}
}

func (s *FrozenSet) Intersection(other *FrozenSet) *FrozenSet {
	left := elemsOf(s)
	right := elemsOf(other)
	out := make([]int, 0, min(len(left), len(right)))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			i++
			continue
		}
		if left[i] > right[j] {
			j++
			continue
		}
		out = append(out, left[i])
		i++
		j++
	}
	return &FrozenSet{elems: out}
}

func (s *FrozenSet) Difference(other *FrozenSet) *FrozenSet {
	left := elemsOf(s)
	right := elemsOf(other)
	out := make([]int, 0, len(left))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			out = append(out, left[i])
			i++
			continue
		}
		if left[i] > right[j] {
			j++
			continue
		}
		i++
		j++
	}
	for i < len(left) {
		out = append(out, left[i])
		i++
	}
	return &FrozenSet{elems: out}
}

func (s *FrozenSet) Equals(other *FrozenSet) bool {
	left := elemsOf(s)
	right := elemsOf(other)
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func elemsOf(s *FrozenSet) []int {
	if s == nil {
		return []int{}
	}
	return s.elems
}

func normalizeSet(s *FrozenSet) *FrozenSet {
	if s == nil {
		return &FrozenSet{}
	}
	return s
}

func lowerBound(values []int, target int) (int, bool) {
	left, right := 0, len(values)
	for left < right {
		mid := left + (right-left)/2
		if values[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left < len(values) && values[left] == target {
		return left, true
	}
	return left, false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
