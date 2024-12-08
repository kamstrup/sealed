package sealed

import (
	"iter"
	"maps"
)

type Mapper[K comparable, V any] struct {
	m map[K]V
}

func (m *Mapper[K, V]) Seal() Map[K, V] {
	a := m.m
	m.m = nil // important that Map is the sole owner of m.m!
	return Map[K, V]{m: a}
}

func (m *Mapper[K, V]) Put(k K, v V) *Mapper[K, V] {
	m.m[k] = v
	return m
}

func (m *Mapper[K, V]) Copy(src map[K]V) *Mapper[K, V] {
	maps.Copy(m.m, src)
	return m
}

func (m *Mapper[K, V]) Collect(seq iter.Seq2[K, V]) *Mapper[K, V] {
	for k, v := range seq {
		m.m[k] = v
	}
	return m
}

func (m *Mapper[K, V]) Len() int {
	return len(m.m)
}
