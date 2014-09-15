package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	t0 := time.Now()
	number := big.NewInt(0)
	biggestNum := big.NewInt(0)

	// Find all primes less than the squareRoot of 600851475143
	for i := int64(999); i > 0; i-- {
		for j := int64(999); j > 0; j-- {
			number = big.NewInt(i * j)
			if Reverse(number.String()) == number.String() {
				if number.Int64() > biggestNum.Int64() {
					biggestNum = number
				}
			}
		}
	}
	fmt.Println(biggestNum)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
