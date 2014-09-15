// https://projecteuler.net/problem=17
//
// Output:
// 21124
// Runtime:  247.145us

package main

import (
	"fmt"
	"time"
)

const gridSize int = 20

func main() {
	t0 := time.Now()
	// Use Pascal's triangle to find the combinations. Row is twice the size of the grid side length. Index/ Column is the same
	numWord := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty", 30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety"}
	sum := 0

	// h: hundreds, t: tens, o: ones
	for h := 0; h < 1000; h += 100 {
		for t := 0; t < 100; t += 10 {
			for o := 0; o < 10; o++ {
				if h > 0 {
					sum += len(numWord[h/100]) + len("hundred")
					if o+t > 0 {
						sum += len("and")
					}
				}
				if (o+t) < 20 && (o+t) > 0 {
					sum += len(numWord[t+o])
				} else {
					if t >= 20 {
						sum += len(numWord[t])
					}
					if o > 0 {
						sum += len(numWord[o])
					}
				}
			}
		}
	}

	fmt.Println(sum + len("onethousand"))
	//fmt.Println(strconv.FormatInt(5, 10))
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
