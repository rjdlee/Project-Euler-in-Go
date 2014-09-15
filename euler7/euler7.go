package main

import (
	"fmt"
	"math"
	"time"
)

const amount int = 2E5

func main() {
	var input = 10001

	var sqrt1E6 int = int(math.Sqrt(float64(amount)))

	t0 := time.Now()
	numbers := make([]bool, amount)
	primeCount := 0

	// Work in increments of 1E6
	for k := 0; true; k++ {
		for i := 2; i <= sqrt1E6; i++ {
			if numbers[i] == false {
				for j := 2 * i; j < len(numbers); j += i {
					numbers[j] = true
				}
			}
		}
		for j := 0; j < len(numbers); j++ {
			if numbers[j] == false {
				primeCount++
				if primeCount == input+2 {
					fmt.Println(j)
					t1 := time.Now()
					fmt.Println("Runtime: ", t1.Sub(t0))
					return
				}
			}
		}
	}
}
