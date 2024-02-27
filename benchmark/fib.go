package main

import (
	"fmt"
	"time"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	start := time.Now()

	fib := fibonacci()
	for i := 0; i < 100000000; i++ {
		fib()
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time for Go (Golang): %s\n", elapsed)
}
