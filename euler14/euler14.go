// https://projecteuler.net/problem=14
//
// Output:
// 837799  with  524
// Runtime:  43.372248ms

package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	var (
		currNum int
		counter int
	)
	chainLen, bigLength, bigIndex := make([]int, 1E6), 0, 0

	for i := 2; i < 1E6; i++ {
		currNum = i
		counter = 0

		for currNum != 1 && currNum >= i {
			counter++
			if currNum%2 == 0 {
				currNum /= 2
			} else {
				currNum = currNum*3 + 1
			}
		}
		counter += chainLen[currNum]
		if counter > bigLength {
			bigLength = counter
			bigIndex = i
		}
		chainLen[i] = counter
	}
	fmt.Println(bigIndex, " with ", bigLength)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
