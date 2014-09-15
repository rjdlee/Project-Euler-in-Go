package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t0 := time.Now()
	var c int

	for b := 0; b < 500; b++ {
		for a := b + 1; a < 500; a++ {
			c = 1000 - b - a
			if math.Pow(float64(b), 2)+math.Pow(float64(a), 2) == math.Pow(float64(c), 2) {
				fmt.Println(c * a * b)
			}
		}
	}
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
