package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	var numbers []int
	numbersTemp := make([]bool, 20)
	divisible := 0

	// Find all primes less than the squareRoot of 600851475143
	for i := len(numbersTemp) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if numbersTemp[j] == false {
				if (i+1)%(j+1) == 0 {
					numbersTemp[j] = true
				}
			}
		}
	}

	for i := len(numbersTemp) - 1; i >= 0; i-- {
		if numbersTemp[i] == false {
			numbers = append(numbers, i+1)
		}
	}

	for i := 2 * numbers[0]; true; i += numbers[0] {
		divisible = 0
		for j := 1; j < len(numbers); j++ {
			if i%numbers[j] == 0 {
				divisible++
			}
		}
		if divisible == len(numbers)-1 {
			fmt.Println(i)
			t1 := time.Now()
			fmt.Println("Runtime: ", t1.Sub(t0))
			return
		}
	}

}
