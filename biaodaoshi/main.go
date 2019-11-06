package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ifbiaodashi() {
	filename := "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//fmt.Println(contents) 出了if语句就不能用
}
func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic("unknown")
	}
	return g
}
func switchbiaodashi(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("未知的操纵符")

	}
}

func convertToBin(n int) string {
	rs := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		rs = strconv.Itoa(lsb) + rs
	}
	return rs
}

/*
一行一行的读文件
*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
/* 死循环*/
func forever()  {
	for{
		fmt.Println("233333")
	}
}
//form
func forbiaodashi(n int) int {
	g := 0
	for i := 0; i < n; i++ {
		g += i
	}
	return g
}
func main() {
	//forever();
	//printFile("abc.txt")
	//fmt.Println(convertToBin(13))
	//fmt.Println(forbiaodashi(10))
	//fmt.Println(grade(10))
	//	fmt.Println(grade(80))
	//fmt.Println(grade(90))
	//fmt.Println(grade(190))

	//fmt.Println(switchbiaodashi(1,2,"+"))
	//fmt.Println(switchbiaodashi(1,2,"-"))
	//fmt.Println(switchbiaodashi(1,2,"*"))
	//fmt.Println(switchbiaodashi(1,2,"/"))
	//fmt.Println(switchbiaodashi(1,2,"a"))

}
