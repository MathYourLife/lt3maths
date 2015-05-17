package lexperm

import (
	"testing"
)

func TestPermute3(t *testing.T) {
	a := []int{0, 1, 2}
	lp := LexPerm{}
	more := lp.Next(a)
	if !more {
		t.Errorf("There should be more permutations")
	}
	expected := []int{0, 2, 1}
	for i, v := range expected {
		if v != a[i] {
			t.Errorf("Unexpected sequence %v != %v", a, expected)
		}
	}
}

func TestLastPermutation(t *testing.T) {
	a := []int{2, 1, 0}
	lp := LexPerm{}
	more := lp.Next(a)
	if more {
		t.Errorf("There should not be more permutation")
	}
}
