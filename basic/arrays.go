package basic

import "fmt"

func printarray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printarrayponiter(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int

	fmt.Println("init arrs.")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grid)

	//数组的遍历
	fmt.Println("get arr3 by for.")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	//获取下标
	fmt.Println("get arr3 by range i.")
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	//获取下标和值
	fmt.Println("get arr3 by range i & v.")
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	//获取值
	fmt.Println("get arr3 by range v.")
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println("get printarray(arr3).")
	printarray(arr3)
	fmt.Println("get printarray(arr1).")
	printarray(arr1)
	fmt.Println("after execute.")
	fmt.Println(arr1, arr3)

	fmt.Println("get printarrayponiter(&arr1).")
	printarrayponiter(&arr1)
	fmt.Println("get printarrayponiter(&arr3).")
	printarrayponiter(&arr3)
	fmt.Println("after execute.")
	fmt.Println(arr1, arr3)

}
