package main

import "fmt"

var g int
func main() {
   /* 声明局部变量 */
   var a, b, c int 

   /* 初始化参数 */
   a = 10
   b = 20
   c = a + b
   g = c

   fmt.Printf ("结果： a = %d, b = %d and c = %d and g = %d\n", a, b, c, g)
}
