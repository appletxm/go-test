package main
import "fmt"

func main() {
  countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
  fmt.Println("地图是：")
  for country := range countryCapitalMap {
    fmt.Println(country, "首都是：", countryCapitalMap[country])
  }

  /*删除元素*/ delete(countryCapitalMap, "France")
  fmt.Println("法国条目被删除")

  fmt.Println("删除元素后地图")

  countryCapitalMap["Chnia"] = "Beijing"

  /*打印地图*/
  for country := range countryCapitalMap {
    fmt.Println(country, "首都是", countryCapitalMap [ country ])
  }
}
