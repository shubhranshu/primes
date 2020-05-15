/*
Generate the fastest primes < 10,000,000

Basic
Total primes  664579
Basic primes : 2.5882856s

This

Primes  664579
First prime  2
Last prime  9999991
Basic primes : 255.3159ms

100 million
Primes  5761455
First prime  2
Last prime  99999989
Basic primes : 2.4943548s



*/

package main

import (
	"../../../helpers"
	"time"
)

const primeLimit = 100000000

func main() {

	defer helpers.TimeTrack(time.Now(), "Basic primes")

	primeSieve := make([]int, primeLimit, primeLimit)
	primeSieve[2] = 0
	for i := 2; i < primeLimit; i++ {
		if primeSieve[i] == -1 {
			continue
		}
		for j := 2 * i; j < primeLimit; j = j + i {
			primeSieve[j] = -1
		}
	}
	primeSlice := make([]int, 0, primeLimit/10)
	for index, val := range primeSieve {
		if val != -1 {
			primeSlice = append(primeSlice, index)
		}
	}
	println("Primes ", len(primeSlice[2:]))
	println("First prime ", primeSlice[2])
	println("Last prime ", primeSlice[len(primeSlice)-1])
}
