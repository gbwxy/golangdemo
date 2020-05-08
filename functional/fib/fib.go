package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() func() uint64 {
	a, b := uint64(0), uint64(1)
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func fibonacci2() intGen {
	a, b := uint64(0), uint64(1)
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type intGen func() uint64

func (g *intGen) Read(p []byte) (n int, err error) {
	next := (*g)()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func main() {
	//f := fibonacci()
	////for i := 0; i < 1000000; i++ {
	////	fmt.Println(f())
	////}
	//
	//for i := 0; i < 15; i++ {
	//	fmt.Println(f())
	//}

	f := fibonacci2()
	printFileContents(&f)

}
