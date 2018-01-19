package main

import (
	"fmt"
)

func compress(input string) (output string) {
	inx := make(map[string]uint)
	for _, char := range input {
		inx[string(char)]++
	}
	for key, val := range inx {
		output += fmt.Sprintf("%d%s", val, key)
	}
	return output
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(compress(input))
}