package main

import (
    "fmt"
    "strconv"
)

func main() {
    s := "123"

    // string to int
    i, err := strconv.Atoi(s,)
    if err != nil {
        // ... handle error
        panic(err)
    }

    fmt.Println(s, i)
}