// https://projecteuler.net/problem=16
//
// Output:
// 1366
// Runtime:  189.094us

package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	t0 := time.Now()
	numString := strconv.FormatFloat(math.Pow(2, 1000), 'f', 0, 64)
	numRune := []rune(numString)
	sum := 0

	for i := 0; i < len(numRune); i++ {
		numba, _ := strconv.Atoi(string(numRune[i]))
		sum += numba
	}

	fmt.Println(sum)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
