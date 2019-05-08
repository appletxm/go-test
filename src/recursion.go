package main
import "fmt"

func Factorial(n int)(result int) {
  if (n > 0) {
    result = n * Factorial(n-1)
    return result
  }

  return 1
}

func main() {
  var i int = 5
  fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))
}
