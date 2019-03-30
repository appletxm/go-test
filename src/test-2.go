package main

var a = "菜鸟教程"
var b string = "runoob.com"
var c bool

var var1, var2, var3 = 1, "aaa", "true"
var (  // 这种因式分解关键字的写法一般用于声明全局变量
  global1 int
  global2 bool
)

func main(){
  g, h := 123, "hello"
  println(a, b, c)
  println(var1, var2, var3, g, h, global1, global2)
}
