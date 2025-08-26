package main

import (
    "fmt"
    "time"
    "github.com/expr-lang/expr"
)

func main() {
    // This will fail
    d, err := time.ParseDuration("60d")
    fmt.Println(d, err) // => 0s time: unknown unit "d"

    // This will work
    d2, err := time.ParseDuration("1440h") // 60 * 24
    fmt.Println(d2, err) // => 1440h0m0s <nil>

    env := map[string]interface{}{}

    
    result, err := expr.Eval(`now() - date("2025-06-16") > duration("24h")`, env)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)

}
