package main

import "fmt"

//指针不能运算
var a int = 2
var pa *int = &a //把a的地址给 pa指针变量

func main() {
	*pa = 4 //a的值也改变了
	fmt.Println(a)
}
