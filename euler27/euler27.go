// https://projecteuler.net/problem=25
//
// Go implementation of http://www.mathblog.dk/project-euler-27-quadratic-formula-primes-consecutive-values/
// Can be solved easier using: n^2-(2p-1)n+p^2-p+41
// Where |2p-1| < 1000 and |p^2-p+41| < 1000
// http://mathworld.wolfram.com/Prime-GeneratingPolynomial.html
//
// Output:
// Nan
package main

import (
	"fmt"
	"math"
	"time"
)

const n float64 = 4762

func main() {
	t0 := time.Now()

	aMax, bMax, nMax := 0, 0, 0
	primes := make([]bool, 90000)
	bPos := make([]bool, 1000)

	// Find prime numbers for primes
	for i := 2; i^2 < len(primes); i++ {
		if primes[i] == false {
			for j := 2 * i; j < len(primes); j += i {
				primes[j] = true
			}
		}
	}

	// Find prime numbers for bPos
	for i := 2; i^2 < len(bPos); i++ {
		if bPos[i] == false {
			for j := 2 * i; j < len(bPos); j += i {
				bPos[j] = true
			}
		}
	}

	for a := -999; a < 1000; a += 2 {
		for i := 1; i < len(bPos); i++ {
			if bPos[i] == false {
				for j := 1; j < 2; j++ {
					n := 0
					sign := -1
					if j == 0 {
						sign = 1
					}
					aodd := 0
					if i%2 == 0 {
						aodd = -1
					}
					prime := int(math.Abs(float64(n*n + (a+aodd)*n + sign*i)))

					if primes[k] == false && k == prime {
						n++
					}
					/*for k := 0; k <= prime; k++ {
						if primes[k] == false && k == prime {
							n++
						}
					}*/
					if n > nMax {
						aMax = a
						bMax = i
						nMax = n
					}
				}
			}
		}
	}

	fmt.Println(aMax, bMax, nMax)

	/*aMax, bMax = 0, nMax = 0;
	primes = ESieve(87400);
	int[] bPos = ESieve(1000);

	for (int a = -999; a < 1001; a +=2) {
	    for (int i = 1; i < bPos.Length; i++) {
	        for (int j = 0; j < 2; j++) {
	            int n = 0;
	            int sign = (j == 0) ? 1 : -1;
	            int aodd = (i % 2 == 0) ? -1 : 0; // Making a even if b is even
	            while (isPrime(Math.Abs(n * n + (a + aodd) * n + sign*bPos[i]))) {
	                n++;
	            }

	            if (n > nMax) {
	                aMax = a;
	                bMax = bPos[i];
	                nMax = n;
	            }
	        }
	    }
	}*/

	t1 := time.Now()
	fmt.Println("Runtime: ", t1.Sub(t0))
}
