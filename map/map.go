package main

import "fmt"
//map 不保证遍历顺序，如需顺序需要手动对key排序
//map 的key,除了slice,map,function其他内建类型都可以作为key,struct类型不包括slice,map和function也能做为key
func main() {
	//建立方法一
	m := map[string]string{
		"name":  "json",
		"couse": "go",
	}
	fmt.Println(m)
	//建立方法二：
	m2 := make(map[string]int) //m2值是empty map
	var m3 map[string]int      //m3的值是nil
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	for k, v := range m3 {
		fmt.Println(k, v)
	}
	m2["name"] = 10
	fmt.Println(m2 == nil)
	fmt.Println(m3 == nil)
	//m3["name"]=10  //map的值是nil时不能直接赋值，这里会报错
	m3 = map[string]int{
		"name": 10, //这种写法就可以
	}

	fmt.Println(m2, m3)
	fmt.Println("GETting values")
	courseName := m["course"]
	fmt.Println(courseName)
	//判断key是否存在,不存在时，获得零值
	if courseName, ok := m["couse"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("unknown")
	}
	fmt.Println("删除元素")
	delete(m, "couse")
	if courseName, ok := m["couse"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("unknown")
	}
	fmt.Println(norepeating("aaabbccesaabc"))
	fmt.Println(norepeating("123412"))
	fmt.Println(norepeating("一二三一"))
}

//最长不重复子串
func norepeating(s string) int  {
	lastoccurred := make(map[rune]int)
	start := 0;
	maxLength := 0;
	for i,ch := range []rune(s){
		if lastI,ok :=  lastoccurred[ch];ok && (lastI >= start){
			start = lastoccurred[ch] +1
		}
		if i-start+1 >maxLength {
			maxLength = i - start +1
		}
		lastoccurred[ch] = i
	}
	return  maxLength
}
