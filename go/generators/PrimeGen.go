package generators

type PrimeGen interface {
	GeneratePrimes(primeLimit int) []int
}
