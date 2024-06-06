// pointers/main.go
package main

import (
    "fmt"
    // "github.com/MICHAELKITH/Golang_full_course/pointer"
)

func main() {
    n := 10
    fmt.Printf("Fibonacci(%d) is = %d \n", n, fibonacci(n))
    
    // 0, 1, 2, 3, 5, 8, 13, 21 ......
}

func fibonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }

    a, b := 0, 1

    for i := 2; i <= n; i++ {
        a, b = b, a+b 
    }

    return b
}
