package main

import "fmt"

// Embedded type
type Logger struct{}

func (Logger) Log(s string) {
    fmt.Println(s)
}

type Service struct {
    Logger
}

// This is a compile-time check that Service implements an interface
var _ interface{ Log(string) } = Service{}

func main() {
    s := Service{}
    s.Log("Hello from Service")
}
