package main

import (
	"fmt"
	"strings"
)

func main() {
	var level int
	fmt.Scan(&level)

	for i := 0; i < level; i++ {
		fmt.Println(strings.Repeat(" ", level-i-1) + strings.Repeat("#", i+1))
	}
}
