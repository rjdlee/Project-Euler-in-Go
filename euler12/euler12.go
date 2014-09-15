package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t0 := time.Now()
	currTriangle, divisors := 0, 0

	for i := 1; true; i++ {
		divisors = 0
		currTriangle += i
		for j := 1; j <= int(math.Sqrt(float64(currTriangle))); j++ {
			if currTriangle%j == 0 {
				divisors += 2
			}
		}
		if divisors > 500 {
			fmt.Println(currTriangle, " with ", divisors, " factors ")
			t1 := time.Now()
			fmt.Println("Runtime: ", t1.Sub(t0))
			return
		}
	}
}
