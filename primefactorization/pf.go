package primefactorization

import (
	"github.com/mathyourlife/lt3maths/prime"
)

type PrimeFactorization struct {
	cache map[uint64]map[uint64]uint64
}

func NewPrimeFactorization() *PrimeFactorization {

	return &PrimeFactorization{
		cache: make(map[uint64]map[uint64]uint64),
	}
}

func (pf *PrimeFactorization) Of(n uint64) map[uint64]uint64 {

	if len(pf.cache[n]) > 0 {
		return pf.cache[n]
	}

	var t bool

	orig := n

	p := prime.NewPrimeGenerator()
	factor := p.Next()
	f := [][]uint64{}

	use_cache := false
	for {
		if len(pf.cache[n]) > 0 {
			use_cache = true
			fs := map[uint64]uint64{}
			for k, v := range pf.cache[n] {
				fs[k] = v
			}
			for i := len(f) - 1; i >= 0; i-- {
				fs[f[i][1]]++
				pf.cache[f[i][0]] = map[uint64]uint64{}
				for k, v := range fs {
					pf.cache[f[i][0]][k] += v
				}
			}
			break
		}
		t, n = pf.is_a_factor(n, factor)
		if t {
			f = append(f, []uint64{n * factor, factor})
			if n <= 1 {
				break
			}
		} else {
			factor = p.Next()
		}
		if n <= 1 {
			break
		}
	}
	if !use_cache {
		pf.cache_it(f)
	}
	return pf.cache[orig]
}

func (pf *PrimeFactorization) is_a_factor(composite uint64, factor uint64) (bool, uint64) {
	if composite < 2 {
		return false, composite
	}
	if composite%factor == 0 {
		return true, composite / factor
	}
	return false, composite
}

func (pf *PrimeFactorization) cache_it(f [][]uint64) {
	fs := map[uint64]uint64{}
	for i := len(f) - 1; i >= 0; i-- {
		fs[f[i][1]]++
		pf.cache[f[i][0]] = map[uint64]uint64{}
		for k, v := range fs {
			pf.cache[f[i][0]][k] = v
		}
	}
}
