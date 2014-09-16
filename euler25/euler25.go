// https://projecteuler.net/problem=23
//
// Output:
// 4179871
// Runtime:  552.135279ms

package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

const n float64 = 4782
const decimalPlaces float64 = 10000

func main() {
	t0 := time.Now()

	phi := big.NewInt(int64(math.Floor(math.Phi * decimalPlaces)))
	phi.Exp(phi, big.NewInt(int64(n)), big.NewInt(0))

	sqrt5 := big.NewInt(int64(math.Floor(math.Sqrt(5) * decimalPlaces)))

	num := big.NewInt(int64(math.Pow(-1*10000, n))) // (-1)^n
	num.Div(num, phi)
	num.Sub(phi, num)
	num.Div(num, sqrt5)
	num.Div(num, big.NewInt(int64(math.Pow(10000, n*n))))
	fmt.Println(num)

	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
