package main

import "fmt"
//内建变量类型
//1、bool,string
//2、(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
//3、byte,rune
//4、float32,float64,complex64,complex128
var a,b = 1,2
var (
	c,d ,e= 3,4,-12.2
)
func variable()  {
	var a int
	var s string
	fmt.Println(a,s)
	fmt.Printf("%q=%s=\n",s,s)
}
func variableInitialVal()  {
	var a,b = 3,4
	var s string ="string"
	c,d,e := 4,'a',"sss"
	fmt.Println(a,b,s,c,d,e)
}
func enums()  {
	const (
		cpp = iota
		java
		js
		_
		php

	)
	//b,kb,mb,gb,tb,pb
	const(
		b = 1<<(10*iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,js,php)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	variable()
	variableInitialVal()
	fmt.Println(a,b,c,d,e)
	enums();
}
