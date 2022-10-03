map = fn(arr, f) {
    arr_length = len(arr)
    result = vector(len(arr))
    for (i in 1:arr_length) {
        result[i] = f(arr[i])
    }
    result
}

[1, 2, 3] |> map(fn(x) x * 2)