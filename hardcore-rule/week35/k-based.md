# Problem 1

(Link Timus)[http://acm.timus.ru/problem.aspx?space=1&num=1009<Paste>]

# Solution

```go
package main

import (
  "fmt"
)

func main() {
  var n, k int
  fmt.Scan(&n)
  fmt.Scan(&k)

  fmt.Println(find_quantity(n, k))
}

func find_quantity(n int, k int) int {
  if n == 0 { return 0 }
  if n == 1 { return k }

  return (k - 1) * count(n - 1, k, false)
} 

func count(i int, k int, isZero bool) int {
  if i == 0  { return 0 }
  if isZero {
    if i == 1 { return k - 1 }
    return (k - 1) * count(i - 1, k, false)
  } else {
    if i == 1 { return k } 
    return count(i - 1, k, true) + (k - 1) * count(i - 1, k , false)
  }
}
```
