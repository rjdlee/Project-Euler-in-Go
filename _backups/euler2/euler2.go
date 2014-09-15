package main

import "fmt"

func main() {
	lastTwo := []int{1, 1}
	sum, counter := 0, 2

	for currNum := 2; currNum < 4E6; currNum += 3 {
		// Add the last to numbers to form the new one
		currNum = lastTwo[0] + lastTwo[1]
		lastTwo[0] = lastTwo[1]
		lastTwo[1] = currNum

		// Every third is an even number
		if counter == 2 {
			counter = 0
			sum += currNum
		} else {
			counter++
		}
	}
	fmt.Println(sum)
}
