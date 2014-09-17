// https://projecteuler.net/problem=25
//
// Manually calculated rather than computer calculated due to poor support for large numbers
//
// Output:
// Nan
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
