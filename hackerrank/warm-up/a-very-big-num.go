package main

import (
	"fmt"
)

func main() {
	var n int
	var sum int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var element int
		fmt.Scan(&element)
		sum = sum + element
	}

	fmt.Println(sum)
}
