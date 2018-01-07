package main

import "fmt"

func main() {
	var list []int64
	var sum, min, max int64
	sum = 0
	max = 0
	for i := 0; i < 5; i++ {
		var element int64
		fmt.Scan(&element)
		list = append(list, element)
		sum = sum + element
	}
	min = sum
	for i := 0; i < 5; i++ {
		checkSum := sum - list[i]
		if checkSum < min {
			min = checkSum
		}
		if checkSum > max {
			max = checkSum
		}
	}

	fmt.Println(min, max)
}
