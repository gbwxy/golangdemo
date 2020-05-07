package basic

import "fmt"

////switch 会自动break，除非使用fallthrough
//func eval(a, b int, op string) int {
//	var result int
//	switch op {
//	case "+":
//		result = a + b
//	case "-":
//		result = a - b
//	case "*":
//		result = a * b
//	case "/":
//		result = a / b
//	default:
//		panic("unsupported operator:" + op)
//	}
//	return result
//}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	//const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	//fmt.Println()

	//a, b := 12, 4
	//fmt.Println(eval(a, b, "*"))
	//fmt.Println(eval(a, b, ".."))

	fmt.Println(grade(99))
	fmt.Println(grade(88))
	fmt.Println(grade(77))
	fmt.Println(grade(55))
	fmt.Println(grade(199))
}
