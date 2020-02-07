package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	limit = int64(600851475143)
	// limit = int64(1000)
)

type message struct {
	Min int64
	Max int64
}

func main() {
	divSet := map[int64]struct{}{}
	start := int64(2)
	min := int64(math.MaxInt64)
	fmt.Printf("Searching divisors from %d to %d (all divisors for %d)\n", start, limit/2, limit)
	for i := int64(2); i <= limit/2; i++ {
		if limit%i == 0 {
			if i < min {
				min = i
			}
			div := limit / i
			divSet[i], divSet[div] = struct{}{}, struct{}{}
			fmt.Printf(">> Adding divisors: %d, %d | (min: %d)\n", i, div, min)
			if div == min {
				break
			}
		}
	}

	div := []int64{}
	for k := range divSet {
		div = append(div, k)
	}
	fmt.Printf("Found divisors: %d total\nSearching prime divisors...\n", len(div))

	var max int64
	for _, d := range div {
		if checkPrime(d) && d > max {
			max = d
			fmt.Printf("Max prime divisor at the moment: %d\n", max)
		}
	}

	fmt.Printf("Result: %d\n", max)
}

func checkPrime(n int64) bool {
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
	var isDivisible bool
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			isDivisible = true
			break
		}
	}
	return !isDivisible
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
