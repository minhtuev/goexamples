package main

import "fmt"

func main() {
	defer fmt.Println("first defer")
	defer fmt.Println("second defer")
	defer fmt.Println("third defer")

	fmt.Println("main body")
}
