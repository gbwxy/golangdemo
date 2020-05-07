package basic

import "fmt"

func swap(a, b int) {
	a, b = b, a
}
func swappointer(a, b *int) {
	*a, *b = *b, *a
}
func swapreturn(a, b int) (int, int) {
	return b, a
}
func main() {

	//var a int = 2
	//fmt.Println(a)
	//var pa *int = &a
	//fmt.Printf("*** before *** pa = %s, *pa = %d, &a = %s, a = %d.\n", pa, *pa, &a, a)
	//*pa = 3
	//fmt.Printf("*** after ***  pa = %s, *pa = %d, &a = %s, a = %d. \n", pa, *pa, &a, a)

	a, b := 3, 4
	fmt.Println(a, b)
	swap(a, b)
	fmt.Println(a, b)
	swappointer(&a, &b)
	fmt.Println(a, b)
	a, b = swapreturn(a, b)
	fmt.Println(a, b)
}
