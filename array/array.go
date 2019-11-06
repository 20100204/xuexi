package main

import "fmt"

func printArray(arr [5]int)  {

}
func main() {
	var arr1 [5]int //定义方法1
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [2][3]int
	fmt.Println(arr1, arr2, arr3, grid)
	//遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//遍历2
	for k, v := range arr3 {
		fmt.Printf("key is %d,value is %d\n", k, v)
	}
	printArray(arr3)
	//printArray(arr2) //数组长度不一样，类型不一样
}
