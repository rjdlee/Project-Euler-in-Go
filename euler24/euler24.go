// https://projecteuler.net/problem=24
//
// Output:
// [2 7 8 3 9 1 5 4 6 0]
// Runtime:  131.02us

package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	t0 := time.Now()
	var comb, combNum float64
	target := float64(1000000) - 1
	allDigits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	digits := []int{}

	for i := len(allDigits) - 1; i >= 0; i-- {
		comb = float64(factorial(i).Int64()) //Get the factorial of i
		combNum = target / comb

		digits = append(digits, allDigits[int(combNum)])
		target = float64(int(target) % int(comb))

		allDigits = append(allDigits[:int(combNum)], allDigits[int(combNum)+1:]...)
	}

	fmt.Println(digits)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}

func factorial(num int) *big.Int {
	factorialNum := big.NewInt(1)
	for i := 2; i <= num; i++ {
		factorialNum.Mul(factorialNum, big.NewInt(int64(i)))
	}
	return factorialNum
}

func indexOf(arr interface{}, val interface{}) int {
	for i := 0; i < len(arr.([]int)); i++ {
		if arr.([]int)[i] == int(val.(int)) {
			return i
		}
	}
	return -1
}
