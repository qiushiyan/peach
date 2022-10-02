map = fn(arr, f) {
    result = []
    for (i in arr) {
        result = append(result, f(i))
    }
    result
}

[1, 2, 3] |> map(fn(x) x * 2)