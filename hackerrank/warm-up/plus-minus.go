package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	pos, neg, zero := 0, 0, 0

	for i := 0; i < n; i++ {
		var element int
		fmt.Scan(&element)

		if element > 0 {
			pos = pos + 1
		} else if element == 0 {
			zero = zero + 1
		} else {
			neg = neg + 1
		}
	}

	fmt.Printf("%f\n", float64(pos)/float64(n))
	fmt.Printf("%f\n", float64(neg)/float64(n))
	fmt.Printf("%f\n", float64(zero)/float64(n))
}
