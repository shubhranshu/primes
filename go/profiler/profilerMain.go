package main

import (
	"../generators/sieves/atkins"
	"fmt"
	"math"
	"time"
)

func main() {

	for limitFactor := 7; limitFactor <= 10; limitFactor++ {
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

		start1 := time.Now()
		primes1 := atkins.PrimeGen(primeLimit)
		elapsed1 := time.Since(start1)
		printStat("Atkins", primeLimit, primes1, elapsed1)

		start2 := time.Now()
		primes2 := atkins.PrimeGenGithubExample(primeLimit)
		elapsed2 := time.Since(start2)
		printStat("Atkins Github", primeLimit, primes2, elapsed2)

		start3 := time.Now()
		primes3 := atkins.PrimeGen2(primeLimit)
		elapsed3 := time.Since(start3)
		printStat("Atkins fast", primeLimit, primes3, elapsed3)

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
