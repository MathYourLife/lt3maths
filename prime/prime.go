package prime

import (
	"math"
)

type PrimeGenerator struct {
	Primes    []uint64
	current   []uint64
	n         uint64
	threshold uint64
}

func NewPrimeGenerator() *PrimeGenerator {
	return &PrimeGenerator{
		Primes:    make([]uint64, 0),
		current:   make([]uint64, 0),
		n:         2,
		threshold: uint64(math.Sqrt(2)),
	}
}

func (p *PrimeGenerator) Next() uint64 {
	if p.n == 2 {
		p.Primes = append(p.Primes, 2)
		p.current = append(p.current, 2)
		p.n++
		return 2
	}

	for {
		t := p.loop_next()
		if t {
			break
		}
	}
	return p.n - 2
}

func (p *PrimeGenerator) loop_next() bool {
	prime := true
CurrentLoop:
	for i, c := range p.current {
		for {
			if c == p.n {
				prime = false
				p.current[i] = c
				break CurrentLoop

			} else if p.Primes[i] > p.threshold {
				break CurrentLoop

			} else if c < p.n {
				c += p.Primes[i]
				continue

			} else {
				p.current[i] = c
				break
			}
		}
	}
	if prime {
		p.threshold = uint64(math.Sqrt(float64(p.n))) + 1
		p.Primes = append(p.Primes, p.n)
		p.current = append(p.current, p.n)
	}
	p.n += 2
	return prime
}
