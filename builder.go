package sealed

import (
	"iter"
	"slices"
)

// Builder is a write-only data structure used to create sealed Slices.
// The zero builder is valid, but a nil builder is not valid and will trigger panics if any method is invoked.
type Builder[T any] struct {
	s []T
}

func NewBuilder[T any](length, capacity int) *Builder[T] {
	return &Builder[T]{
		s: make([]T, length, capacity),
	}
}

// Seal clears the Builder and returns a Slice with the elements.
//
// All internal state is cleared after calling Seal.
// The Builder can be reused, but should generally not be.
func (b *Builder[T]) Seal() Slice[T] {
	s := b.s
	b.s = nil // important that the new Slice is the sole owner of s!
	return Slice[T]{s: s}
}

func (b *Builder[T]) Append(elem ...T) *Builder[T] {
	b.s = append(b.s, elem...)
	return b
}

// Collect appends all values from seq and returns the builder.
//
// If you know the number of elements in the iter it is usually worthwhile calling Grow before calling Collect,
// or ensure that the builder is created with enough initial capacity when calling NewBuilder.
func (b *Builder[T]) Collect(seq iter.Seq2[int, T]) *Builder[T] {
	for _, elem := range seq {
		b.s = append(b.s, elem)
	}
	return b
}

// Grow ensures there is underlying capacity for appending another n elements without allocations.
func (b *Builder[T]) Grow(n int) *Builder[T] {
	b.s = slices.Grow(b.s, n)
	return b
}

func (b *Builder[T]) Len() int {
	return len(b.s)
}

func (b *Builder[T]) Cap() int {
	return cap(b.s)
}

func (b *Builder[T]) Sort(cmp func(a, b T) int) *Builder[T] {
	slices.SortStableFunc(b.s, cmp)
	return b
}

func (b *Builder[T]) Reverse() *Builder[T] {
	slices.Reverse(b.s)
	return b
}
