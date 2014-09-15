package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	sumSquare, squareSum, sum := 0, 0, 0

	for i := 1; i <= 100; i++ {
		sumSquare += i * i
	}
	for i := 1; i <= 100; i++ {
		sum += i
	}
	squareSum = sum * sum
	fmt.Println(squareSum - sumSquare)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
