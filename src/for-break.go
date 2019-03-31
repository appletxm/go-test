package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var i, j int

  for i = 0; i < 10; i++ {
    fmt.Printf("===%d\n", i)
    for j = 0; j < 10; j++ {
      fmt.Printf("***%d\n", j)
      if (j % 2 == 0) {
        // break
        continue
      }
    }
  }  
}
