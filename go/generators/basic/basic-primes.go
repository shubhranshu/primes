/*
Generate the fastest primes < 10,000,000

Total primes  664579
Basic primes : 2.5882856s

*/

package main

import (
	"../../helpers"
	"math"
	"time"
)

const primeLimit = 10000000

func main() {
	defer helpers.TimeTrack(time.Now(), "Basic primes")
	primes := make([]int, 0, primeLimit/10)
	primes = append(primes, 2)
	for i := 3; i < primeLimit; i = i + 2 {
		sqPrime := int(math.Sqrt(float64(i)))
		for _, prime := range primes {
			if prime > sqPrime {
				primes = append(primes, i)
				break
			}
			if i%prime == 0 { //not prime
				break
			}
		}
	}

	println("Total primes ", len(primes))
	//for _,prime := range primes {
	//	println(prime)
	//}
}
