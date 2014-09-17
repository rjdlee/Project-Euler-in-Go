// https://projecteuler.net/problem=25
//
// Calculated by hand using: http://math.fau.edu/richman/Liberal/decimal.htm and https://en.wikipedia.org/wiki/List_of_prime_numbers#The_first_500_prime_numbers
// Since https://en.wikipedia.org/wiki/Repeating_decimal, the largest repeating decimal with d(denominator) under 1000 we can have is d - 1. The rules says 1/d produces this.
//
// Output:
// Nan
package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
