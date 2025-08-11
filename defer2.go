package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println("arg captured at defer time:", i) // prints 0
	defer func() { fmt.Println("closure sees final i:", i) }() // prints 42

	i = 42
	fmt.Println("set i to", i)
}
