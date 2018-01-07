package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	sum = 0
	for i := 0; i < n; i++ {
		var element int
		fmt.Scan(&element)
		sum = sum + element
	}

	fmt.Println(sum)
}
