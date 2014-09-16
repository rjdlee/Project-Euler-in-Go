// https://projecteuler.net/problem=22
//
// Output:
// 871198282
// Runtime:  9.460917ms

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	t0 := time.Now()

	alphabet := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9, "J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17, "R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25, "Z": 26}
	var nameRune []rune
	score, totalScore := 0, 0

	namesString, _ := readLines("names.txt")
	names := strings.Split(strings.Replace(namesString[0], `"`, "", -1), `,`) // Parse the file
	sort.Strings(names)                                                       // Alphabetical order

	for i := 0; i < len(names); i++ {
		nameRune = []rune(names[i])
		score = 0
		for j := 0; j < len(nameRune); j++ {
			score += alphabet[string(nameRune[j])]
		}
		score *= (i + 1)
		totalScore += score
	}

	fmt.Println(totalScore)
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
