/*
Lexicographic Permutations

source: http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order

The following algorithm generates the next permutation lexicographically
after a given permutation. It changes the given permutation in-place.

1) Find the largest index k such that a[k] < a[k + 1].
   If no such index exists, the permutation is the last permutation.
2) Find the largest index l greater than k such that a[k] < a[l].
3) Swap the value of a[k] with that of a[l].
4) Reverse the sequence from a[k + 1] up to and including the
   final element a[n].
*/

package lexperm

import ()

// Lexicographic Permutations
type LexPerm struct{}

/*
Find the largest index k such that a[k] < a[k + 1].
If no such index exists, the permutation is the last permutation.

Return -1 if items are in descending order.
*/
func (lp LexPerm) get_k(a []int) int {
	k := -1
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			k = i
		}
	}
	return k
}

/*
Find the largest index l greater than k such that a[k] < a[l].
*/
func (lp LexPerm) get_l(a []int, k int) int {
	var l int
	for i := k + 1; i < len(a); i++ {
		if a[k] < a[i] {
			l = i
		}
	}
	return l
}

/*
Modify the provided slice in place, and return false if the slice
is already in descending order (no more permutations exist).
*/
func (lp LexPerm) Next(a []int) bool {
	k := lp.get_k(a)
	if k < 0 {
		return false
	}
	l := lp.get_l(a, k)
	// Step 3 swap k and l
	a[k], a[l] = a[l], a[k]
	// Step 4 revers items for elements > k+1
	for i, j := k+1, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return true
}
