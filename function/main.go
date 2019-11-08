package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"reflect"
	"runtime"
	"strings"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

//匿名返回
func div(a, b int) (int, int) {
	return a / b, a % b
}

//非匿名返回
func bdiv(a, b int) (q, r int) {
	//return a/b,a%b
	q = a / b
	r = a % b
	return
}

//函数做参数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	pName := runtime.FuncForPC(p).Name() //取得函数的名字
	fmt.Printf("Calling funciton %s with args "+"(%d,%d,) \n", pName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//不定长参数
func sum(nums ...int) int {
	s := 0
	for k, i := range nums {
		s += i
		fmt.Printf("key is %d,value is %d\n", k, i)
	}
	return s
}
func adder()func(int)int  {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
func fibonacci() intGen  {
	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}

}
type intGen func() int

func (i intGen) Read(p []byte) (n int, err error) {
	 next := i()
	 if next>3 {
	 	return 0,io.EOF //io.eof  scanner扫描到后会自动结束
	 }
	 s := fmt.Sprintf("%d\n",next) //转成字符串并换一行
	 return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printFileContents(f)
	//fmt.Println(eval(1, 2, "*"))
	//fmt.Println(div(3, 2))
	//q, r := bdiv(3, 2) //写完后面的函数后，按住ctrl+alt+v 会自动生成函数的返回命名参数
	//a, _ := bdiv(3, 2) //只接收一个返回值
	//	fmt.Println(q, r, a)
	//fmt.Println(apply(pow,3,4))
	//匿名函数
	//    fmt.Println(apply(func(i int, i2 int) int {
	//		return int(math.Pow(float64(i),float64(i2)))
	//	},3,4))
	//fmt.Println(sum(1, 2, 3, 4))
	//a := adder()
	//for i :=0;i<10;i++{
	//	fmt.Println(a(i))
	//}
}
