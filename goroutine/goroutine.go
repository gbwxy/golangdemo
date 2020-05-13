package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	var a [10]int
	//for i := 0; i < 10; i++ {
	//	go func(ii int) {
	//		for {
	//			//fmt.Println(a[ii])
	//			a[ii]++
	//			runtime.Gosched()
	//		}
	//	}(i)
	//}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				//fmt.Println(a[ii])
				a[i]++ //这里的i 为闭包
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(a)
}
