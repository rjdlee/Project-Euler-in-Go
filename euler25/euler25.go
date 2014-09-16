// https://projecteuler.net/problem=23
//
// Output:
// 4179871
// Runtime:  552.135279ms

package main

import (
	"fmt"
	"math/big"
	"time"
)

// Limit can be lowered from 28123 to 20161
const Phi int = 20161 + 1

func main() {
	t0 := time.Now()

	fmt.Println(fib(big.NewInt(int64(5))).String())
	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}

func fib(n *big.Int) *big.Int {
	one, two := big.NewInt(int64(1)), big.NewInt(int64(2))

	if n.Cmp(two) <= 0 {
		return one
	}
	k := n.Div(n, two)
	a := fib(k.Add(k, two))
	b := fib(k)
	var mod, total *big.Int
	n.DivMod(n, two, mod)
	if mod.Cmp(one) == 0 {
		return total.Add(total.Mul(a, a), total.Mul(b, b))
	} else {
		return total.Mul(b, total.Sub(total.Mul(two, a), b))
	}
}
