package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//switch 会自动break，除非使用fallthrough
func eval(a, b int, op string) (result int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//result = a / b
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//函数声明这里就相当于声明了 q、r
func div(a, b int) (q, r int) {
	q, r = a/b, a%b
	return q, r
}

//
//func apply(op func(int, int) int, a, b int) int {
//	pointer := reflect.ValueOf(op).Pointer()
//	name := runtime.FuncForPC(pointer).Name()
//	fmt.Println("Calling functional %s with args "+
//		"(%d,%d)", name, a, b)
//	return op(a, b)
//}

func apply(op func(int, int) int, a, b int) int {
	fmt.Println("Calling functional %s with args "+
		"(%d,%d)", runtime.FuncForPC(reflect.ValueOf(op).Pointer()), a, b)
	return op(a, b)
}

func power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sumcan(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func main() {
	//fmt.Println(evaltest(3, 4, "*"))
	//q, r := div(13, 3)
	//fmt.Println(q, r)

	//if result, err := eval(3, 4, "*"); err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("result:", result)
	//}
	//
	//if result, err := eval(3, 4, "@*"); err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("result:", result)
	//}

	//a, b := 3, 4
	//i := apply(power, a, b)
	//i2 := apply(func(a int, b int) int {
	//	return int(math.Pow(float64(a), float64(b)))
	//}, a, b)
	//
	//fmt.Println(i, i2)
	fmt.Println(sumcan(1, 2, 3, 4, 5))
	fmt.Println(sumcan(1, 2, 3, 3, 444, 4, 5))
}
