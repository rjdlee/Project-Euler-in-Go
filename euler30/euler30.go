// https://projecteuler.net/problem=30
//
// Output:
// 443839
// Runtime:  247.649052ms

package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const powNum float64 = 5

func main() {
	t0 := time.Now()

	max := int(math.Pow(9, powNum)*powNum - 1)

	pow := []int{0, 1}
	var num []byte
	digit, numSum, sum, numInt := 0, 0, -1, 0

	for i := 2; i < 10; i++ {
		pow = append(pow, int(math.Pow(float64(i), powNum)))
	}

	for i := 0; i < max; i++ {
		num = []byte(strconv.FormatInt(int64(i), 10))
		numSum = 0
		for k := 0; k < len(num); k++ {
			digit, _ = strconv.Atoi(string(num[k]))
			numSum += pow[digit]
		}
		numInt, _ = strconv.Atoi(string(num))
		if numInt == numSum {
			sum += numSum
		}
	}

	fmt.Println(sum)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
