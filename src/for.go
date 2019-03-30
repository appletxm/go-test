package main

import "fmt"

func main() {

   var b int = 15
   var a int

   numbers := [4]string{"a", "b", "c", "d"} 

   /* for 循环 */
   for a := 0; a < 10; a++ {
      fmt.Printf("a 的值为: %d\n", a)
   }

   for a < b {
      a++
      fmt.Printf("a 的值为: %d\n", a)
   }

   for i, x := range numbers {
      fmt.Printf("第 %d 位 x 的值 = %s\n", i, x)
   }   
}