package main

import (
	"../generators/sieves/atkins"
	"fmt"
	"math"
	"time"
)

func main() {

	for limitFactor := 3; limitFactor <= 8; limitFactor++ {
		primeLimit := int(math.Pow10(limitFactor))

		//results := [3]Results{}
		//println("###################################")
		//start:= time.Now()
		//primes := basic.PrimeGen(primeLimit)
		//elapsed := time.Since(start)
		//printStat("Basic", primeLimit, primes, elapsed)
		//
		//start= time.Now()
		//primes = eratosthenes.PrimeGen(primeLimit)
		//elapsed = time.Since(start)
		//printStat("Eratosthenes", primeLimit, primes, elapsed)


		//start2 := time.Now()
		//primes2 := atkins.PrimeGenGithubExample(primeLimit)
		//elapsed2 := time.Since(start2)
		//printStat("Atkins Github", primeLimit, primes2, elapsed2)

		//Atkins fast         10 mill         664579      96.7412ms
		//Atkins fast        100 mill        5761455     816.8189ms
		//Atkins fast       1000 mill       50847576     8.5501522s


		start := time.Now()
		primes := atkins.PrimeGen(primeLimit)
		elapsed := time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)


		start = time.Now()
		primes = atkins.PrimeGen2(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins fast", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)


		start = time.Now()
		primes = atkins.PrimeGen2(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins fast", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)


		start = time.Now()
		primes = atkins.PrimeGen2(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins fast", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)


		start = time.Now()
		primes = atkins.PrimeGen2(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins fast", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)


		start = time.Now()
		primes = atkins.PrimeGen2(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins fast", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)


		start = time.Now()
		primes = atkins.PrimeGen2(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins fast", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)


		start = time.Now()
		primes = atkins.PrimeGen2(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins fast", primeLimit, primes, elapsed)

		start = time.Now()
		primes = atkins.PrimeGen(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins", primeLimit, primes, elapsed)


		start = time.Now()
		primes = atkins.PrimeGen2(primeLimit)
		elapsed = time.Since(start)
		printStat("Atkins fast", primeLimit, primes, elapsed)




	}
}

type Results struct {
	name    string
	primes  []int
	elapsed time.Duration
}

func printStat(method string, primeLimit int, primes []int, elapsed time.Duration) {
	fmt.Printf("%14s %10d mill %14d %14s\n", method, primeLimit/1000000, len(primes), elapsed)
	//println("First prime : ", primes[0])
	//println("Last prime  : ", primes[len(primes)-1])
	//fmt.Printf("Elapsed     :  %s\n", elapsed)
}
