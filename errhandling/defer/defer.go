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

	fmt.Println(3)

	panic("error occurred ...")
	fmt.Println(4)
}

func fibonacci() func() uint64 {
	a, b := uint64(0), uint64(1)
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)

	//自己创建error
	err = errors.New("This is a custom error.")

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			//fmt.Println("error : ", err)
			//非已知的可以直接抛出
			panic(err)
		} else {
			//处理已知的error
			fmt.Printf("%s, %s, %s \n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
