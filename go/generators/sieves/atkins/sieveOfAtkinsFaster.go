/*
        Atkins         10 mill         664579      91.6101ms
 Atkins Github         10 mill         664579      90.5216ms
   Atkins fast         10 mill         664579      64.8269ms
        Atkins        100 mill        5761455     878.6588ms
 Atkins Github        100 mill        5761455     1.3161449s
   Atkins fast        100 mill        5761455      786.908ms
        Atkins       1000 mill       50847576     9.2154671s
 Atkins Github       1000 mill       50847534    13.3827379s
   Atkins fast       1000 mill       50847576     8.9590751s
        Atkins      10000 mill      455052511  3m32.0621299s
 Atkins Github      10000 mill      455052511  6m29.6870127s
   Atkins fast      10000 mill      455052511  4m25.8317892s

*/
package atkins

import (
	"math"
)

func PrimeGen2(primeLimit int) []int {

	sqLimit := int(math.Sqrt(float64(primeLimit)))
	primes := []int{2, 3, 5}
	primeSieve := make([]bool, primeLimit, primeLimit)

	//s := []int{1,7,11,13,17,19,23,29,31,37,41,43,47,49,53,59}
	for x := 1; x < sqLimit; x++ {
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

	for x := 1; x < sqLimit; x += 2 {
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

	for index, prime := range primeSieve {
		if prime {
			primes = append(primes, index)
		}
	}

	return primes
}
