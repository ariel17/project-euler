package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Printf("Result: %d", total)
}
