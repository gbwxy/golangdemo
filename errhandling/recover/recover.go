package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

func tryRecover() {

	//defer 必须是一个操作，所以后面跟一个匿名函数，然后需要执行这个匿名函数
	//func() {}()
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			str := fmt.Sprintf("Error Occurred, error : %s", err)
			fmt.Println(str)
			log.Error(str)
		} else if err != nil {
			panic(err)
		}
	}()

	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {

	tryRecover()
}
