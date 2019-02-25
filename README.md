# FunctionalGo

This repository holds code snippets showing how to program basic functional ideas using Go.
Whilst it is not recommended to actually use this code in production, it exists to show it is relatively easy to use functional ideas within Go.
This is not necessarily the best implementation of these functional ideas but should be easy to understand. 

## Contents:
- [Closures](#closures)
- [Slices](#slices)

---

### Closures
As Go's functions are first class types, they can be defined inline. Functions declared inline have access to all the variables declared in the current scope, without having to explicitly pass them as arguments.
This allows for basic variable saving such as in the `AnnotatedLogger` function.
```go
func AnnotatedLogger(annotation string) func(string) {
	//Sometimes we only want to store some variables before hand, and allow others to be passed in later
	return func(message string) {
		fmt.Printf("%s: %s", annotation, message)
	}
}
```
This function takes a header and applies returns a function prepends the original header to all messages passed to it.
This could allow you to create different Logging functions for each package in your application, and differentiate between them using the headers

Closures also allow for **Currying**, as this example shows:
```go
func Add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
``` 
Using this add method we can immediately add two numbers by invoking the function twice (e.g. `result := Add(1)(2)`),
or invoke it only once to prepare a function which always adds the same number to it's arguement (e.g. `add5 := Add(5)` where `add5` is a function that takes an int and returns int + 5)

---

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