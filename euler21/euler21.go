// https://projecteuler.net/problem=21
//
// Output:
// 31626
// Runtime:  583.833us

package main

import (
	"fmt"
	"time"
)

const maxNum int = 10000

func main() {
	t0 := time.Now()

	sums := make([]int, 10000) // Hold the sum of each number for faster calculation
	nums := make([]bool, 10000)
	numsAdded := make([]bool, 10000) // Prevents double adding to totalSum
	totalSum, match := 0, 0

	// Generate a list of the divisors for numbers 0-10000
	for i := 1; i < maxNum/2; i++ {
		for j := i * 2; j < maxNum; j += i {
			if nums[j] == false && i > 1 {
				nums[j] = true
			}
			sums[j] += i
		}
	}
	// Determine if a pair exists while preventing a number from matching up with itself or a pair from being added twice
	for j := 0; j < maxNum; j++ {
		if sums[j] < maxNum && numsAdded[j] == false {
			match = sums[sums[j]]
			if j == match && j != sums[j] {
				numsAdded[j] = true
				numsAdded[sums[j]] = true
				totalSum += j + sums[j]
			}
		}
	}
	fmt.Println(totalSum)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
