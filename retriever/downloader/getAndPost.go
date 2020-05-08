package main

import (
	"fmt"
	"golangdemo/retriever/mock"
)

type getItf interface {
	Get(url string) string
}

type postItf interface {
	Post(url string, from map[string]string) string
}
type getAndPostI interface {
	getItf
	postItf
}

func getAndPostMoth(r getAndPostI) string {
	r.Post("http://www.imooc.com", map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
	return r.Get("http://www.imooc.com")
}

func main() {
	var r getAndPostI
	fmt.Println("mock ...")
	r = &mock.Retriever2{}

	i := getAndPostMoth(r)
	fmt.Println(i)

}
