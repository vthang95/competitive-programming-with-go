# Problem 2

(Link Leetcode)[https://leetcode.com/contest/weekly-contest-97/problems/spiral-matrix-iii/]

# Solution

```go
package main

import (
  "fmt"
)

func spiralMatrix(R int, C int, r0 int, c0 int) [][]int {
  dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
  totalStep := R * C
  res := make([][]int, totalStep) 
  res[0] = []int{r0, c0}
  step := 1
  curr := 0
  edge := 0

  for step < totalStep {
    if curr == 0 || curr == 2 {
      edge++
    }

    for i := 0; i < edge; i++ {
      r0 += dirs[curr][0] 
      c0 += dirs[curr][1]

      if isValidCor(R, C, r0, c0) {
        res[step] = []int{r0, c0}
        step++
      }
    }

    curr = (curr + 1) % 4
  }

  return res
}

func isValidCor(R int, C int, r0 int, c0 int) bool {
  return r0 >= 0 && r0 < R && c0 >= 0 && c0 < C
}

func main() {
  var R, C, r0, c0 int
  fmt.Scan(&R)
  fmt.Scan(&C)
  fmt.Scan(&r0)
  fmt.Scan(&c0)

  fmt.Println(spiralMatrix(R, C, r0, c0))
}
```
