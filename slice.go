package sealed

// TODO
// Slice.Index(func) int
// Slice.Contains(func) bool
// Slice.Slice(start, end int) Slice[T]
// PopFirst, PopLast

// Slice is an immutable representation of a standard Go slice.
// Use Builder.Seal to create a new slice.
// The zero-value of Slice is a valid representation of an empty slice.
type Slice[T any] struct {
	s []T
}

func (s Slice[T]) Len() int {
	return len(s.s)
}

// Get return the ith element of s.
// Exactly like normal raw slice access this method will panic if i is out of bounds.
func (s Slice[T]) Get(i int) T {
	return s.s[i]
}

// All iterates over all elements in s.
// Return false to stop early.
func (s Slice[T]) All(yield func(int, T) bool) {
	for i, t := range s.s {
		if !yield(i, t) {
			return
		}
	}
}

func (s Slice[T]) Empty() bool {
	return len(s.s) == 0
}

func (s Slice[T]) First() (T, bool) {
	if len(s.s) == 0 {
		var zero T
		return zero, false
	}
	return s.s[0], true
}

func (s Slice[T]) Last() (T, bool) {
	if len(s.s) == 0 {
		var zero T
		return zero, false
	}
	return s.s[len(s.s)-1], true
}
