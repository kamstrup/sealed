package sealed

type Map[K comparable, V any] struct {
	m map[K]V
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	v, ok := m.m[k]
	return v, ok
}

func (m *Map[K, V]) GetOr(k K, def V) V {
	v, ok := m.m[k]
	if ok {
		return v
	}
	return def
}

func (m *Map[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m *Map[K, V]) All(yield func(K, V) bool) {
	for k, v := range m.m {
		if !yield(k, v) {
			return
		}
	}
}

func (m *Map[K, V]) Len() int {
	return len(m.m)
}

func (m *Map[K, V]) Empty() bool {
	return len(m.m) == 0
}
