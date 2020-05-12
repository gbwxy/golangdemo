package basic

import (
	"bytes"
	"fmt"
	"math"
	"math/cmplx"
)

//包变量- 包里面是可用的，并不是全局变量
var (
	a_pkg = 123
	b_pkg = true
	c_pkg = "kkk"
)

//包内不能用 := 声明变量
//abc, bcd := 123, false

func variableZeroValue() {
	//函数局部变量
	var a int
	var s string
	fmt.Printf("%d    %q", a, s)
	fmt.Println()
}

func variableInitialValue() {
	var a, b int = 1, 2
	var s string = "abd"
	fmt.Println(a, b, s)
}

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 30, 40
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

	////var a, b  = 30, 40
	//var a, b  float64 = 30, 40
	////var c int
	//c = int(math.Sqrt(a*a + b*b))
	////fmt.Println(c)
}

//func Triangle(a int, b int) int {
//	c := int(math.Sqrt(float64(a*a + b*b)))
//	return c;
//}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, false, "def"
	fmt.Println(a, b, c, s)
}

func variableShortor() {
	abc, bcd := 123, false
	fmt.Println(abc, bcd)
	ss := "abs"
	fmt.Println(ss)
	ss = "abcd"
	fmt.Println(ss)

	var (
		s_a = 2
		s_d = 3
	)
	fmt.Println(s_a, s_d)
}

const fileName = "abc.txt"
const (
	file1 = "file1"
	file2 = "file2"
)

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
	//fmt.Println(fileName, file1, file2)
}

func hello(abc ...interface{}) {

	//fmt.Println(abc)

	for argNum, arg := range abc {
		if argNum > 0 {
			fmt.Println(argNum, arg)
		}
	}

}

func enums() {
	const (
		cpp = iota
		java
		python
		c
		golang
		javascript = 100
		_
		csharp = iota
	)
	fmt.Println(cpp, java, python, c, golang, javascript, csharp)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	//fmt.Println("Hello world!")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShortor()
	//
	//fmt.Println(a_pkg, b_pkg, c_pkg)

	//euler()
	//triangle()

	//var str string
	//str = "anc"
	//fmt.Println(str)
	//consts()
	//hello(fileName, file1, file2, 123)

	enums()

	var bufabc bytes.Buffer
	fmt.Println(bufabc)

}
