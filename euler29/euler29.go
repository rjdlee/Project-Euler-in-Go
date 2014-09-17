// https://projecteuler.net/problem=28
//
// Solved by hand
//
// Output:
// 669171001
// Runtime:  71.224us

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
