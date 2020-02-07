package main

import (
	"fmt"
	"strconv"
)

const (
	limit = int64(600851475143)
	// limit = int64(1000)
)

func main() {
	primes := []int64{}
	for i := int64(0); i <= limit; i++ {
		if isPrime(primes, i) {
			primes = append(primes, i)
		}
	}

	gpd := greatestPrimeDivisor(primes, limit)
	fmt.Printf("Primes: %v", primes)
	fmt.Printf("Result: %d", gpd)
}

func isPrime(primes []int64, n int64) bool {
	if n == 0 || n == 1 {
		return false
	}
	v := fmt.Sprintf("%d", n)
	var digitSum int64
	for _, d := range v {
		digit, err := strconv.ParseInt(string(d), 10, 64)
		if err != nil {
			panic(err)
		}
		digitSum += digit
	}
	if digitSum%3 == 0 {
		return false
	}

	return greatestPrimeDivisor(primes, n) == 0
}

// greatestPrimeDivisor returns the greatest prime number that divides n; if
// n == 0 then none number in list divides n.
func greatestPrimeDivisor(primes []int64, n int64) int64 {
	input := make(chan int64, len(primes))
	output := make(chan int64, len(primes))

	for i := 0; i < len(primes); i++ {
		go func() {
			p := <-input
			if n%p == 0 {
				output <- p
			}
			output <- 0
		}()
	}

	for _, p := range primes {
		input <- p
	}
	close(input)

	var greatest int64
	for i := 0; i < len(primes); i++ {
		if v := <-output; v > greatest {
			greatest = v
		}
	}
	return greatest

}
