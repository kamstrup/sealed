package sealed_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/kamstrup/sealed"
)

func TestSlice(t *testing.T) {
	b := sealed.Builder[int]{}
	b.Append(2, 1, 3).
		AppendSeq2(slices.All([]int{6, 5, 4})).
		Sort(cmp.Compare[int])

	if b.Len() != 6 {
		t.Fatalf("expected %d elements in builder, found %d", 6, b.Len())
	}

	s := b.Seal()
	if s.Len() != 6 {
		t.Fatalf("expected %d elements in slice, found %d", 6, s.Len())
	}

	idx := 0
	for i, v := range s.All {
		if i != idx {
			t.Fatalf("expected index %d , found %d", idx, i)
		}
		if v != i+1 {
			t.Fatalf("expected value %d, at index %d, found %d", i+1, i, v)
		}
		idx++
	}

	for i := 0; i < s.Len(); i++ {
		if s.Get(i) != i+1 {
			t.Fatalf("expected value %d, at index %d, found %d", i+1, i, s.Get(i))
		}
	}
}
