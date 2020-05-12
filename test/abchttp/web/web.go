package main

import (
	"golangdemo/test/abchttp"
	"net/http"
)

func main() {

	http.HandleFunc("/", abchttp.ErrWrapper(abchttp.FileListingHandleThrows))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
