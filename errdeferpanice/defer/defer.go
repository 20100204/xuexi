package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

type intGen func() int

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custome error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError.Err)
			fmt.Println(pathError.Path)
			fmt.Println(pathError.Op)
			fmt.Println(pathError.Error())
			return
		} else {
			panic(err)

		}
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}
func main() {
	tryDefer()
	writeFile("fibonacci.txt")
}
