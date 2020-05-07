package container

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len(s) = %d, cap(s) = %d \n", len(s), cap(s))
}
func main() {

	var s []int
	for i := 0; i < 20; i++ {
		//fmt.Println(s)
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d ", s, len(s), cap(s))

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	fmt.Println(s2)
	printSlice(s2)

	s3 := make([]int, 16, 32)
	fmt.Println(s2)
	printSlice(s3)

	fmt.Println("Copy slice.")
	copy(s2, s1)
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))

	fmt.Println("Deleting element.")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))

	fmt.Println("Deleting top.")
	s2 = s2[1:]
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))

	fmt.Println("Deleting tail.")
	s2 = s2[:len(s2)-1]
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d \n", s2, len(s2), cap(s2))

}
