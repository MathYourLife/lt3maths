# lt3maths

I <3 maths

Some golang packages to aid in maths exploration.

## prime - Prime Number Generator

```go
package main
import(
  "fmt"
  "github.com/mathyourlife/lt3maths/prime"
)
func main () {
  // Initialize the generator
  p := prime.NewPrimeGenerator()
  // Generate 10 primes
  for j := 0; j < 10; j++ {
    fmt.Println(p.Next())
  }
  // Show the current list of primes
  fmt.Println(p.Primes)
}
```

## primefactorization - Calculate the Prime Factorization

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