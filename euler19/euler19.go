// https://projecteuler.net/problem=19
//
// Output:
// 171
// Runtime:  139.525us

package main

import (
	"fmt"
	"time"
)

const gridSize int = 20

func main() {
	t0 := time.Now()
	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	totalDays, totalSun := 0, 0

	for year := 1901; year < 2000; year++ {
		// Leap year
		if year%4 == 0 {
			monthDays[1] = 29
		}
		for m := 0; m < 12; m++ {
			if (totalDays+6+1)%7 == 0 {
				totalSun++
			}
			totalDays += monthDays[m]
		}
	}

	fmt.Println(totalSun)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
