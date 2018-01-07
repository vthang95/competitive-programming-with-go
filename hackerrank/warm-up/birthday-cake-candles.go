package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	max := 0
	result := 0
	for i := 0; i < number; i++ {
		var element int
		fmt.Scan(&element)
		if max < element {
			max = element
			result = 1
		} else if max == element {
			result += 1
		}
	}

	fmt.Println(result)
}
