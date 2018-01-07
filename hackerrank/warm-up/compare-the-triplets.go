package main

import "fmt"

func main() {
	var a0, a1, a2, b0, b1, b2 int
	inp := []*int{&a0, &a1, &a2, &b0, &b1, &b2}
	for i := 0; i < len(inp); i++ {
		fmt.Scan(inp[i])
	}
	list := []int{a0 - b0, a1 - b1, a2 - b2}
	aPoint := 0
	bPoint := 0
	for i := 0; i < len(list); i++ {
		if list[i] > 0 {
			aPoint = aPoint + 1
		} else if list[i] < 0 {
			bPoint = bPoint + 1
		}
	}
	fmt.Println(aPoint, bPoint)
}
