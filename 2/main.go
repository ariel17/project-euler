package main

import "fmt"

const (
	limit = 4000000
)

func main() {
	fib := generate()
	var evenSum int64
	for _, v := range fib {
		if v%2 == 0 {
			evenSum += v
		}
	}
	fmt.Printf("Result: %d", evenSum)
}

func generate() []int64 {
	fib := []int64{1, 2}
	for {
		l := len(fib)
		a := fib[l-1]
		b := fib[l-2]
		r := a + b
		if r > limit {
			return fib
		}
		fib = append(fib, r)
	}
	return fib
}
