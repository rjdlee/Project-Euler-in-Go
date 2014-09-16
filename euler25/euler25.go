// https://projecteuler.net/problem=23
//
// Output:
// 4179871
// Runtime: 552.135279ms
package main

import (
	"fmt"
	"math"
	"time"
)

const n float64 = 4762

func main() {
	t0 := time.Now()

	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	phiPow := math.Pow(phi, n)

	fmt.Println((phiPow - (math.Pow(float64(-1), n) / phiPow)) / sqrt5)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
