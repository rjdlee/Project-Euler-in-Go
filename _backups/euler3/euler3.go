package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t0 := time.Now()
	highestFactor := int(math.Sqrt(600851475143))
	numbers := make([]bool, highestFactor)

	// Find all primes less than the squareRoot of 600851475143
	for i := 2; i^2 < len(numbers); i++ {
		if numbers[i] == false {
			for j := 2 * i; j < len(numbers); j += i {
				numbers[j] = true
			}
		}
	}

	// Find the highest prime factor for 600851475143
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] == false && 600851475143%i == 0 {
			fmt.Println(i)
			break
		}
	}
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
