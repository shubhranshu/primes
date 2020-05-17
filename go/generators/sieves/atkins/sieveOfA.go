package main

import (
	"../../../helpers"
	"math"
	"time"
)

const primeLimit = 100000000

func main() {

	defer helpers.TimeTrack(time.Now(), "Basic primes")
	sqLimit := int(math.Sqrt(primeLimit))
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

	println(len(primes))
}
