package main

import "fmt"

var arr = [6]int{1,5,6,7,8,9}

func main() {
  var numbers = make([]int,3,5)
  slice1 := []int{1,2,3 }
  slice2 := arr[ : ]
  slice3 := arr[0 : ]
  slice4 := arr[0 : 2]
  slice5 := make([]int, 3, 5) 

  fmt.Printf("slice1=%v slice2=%v slice3=%v slice4=%v slice5=v%\n", slice1, slice2, slice3, slice4, slice5)

  printSlice(numbers)
}

func printSlice(x []int){
  fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
