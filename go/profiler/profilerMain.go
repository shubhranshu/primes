package main

import (
	"../generators/basic"
	"../generators/sieves/atkins"
	"../generators/sieves/eratosthenes"
	"fmt"
	"math"
	"time"
)

func main() {

	for limitFactor := 3; limitFactor <= 8; limitFactor++ {
		primeLimit := int(math.Pow10(limitFactor))

		start := time.Now()
		primes := basic.PrimeGen(primeLimit)
		elapsed := time.Since(start)
		printStat("Basic", primeLimit, primes, elapsed)

		start = time.Now()
		primes = eratosthenes.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Eratosthenes", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)

		start = time.Now()
		primes = basic.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Basic", primeLimit, primes, elapsed)

		start = time.Now()
		primes = eratosthenes.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Eratosthenes", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)

		start = time.Now()
		primes = basic.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Basic", primeLimit, primes, elapsed)

		start = time.Now()
		primes = eratosthenes.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Eratosthenes", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)

		start = time.Now()
		primes = basic.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Basic", primeLimit, primes, elapsed)

		start = time.Now()
		primes = eratosthenes.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Eratosthenes", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)

	}
}

func printStat(method string, primeLimit int, primes []int, elapsed time.Duration) {
	fmt.Printf("%14s %4.3f mill %14d %14s\n", method, float64(primeLimit)/1000000, len(primes), elapsed)
}
