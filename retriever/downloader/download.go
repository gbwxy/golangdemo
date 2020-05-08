package main

import (
	"fmt"
	"golangdemo/retriever/mock"
	real2 "golangdemo/retriever/real"
	"strconv"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {

	var r Retriever

	//fmt.Println("mock ...")
	//r = mock.Retriever{"this is a fake imooc.com"}
	//fmt.Printf("%T  %v \n", r, r)
	//s := download(r)
	//fmt.Println(s)
	//
	//fmt.Println("real ...")
	//r = real2.Retriever{
	//	UserAgent: "Mozilla/5.0",
	//	TimeOut:   time.Minute}
	//fmt.Printf("%T  %v \n", r, r)
	//s = download(r)
	////fmt.Println(s)
	//
	//fmt.Println("real pointer ...")
	//r = &real2.RetrieverPointer{
	//	UserAgent: "Mozilla/5.0",
	//	TimeOut:   time.Minute}
	//fmt.Printf("%T  %v \n", r, r)
	//s = download(r)
	////fmt.Println(s)

	fmt.Println("mock ...")
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)

	fmt.Println("real ...")
	r = real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute}
	inspect(r)

	fmt.Println("real pointer ...")
	r = &real2.RetrieverPointer{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute}
	inspect(r)

	realRep := r.(*real2.RetrieverPointer)
	fmt.Println(realRep.TimeOut)

	//realRe := r.(real2.Retriever)
	//fmt.Println(realRe.TimeOut)

	////查看接口变量-Type Assertion
	//if mockRe, ok := r.(mock.Retriever); ok {
	//	fmt.Println(mockRe.Context)
	//} else {
	//	fmt.Println("not a mock ...")
	//}

	//只有interface类型的变量可用 .(type)
	var itf interface{}
	itf = "123"
	//fmt.Println(str.(type))//type switch 可以用此写法H
	i := itf.(int)
	fmt.Println(i)

	str := "123"
	//ii = int(str)
	if ii, err := strconv.Atoi(str); err != nil {
		panic(err)
	} else {
		fmt.Println(ii)
	}

	idx := 123
	flo1 := float64(idx)
	var flo float64
	flo = float64(idx)
	print(flo, flo1)

	////查看接口变量
	//mockRe := mock.Retriever(r)
	//fmt.Println(mockRe.Context)

}

//查看接口变量-Type switch
func inspect(r Retriever) {
	fmt.Printf("%T  %v \n", r, r)
	//type switch 可以用此写法
	switch r := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", r.Context)
	case real2.Retriever:
		fmt.Println("UserAgent:", r.UserAgent)
	case *real2.RetrieverPointer:
		fmt.Println("UserAgent:", r.UserAgent)
	}

}
