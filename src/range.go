package main
import "fmt"
func main(){
  var nums = []int{1,3,5,7,9}
  var sum = 0
  kvs := map[string]string{"a": "apple", "b": "banana"}

  for _, num := range nums {
    sum += num
  }
  fmt.Println("sum:", sum)

  for i, num := range nums {
    if num == 3 {
      fmt.Println("index:", i)
    }
  }

  for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
  }
}
