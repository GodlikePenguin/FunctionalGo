# FunctionalGo

This repository holds code snippets showing how to program basic functional ideas using Go.
Whilst it is not recommended to actually use this code in production, it exists to show it is relatively easy to use functional ideas within Go.
This is not necessarily the best implementation of these functional ideas but should be easy to understand. 

## Contents:
- [Slices](#slices)

### Slices
The `slices` package in this repository contains functional types on slices which implement the `map`, `filter` and `reduce` functional constructs.
The types in this package provide `From` and `To` methods for their respective basic types to allow you to take an existing slice, perform some functional operations on it, and then get a basic slice result.

Example **Map** usage: square all numbers in a slice

```go
i := slices.FromIntSlice([]int{1, 2, 3})
result := i.Map(func(i int) int {
    return i * i
}).ToSlice()
// result = [1 4 9]
```

Example **Filter** usage: remove numbers less than 2

```go
i := slices.FromIntSlice([]int{1, 2, 3})
result := i.Filter(func(i int) bool {
    return i > 1
}).ToSlice()
// result = [2 3]
```

Example **Reduce** usage: sum the numbers in a slice
```go
i := slices.FromIntSlice([]int{1, 2, 3})
result := i.Reduce(func(acc int, i int) int {
    return acc + i
}, 0)
// result = 6
```

The `map` and `filter` operations return Functional objects, so the operations can be chained together. `Reduce` returns a single item, so must be the last function call in the chain.
More examples are available in the `examples.go` file in the package.

There is currently a lot of duplicated code in the slices package, however without generics and with Go's strict typing of functions (`func(string) != func(int)`) there is no trivial solution to reducing the amount of code.