// https://projecteuler.net/problem=29
//
// Solved by hand by finding the maximum combinations minus a^n when
//
// Output:
// Nan

package main

import (
	"fmt"
	"time"
)

const n float64 = 4762

func main() {
	t0 := time.Now()

	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
