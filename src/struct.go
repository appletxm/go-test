package main

import "fmt"

type Books struct {
  title string
  author string
  price float32
}

func main()  {
  javascript := Books {title: "Javacript", author: "Xiaoming", price: 78.50}
  // java := Books {"Java", "XiaoHong", 109.50}
  // cb := Books {"C", "XiaoHong2", 0}

  fmt.Println(javascript)
}
