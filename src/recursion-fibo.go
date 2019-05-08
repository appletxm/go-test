package main
import "fmt"

var cal int = 0

func fibonacci(n int, nextN int) int {
  // fmt.Printf("current is %d\n", n)
  if (n < 2) {
    if (cal > 0) {
      return cal + n
    }
    return n
  }
  cal = n - 1 + n - 2
  return fibonacci(n - 1, n - 2)
}

func main() {
  var i int
  
  for i = 0; i < 10; i++ {
    cal = 0
    fmt.Printf("%d\t", fibonacci(i, i))
  }
}
