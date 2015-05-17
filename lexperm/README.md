# lexperm

Lexicographic permutation generator.

Given an array of integers in weakly increasing order (including the
possibility for repeated values).  `LexPerm` will generate all
permutations in a lexicographic order.

source: http://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order

## Usage

```go
package main
import(
  "fmt"
  "github.com/mathyourlife/lt3maths/lexperm"
)
func main () {
  lp := lexperm.LexPerm{}
  a := []int{0, 1, 2}
  for {
    fmt.Println(a)
    more := lp.Next(a)
    if !more {
      break
    }
  }
}
```
