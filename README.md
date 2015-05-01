# lt3maths

I <3 maths

Some golang packages to aid in maths exploration.

## prime

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