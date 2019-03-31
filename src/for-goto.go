package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var i, j = 0, 4

  LOOP1: for i = 0; i < 10; i++ {
    j++
    fmt.Printf("===%d\n", i)
  }

  // LOOP2: fmt.Printf("######\n")

  if (j < 2) {
    goto LOOP1
  } else {
    fmt.Printf("******\n")
    // goto LOOP2
  }
}
