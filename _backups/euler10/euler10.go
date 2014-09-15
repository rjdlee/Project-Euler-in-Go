package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	numbers := make([]bool, 2E6)
	sum := 0

	// Find all primes less than the squareRoot of 600851475143
	for i := 2; i^2 < len(numbers); i++ {
		if numbers[i] == false {
			for j := 2 * i; j < len(numbers); j += i {
				numbers[j] = true
			}
		}
	}

	// Find the highest prime factor for 600851475143
	for i := len(numbers) - 1; i > 1; i-- {
		if numbers[i] == false {
			sum += i
		}
	}
	fmt.Println(sum)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
