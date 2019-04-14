package main

import "fmt"

var arr = [10]float32{1.2, 56.09}
var a = [3][4]int{  
  {0, 1, 2, 3} ,   /*  第一行索引为 0 */
  {4, 5, 6, 7} ,   /*  第二行索引为 1 */
  {8, 9, 10, 11},   /* 第三行索引为 2 */
 }

func main()  {
  fmt.Println(arr[0])
  arr[2] = 123.60
  fmt.Println(arr)
}
