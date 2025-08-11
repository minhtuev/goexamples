package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "demo-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name()) // delete the file at the end
	defer f.Close()           // close the handle at the end

	if _, err := f.WriteString("hello\n"); err != nil {
		panic(err)
	}
	data, _ := os.ReadFile(f.Name())
	fmt.Printf("file %s has: %q\n", f.Name(), string(data))
}
