package main

import (
	"errors"
	"fmt"
	"time"
)

func work(ok bool) (err error) {
	start := time.Now()
	defer func() {
		fmt.Println("took", time.Since(start))
		if err != nil {
			fmt.Println("work failed:", err)
		}
	}()

	if !ok {
		return errors.New("boom")
	}
	time.Sleep(200 * time.Millisecond)
	return nil
}

func main() {
	_ = work(true)
	_ = work(false)
}
