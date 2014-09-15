// https://projecteuler.net/problem=15
//Resources:
// http://copingwithcomputers.com/2013/07/06/lattice-paths/
// Note: Merely was searching for more information about lattice paths
//
// Output:
// 1.3784652882e+12
// Runtime:  83.339us

package main

import (
	"fmt"
	"math"
	"time"
)

const gridSize int = 20

func main() {
	t0 := time.Now()
	// Use Pascal's triangle to find the combinations. Row is twice the size of the grid side length. Index/ Column is the same
	pTRow := gridSize * 2
	pTIndex := gridSize

	paths := factorial(pTRow) / math.Pow(float64(factorial(pTIndex)), 2)

	fmt.Println(paths)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}

func factorial(num int) float64 {
	factorialNum := float64(2)
	for i := float64(3); i < float64(num); i++ {
		factorialNum *= i
	}
	return factorialNum
}
