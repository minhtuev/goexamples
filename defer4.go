package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)
		}
	}()

	fmt.Println("about to panic")
	panic("bad thing")
	// not reached
}
