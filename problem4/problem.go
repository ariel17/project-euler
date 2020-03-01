package problem4

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

type pair struct {
	x int
	y int
}

// Solve Largest palindrome product
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
func Solve() {

	pairs := []pair{}

	for i := 100; i <= 999; i++ {
		for j := 1; j <= 999; j++ {
			if i > j {
				continue
			}
			pairs = append(pairs, pair{x: i, y: j})
		}
	}

	maxPalindrome := 0
	var mp pair
	for _, p := range pairs {
		v := p.x * p.y
		sv := fmt.Sprintf("%d", v)
		r := stringutil.Reverse(sv)
		if sv == r && v > maxPalindrome {
			maxPalindrome = v
			mp = p
		}
	}

	fmt.Printf("> Result: %d x %d = %d", mp.x, mp.y, maxPalindrome)
}
