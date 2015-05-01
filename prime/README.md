# prime

Prime number generator.

## Usage

Initialize the generator.

```golang
p := NewPrimeGenerator()
```

Make some primes

```golang
for i = 0; i < 10; i++ {
  fmt.Println(p.Next())
}
```

Show the current list of primes

```golang
fmt.Println(p.Primes)
```

## Benchmarking

```
$ go test -bench=.
PASS
Benchmark10Primes         500000          3442 ns/op
Benchmark100Primes         50000         31121 ns/op
Benchmark1000Primes         5000        400247 ns/op
Benchmark10000Primes         200       6500074 ns/op
Benchmark100000Primes         10     114879051 ns/op
ok      _/home/dcouture/git/mathyourlife/lt3maths/prime 8.925s
```