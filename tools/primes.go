package tools

import (
	"fmt"
	"strconv"
)

// IsPrime performs checks to determine if n is a prime number.
func IsPrime(n int64) bool {
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
	for i := int64(2); i <= n/2; i++ {
		if n%i == 0 {
			isDivisible = true
			break
		}
	}
	return !isDivisible
}
