package prime

import (
	"testing"
)

func TestPrime(t *testing.T) {
	p := NewPrimeGenerator()
	check := []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

	for i := 0; i < len(check); i++ {
		v := p.Next()
		if v != check[i] {
			t.Errorf("%dth prime was %d instead of %d", i+1, v, check[i])
		}
	}
	for i := 0; i < len(check); i++ {
		if p.Primes[i] != check[i] {
			t.Errorf("%dth prime was listed as %d instead of %d", i+1, p.Primes[i], check[i])
		}
	}
}

func Benchmark10Primes(b *testing.B) {
	L := 10
	for i := 0; i < b.N; i++ {
		p := NewPrimeGenerator()
		for j := 0; j < L; j++ {
			p.Next()
		}
	}
}

func Benchmark100Primes(b *testing.B) {
	L := 100
	for i := 0; i < b.N; i++ {
		p := NewPrimeGenerator()
		for j := 0; j < L; j++ {
			p.Next()
		}
	}
}

func Benchmark1000Primes(b *testing.B) {
	L := 1000
	for i := 0; i < b.N; i++ {
		p := NewPrimeGenerator()
		for j := 0; j < L; j++ {
			p.Next()
		}
	}
}

func Benchmark10000Primes(b *testing.B) {
	L := 10000
	for i := 0; i < b.N; i++ {
		p := NewPrimeGenerator()
		for j := 0; j < L; j++ {
			p.Next()
		}
	}
}

func Benchmark100000Primes(b *testing.B) {
	L := 100000
	for i := 0; i < b.N; i++ {
		p := NewPrimeGenerator()
		for j := 0; j < L; j++ {
			p.Next()
		}
	}
}
