package main

import "fmt"

//切片半开半闭区间
//slice肚子里面有一个结构，有数组的指针
func updateSlice(s []int) {
	s[0] = 1000
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6] //2,3,4,5 半开半闭区间
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println(arr)
	fmt.Println(s1)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(s2)
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
	s3 := append(s2, 10) //添加元素时如果超越cap,就会分配更大的底层数组，
	s4 := append(s3, 11)
	s5 := append(s3, 12)
	fmt.Println(s3, s4, s5, arr)
	var ss []int
	sss := make([]int, 16)
	ssss := make([]int, 2, 32)
	for i := 0; i < 100; i++ {
		fmt.Printf("len=%d,cap=%d\n", len(ss), cap(ss))
		ss = append(ss, 2*i+1)
	}
	fmt.Printf("len=%d,cap=%d\n", len(sss), cap(sss))
	fmt.Printf("len=%d,cap=%d\n", len(ssss), cap(ssss))
	//copying slice
	copy(ssss, s2)
	fmt.Printf("len=%d,cap=%d,value=%v\n", len(ssss), cap(ssss), ssss)

	//delete
	fmt.Println(s2)
	ss2 := append(s2[:2], s2[3:]...) //删除其中的一个元素 注意...的用法
	fmt.Println(ss2)
	fmt.Println("删除头元素")
	s2 = s2[1:]
	fmt.Println(s2)
	fmt.Println("删除尾部的元素")
	s2 = s2[:len(s2)-1]
	fmt.Println(s2)

}
