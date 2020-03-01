package problem3

import (
	"fmt"
	"math"

	"github.com/ariel17/project-euler/tools"
)

const (
	limit = int64(600851475143)
	// limit = int64(1000)
)

type message struct {
	Min int64
	Max int64
}

// Solve Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
func Solve() {
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
		if tools.IsPrime(d) && d > max {
			max = d
			fmt.Printf("Max prime divisor at the moment: %d\n", max)
		}
	}

	fmt.Printf("Result: %d\n", max)
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
