package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	sumSquare, squareSum := 0, 0

	// Find all primes less than the squareRoot of 600851475143
	for i := 1; i <= 100; i++ {
		sumSquare += i * i
		squareSum += i
	}
	squareSum = squareSum * squareSum
	fmt.Println(squareSum - sumSquare)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
