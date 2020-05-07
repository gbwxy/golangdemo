package main

import (
	"fmt"
	"golangdemo/queue"
)

func main() {

	q := queue.Queue{1, 2, 3}
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}
