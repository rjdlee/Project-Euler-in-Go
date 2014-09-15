// https://projecteuler.net/problem=18
//
// Output:
// 1074
// Runtime:  159.064us

package main

import (
	"fmt"
	"strconv"
	"time"
)

const gridSize int = 20

func main() {
	t0 := time.Now()
	// Use Pascal's triangle to find the combinations. Row is twice the size of the grid side length. Index/ Column is the same
	triangle := [][]float64{
		[]float64{75},
		[]float64{95, 64},
		[]float64{17, 47, 82},
		[]float64{18, 35, 87, 10},
		[]float64{20, 4, 82, 47, 65},
		[]float64{19, 1, 23, 75, 3, 34},
		[]float64{88, 2, 77, 73, 7, 63, 67},
		[]float64{99, 65, 4, 28, 6, 16, 70, 92},
		[]float64{41, 41, 26, 56, 83, 40, 80, 70, 33},
		[]float64{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		[]float64{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		[]float64{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		[]float64{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		[]float64{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		[]float64{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 2}}

	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			if triangle[row+1][col] > triangle[row+1][col+1] {
				triangle[row][col] += triangle[row+1][col]
			} else {
				triangle[row][col] += triangle[row+1][col+1]
			}
		}
	}

	fmt.Println(strconv.FormatFloat(triangle[0][0], 'f', 0, 64))

	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
