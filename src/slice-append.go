package main

import "fmt"

func main() {
  // var slice1 = make([]int, 5, 10)
  var numbers = make([]int, 3, 5)

  fmt.Printf("%v\n", numbers)
  numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7)
  fmt.Printf("%v\n", numbers)
  numbers2 := make([]int, len(numbers), cap(numbers)*2)
  fmt.Printf("%v\n", numbers2)
  copy(numbers2, numbers)
  fmt.Printf("%v\n", numbers2)
}
