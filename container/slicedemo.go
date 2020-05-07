package container

import "fmt"

func updateslice(arr []int) {
	arr[0] = 100
}

func main() {
	//var a, b int
	//a = 100
	//b = 222
	//fmt.Println(a, b)

	//arr:=[...]int{0,1,2,3,4,5,6,7,8,9}
	//fmt.Println("arr[2:6] = ",arr[2:6])
	//fmt.Println("arr[:6] = ",arr[:6])
	//fmt.Println("arr[2:] = ",arr[2:])
	//fmt.Println("arr[:] = ",arr[:])

	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := arr[2:]
	//s2 := arr[:]
	//fmt.Println("arr[2:] = ", s1)
	//fmt.Println("arr[:] = ", s2)
	//
	//updateslice(s1)
	//fmt.Println("after update s1.")
	//fmt.Println("arr[2:] = ", s1)
	//fmt.Println("arr[:] = ", s2)
	//
	//updateslice(s2)
	//fmt.Println("after update s2.")
	//fmt.Println("arr[2:] = ", s1)
	//fmt.Println("arr[:] = ", s2)

	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := arr[2:6]
	//s2 := s1[3:5]
	//
	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(s1[5:9])

	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := arr[2:6]
	//s2 := s1[3:5]
	//
	//fmt.Println("arr = ", arr)
	//fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d \n", s1, len(s1), cap(s1))
	//fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))

	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s1 := arr[2:6]
	//s2 := s1[3:5]
	//
	//fmt.Printf("arr = %v, len(arr) = %d, cap(arr) = %d \n", arr, len(arr), cap(arr))
	//fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))
	//
	//
	//s3 := append(s2, 10)
	//s4 := append(s3, 11)
	//s5 := append(s4, 12)
	//s6 := append(s5, 13)
	//s7 := append(s6, 14)
	//
	//fmt.Println("after append. ")
	//fmt.Printf("arr = %v, len(arr) = %d, cap(arr) = %d \n", arr, len(arr), cap(arr))
	//fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))
	//fmt.Printf("s3 = %v, len(s3) = %d, cap(s3) = %d \n", s3, len(s3), cap(s3))
	//fmt.Printf("s4 = %v, len(s4) = %d, cap(s4) = %d \n", s4, len(s4), cap(s4))
	//fmt.Printf("s5 = %v, len(s5) = %d, cap(s5) = %d \n", s5, len(s5), cap(s5))
	//fmt.Printf("s6 = %v, len(s6) = %d, cap(s6) = %d \n", s6, len(s6), cap(s6))
	//fmt.Printf("s7 = %v, len(s7) = %d, cap(s7) = %d \n", s7, len(s7), cap(s7))arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s1 := arr[2:6]
	//s2 := s1[3:5]
	//
	//fmt.Printf("arr = %v, len(arr) = %d, cap(arr) = %d \n", arr, len(arr), cap(arr))
	//fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))
	//
	//
	//s3 := append(s2, 10)
	//s4 := append(s3, 11)
	//s5 := append(s4, 12)
	//s6 := append(s5, 13)
	//s7 := append(s6, 14)
	//
	fmt.Println("after append. ")
	fmt.Printf("arr = %v, len(arr) = %d, cap(arr) = %d \n", arr, len(arr), cap(arr))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))
	fmt.Printf("s3 = %v, len(s3) = %d, cap(s3) = %d \n", s3, len(s3), cap(s3))
	fmt.Printf("s4 = %v, len(s4) = %d, cap(s4) = %d \n", s4, len(s4), cap(s4))
	fmt.Printf("s5 = %v, len(s5) = %d, cap(s5) = %d \n", s5, len(s5), cap(s5))
	fmt.Printf("s6 = %v, len(s6) = %d, cap(s6) = %d \n", s6, len(s6), cap(s6))
	fmt.Printf("s7 = %v, len(s7) = %d, cap(s7) = %d \n", s7, len(s7), cap(s7))
}
