package main
import "fmt"

func main() {
  // var map1 map[string]string
  map1 := make(map[string]string)

  map1["a"] = "aaaaaa"
  map1["b"] = "bbbbbb"
  map1["c"] = "ccccccccc"

  for key,v := range map1 {
    fmt.Printf("%s = %s\n", key, v)
  }

  key2, val2 := map1["c"]
  if (val2) {
    fmt.Println("==1===", key2, val2)
  } else {
    fmt.Println("==2===", key2, val2)
  }
}
