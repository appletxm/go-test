package main

import "fmt"

func main() {
  var a int = 10
  var x int = 10
  var y int = 5  

  fmt.Printf("变量的地址: %x\n", &a  )

  checkAddress(&x, &y)
}

func checkAddress(x *int, y *int) {
  fmt.Printf("x address is %x, y address is %x", &x, &y)
}
