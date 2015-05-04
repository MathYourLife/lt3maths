package primefactorization

import (
	"testing"
)

func match(result map[uint64]uint64, expected map[uint64]uint64) bool {
	if len(expected) != len(result) {
		return false
	}
	for k, v := range expected {
		if result[k] != v {
			return false
		}
	}
	return true
}

func TestPrimeFactorization(t *testing.T) {

	pf := NewPrimeFactorization()

	result := pf.Of(12)
	expected := map[uint64]uint64{
		2: 2,
		3: 1,
	}
	if !match(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}

	result = pf.Of(450)
	expected = map[uint64]uint64{
		2: 1,
		3: 2,
		5: 2,
	}
	if !match(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}

	result = pf.Of(9941)
	expected = map[uint64]uint64{
		9941: 1,
	}
	if !match(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func TestNoFactorization(t *testing.T) {

	pf := NewPrimeFactorization()

	result := pf.Of(1)
	expected := map[uint64]uint64{}
	if !match(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}

	result = pf.Of(0)
	expected = map[uint64]uint64{}
	if !match(result, expected) {
		t.Errorf("%v != %v", expected, result)
	}
}

func BenchmarkPow2_50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pf := NewPrimeFactorization()
		pf.Of(1125899906842624)
	}
}

func Benchmark99991(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pf := NewPrimeFactorization()
		pf.Of(99991)
	}
}

func Benchmark149358(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pf := NewPrimeFactorization()
		pf.Of(149358)
	}
}
