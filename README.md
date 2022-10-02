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

Both R and Python has 1-dimensional containers for storing a series of
values. Q offers the vector data structure as created by `[]`

``` q
[1, 2, 3]
```

    #> NumericVector with 3 elements
    #> [1, 2, 3]

A vector is typed by its inner elements. Vectors containing only numbers
numeric vectors, vectors with only string elements are string vectors,
and so on. A vector with mixed types is simply a base `Vector` type, as
is a Python list. No type conversion is done automatically, if the
elements are heterogeneous the base type will be used.

``` q
[1, 2, "hello"]
```

    #> Vector with 3 elements
    #> [1, 2, "hello"]

Vectors in Q have 1-based indexing: the first element starts at index 1,
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

Inspired by R, operators are vectorized and can be applied directly to
vectors.

``` q
print([1, 2, 3] + [4, 5, 6])
print([1, 2, 3] * [4, 5, 6])
```

    #> NumericVector with 3 elements
    #> [5, 7, 9]
    #> NumericVector with 3 elements
    #> [4, 10, 18]

``` q
["hello ", "good "] + ["world", "morning"]
```

    #> CharacterVector with 2 elements
    #> ["hello world", "good morning"]

Vector elements are recycled only if it has lenght 1 or is a scalar of
primitive type.

``` q
print([1, 2, 3] * 2)
print([1, 2, 3] + [4])
```

    #> NumericVector with 3 elements
    #> [2, 4, 6]
    #> NumericVector with 3 elements
    #> [5, 6, 7]

### Dictionaries

You can create a hash table structure in Q called a dictionary with a
pair of `{`, similar to Python, except that you don’t have to quote the
keys.

``` q
property = "functional"
q = {name: "Q", age: 0, property: true}
print(q)
print(keys(q))
print(values(q))
```

    #> {"functional": true, "name": "Q", "age": 0}
    #> CharacterVector with 3 elements
    #> ["age", "functional", "name"]
    #> Vector with 3 elements
    #> ["Q", 0, true]

### Control flows

``` q
for (name in ["Q", "R", "Python"]) {
  if (name != "Python") {
    print(Python)
  }
}
```

    #> [34mERROR: identifier not found: Python[0m

### Functions

Functions in Q are first-class citizens. They can be passed around as
arguments and returned from other functions. There is a `return` keyword
but functions can also use implicit returns.

``` markdown
map = fn(arr, f) {
  result = []
  for (x in arr) {
    result = append(result, f(x))
  }
  result
}

[1, 2, 3] |> map(fn(x) { x * 2 })
```

A more advanced example

``` markdown
let reduce = fn(f, seed, arr) {
    let iter = fn(acc, arr) {
        if (len(arr) == 0) {
            return acc
        }
        iter(f(acc, head(arr)), tail(arr))
    }

    iter(seed, arr)
}

let sum = fn(arr) {
    reduce(fn(acc, it) { acc + it }, 0, arr)
}

[1, 2, 3, 4, 5] |> sum
```

## Next steps

- `...` for variadic arguments

- fix `append()` to use copy

- dataframe interface

- improve error message with token col and line

- more standard library functions
