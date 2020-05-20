/*
   Atkins fast         10 mill         664579      64.8269ms
   Atkins fast        100 mill        5761455      786.908ms
   Atkins fast       1000 mill       50847576     8.9590751s
   Atkins fast      10000 mill      455052511  4m25.8317892s
*/
package atkins

import (
	"math"
)

func PrimeGen(primeLimit int) []int {

	sqLimit := int(math.Sqrt(float64(primeLimit)))
	primeSieve := make([]bool, primeLimit, primeLimit)

	for x := 1; x <= sqLimit/2; x++ {
		for y := 1; y < sqLimit; y += 2 {
			n := 4*x*x + y*y
			if n >= primeLimit {
				continue
			}
			mod := n % 60
			if mod == 1 || mod == 13 || mod == 17 || mod == 29 || mod == 37 || mod == 41 || mod == 49 || mod == 53 {
				primeSieve[n] = !primeSieve[n]
			}
		}
	}

	for x := 1; x < int(float64(sqLimit)/1.6); x += 2 {
		for y := 2; y < sqLimit; y += 2 {
			n := 3*x*x + y*y
			if n >= primeLimit {
				continue
			}
			mod := n % 60
			if mod == 7 || mod == 19 || mod == 31 || mod == 43 {
				primeSieve[n] = !primeSieve[n]
			}
		}
	}

	for x := 2; x < sqLimit; x += 1 {
		for y := 1; y < sqLimit && y <= x; y += 2 {
			n := 3*x*x - (x-y)*(x-y)
			if n >= primeLimit {
				continue
			}

			mod := n % 60
			if mod == 11 || mod == 23 || mod == 47 || mod == 59 {
				primeSieve[n] = !primeSieve[n]
			}
		}
	}

	for n := 0; n < sqLimit; n++ {
		if primeSieve[n] {
			for y := n * n; y < primeLimit; y += n * n {
				primeSieve[y] = false
			}
		}
	}

	// optimistically allocate memory
	primes := make([]int, 0, primeLimit/20)
	primes = append(primes, 2)
	primes = append(primes, 3)
	primes = append(primes, 5)
	for index, prime := range primeSieve {
		if prime {
			primes = append(primes, index)
		}
	}

	return primes
}
