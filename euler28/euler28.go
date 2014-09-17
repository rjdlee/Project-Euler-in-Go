// https://projecteuler.net/problem=28
//
//
// Output:
// 669171001
// Runtime:  71.224us

package main

import (
	"fmt"
	"time"
)

const n float64 = 4762

func main() {
	t0 := time.Now()
	distance := 2 //Distance is the spacing between each corner
	currNum := 1
	sum := 1

	// i is side length
	for i := 1; i < 1001; i += 2 {
		distance = i + 1
		for c := 0; c < 4; c++ {
			currNum += distance
			sum += currNum
		}
	}

	fmt.Println(sum)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
