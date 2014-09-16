// https://projecteuler.net/problem=23
//
// Output:
// 4179871
// Runtime:  552.135279ms

package main

import (
	"fmt"
	"time"
)

// Limit can be lowered from 28123 to 20161
const maxNum int = 20161 + 1

func main() {
	t0 := time.Now()
	totalSum, sums := 0, make([]int, maxNum)
	abundantNums := []int{}
	sumAbunNums := map[int]bool{}

	// Find the sum of the proper divisors from 0 to the maxNum
	for i := 1; i < maxNum/2; i++ {
		for j := i * 2; j < maxNum; j += i {
			sums[j] += i
		}
	}

	// Find and store all the abundant numbers
	for i := 0; i < maxNum; i++ {
		if i < sums[i] {
			abundantNums = append(abundantNums, i)
		}
	}

	// Find the sums of all abundant numbers
	for i := 0; i < len(abundantNums); i++ {
		for j := i; j < len(abundantNums); j++ {
			if abundantNums[i]+abundantNums[j] < maxNum {
				sumAbunNums[abundantNums[i]+abundantNums[j]] = true
			}
		}

	}

	for i := 0; i < maxNum; i++ {
		if !sumAbunNums[i] {
			totalSum += i
		}
	}

	fmt.Println(totalSum)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
