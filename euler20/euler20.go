// https://projecteuler.net/problem=20
//
// Note: 100! is not calculable with default numeric types, so the math/big package was used.
//
// Output:
// 648
// Runtime: 152.453us

package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

const gridSize int = 20

func main() {
	t0 := time.Now()

	num, sum := factorial(100), 0
	runes := []rune(num.String())

	for i := 0; i < len(runes); i++ {
		j, _ := strconv.Atoi(string(runes[i]))
		sum += j
	}

	fmt.Println(sum)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}

func factorial(num int) *big.Int {
	factorialNum := big.NewInt(2)
	for i := 3; i <= num; i++ {
		factorialNum.Mul(factorialNum, big.NewInt(int64(i)))
	}
	return factorialNum
}
