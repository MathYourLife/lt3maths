# primefactorization

Calculate the prime factorization for a given value.

Uses: `github.com/mathyourlife/lt3maths/prime`

## Usage

```go
package main
import (
  "fmt"
  "github.com/mathyourlife/lt3maths/primefactorization"
)
func main() {
  pf := primefactorization.NewPrimeFactorization()
  result := pf.Of(12)
  fmt.Println(result)
  // map[2:2 3:1]
}
```

## Performance Notes

The `PrimeFactorization` instance will cache the final and intermediate
prime factorizations for each call, so reusing an existing instance for
multiple factorizations will result in higher cache hit ratios.

## Sample

```
Calculate the prime factorization of 5742
```

1) `github.com/mathyourlife/lt3maths/prime` is used to provide a list of
    primes to determine if they are divisors.
2) **2** is found to be a factor of 5742 with the remainder 2871.
3) **3** is found to be a factor of 2871 twice with the remainder 319.
4) **11** is found to be a factor of 319 with the remainder 29.
5) **29** is determined to be prime.
6) The prime factorization is returned `5742 = map[11:1 3:2 2:1 29:1]`

In this case the cache now includes the returned map for 5742 as well as the
intermediate prime factorizations.

```
map[
    29:map[29:1]
    319:map[29:1 11:1]
    957:map[29:1 11:1 3:1]
    2871:map[29:1 11:1 3:2]
    5742:map[29:1 11:1 3:2 2:1]
]
```

## Benchmarking

```
$ go test -bench=.
PASS
BenchmarkPow2_50       20000         81394 ns/op
Benchmark99991           200       6381317 ns/op
Benchmark149358       100000         18359 ns/op
ok      _/home/dcouture/git/mathyourlife/lt3maths/primefactorization    6.422s
```