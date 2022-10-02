Q Progrmaming Language
================

Q is a toy programming language with a mix of R and Python’s syntax. It
was written in Go and inspired by <https://interpreterbook.com/>.

## Data structures

### Primitives

Primitive data structures include: numbers (integers and floats),
strings and booleans

``` q
1 + 1 + (10 * 2) / 4
```

    #> 7

``` q
"hello" + " " + "world"
```

    #> "hello world"

``` q
!false
```

    #> true

### Vectors

Similar to R Q has vectors typed by its inner elements. Vectors with
number elements are numeric vectors, vectors with string elements are
string vectors, and so on. A vector with mixed types is simply a base
`Vector` type.

``` q
[1, 2, 3, 4, 5]
```

    #> NumericVector with 5 elements
    #> [1, 2, 3, 4, 5]

``` q
["hello", "world"]
```

    #> CharacterVector with 2 elements
    #> ["hello", "world"]

Vectors in Q has 1-based indexing, so the first element is at index 1,
not 0. Built-in functions for vectors include `len()`, `append()`,
`head()`, `tail()`

``` q
x = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

print(x[1:3])
print(append(x, [11, 12, 13], 14, "15"))
print(head(x, 10))
```

    #> NumericVector with 3 elements
    #> [1, 2, 3]
    #> Vector with 15 elements
    #> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, "15"]
    #> NumericVector with 10 elements
    #> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

### Dictionaries

The hash table structure in Q is called a dictionary similar in Python,
except that you don’t have to quote the keys.

``` q
property = "functional"
q = {name: "Q", "age": 0, property: true}
q
```

    #> {"name": "Q", "age": 0, "functional": true}

### Control flows

### Functions

Functions in Q are first-class citizens. They can be passed around as
arguments and returned from other functions. There is a `return` keyword
but functions can also use implicit returns.

``` q
let fib = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1
    } else {
      fib(x - 1) + fib(x - 2)
    }
  }
}
fib(10)
```

    #> [34mERROR: identifier not found: fib[0m

## Next steps

- `...` for variadic arguments

- for loops

- dataframe interface

- improve error message with token col and line

- more standard library functions
