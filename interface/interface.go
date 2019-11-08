package main

import (
	"fmt"
	real2 "github.com/20100204/xuexi/interface/real"
	"github.com/20100204/xuexi/interface/retriever"
)

type RetrieverPoster interface {
	retriever.Retriever
	retriever.Poster
}

func main() {
	var r RetrieverPoster
	r = real2.Retriever{}
	fmt.Println(r.Post("wwww",map[string]string{}))
	//fmt.Println(r.Get("http://www.imooc.com"))
	fmt.Printf("%T %v\n", r, r)


}
