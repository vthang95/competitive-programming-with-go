package main

import (
	"fmt"
	"math"
)

func main() {
	var matrixSize int
	fmt.Scan(&matrixSize)
	sumx, sumy := 0, 0

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			var element int
			fmt.Scan(&element)
			if i == j {
				sumx = sumx + element
			}
			if i+j == matrixSize-1 {
				sumy = sumy + element
			}
		}
	}

	fmt.Println(math.Abs(float64(sumx - sumy)))
}
