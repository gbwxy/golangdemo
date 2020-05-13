package main

import (
	"fmt"
	"time"
)

func worker(i int, ch chan int) {
	//for {
	//	//fmt.Printf("worker %d received %c \n", i, <-ch)
	//	fmt.Printf("worker %d received %d \n", i, <-ch)
	//}

	//判断是否发送完成
	//for {
	//	//fmt.Printf("worker %d received %c \n", i, <-ch)
	//	num, ok := <-ch;
	//	if (!ok) {
	//		break
	//	}
	//	fmt.Printf("worker %d received %d \n", i, num)
	//}

	//range 判断是否发送完成
	for num := range ch {
		fmt.Printf("worker %d received %d \n", i, num)
	}
}

// send chan : chan<- int
// receive chan : <-chan int
func createWork(i int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c \n", i, <-c)
		}
	}()
	return c
}

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func chanDemo2() {
	//go func() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for j := 0; j < 20; j++ {
		for i := 0; i < 10; i++ {
			channels[i] <- 'a' + i
		}
	}
	//}()
}

func channelDemo3() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}
	for j := 0; j < 20; j++ {
		for i := 0; i < 10; i++ {
			channels[i] <- 'a' + i
		}
	}
	time.Sleep(10 * time.Millisecond)
}
func bufferChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(10 * time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	//close-终止发送
	close(c)
	time.Sleep(10 * time.Millisecond)
}

func main() {
	//chanDemo()
	//chanDemo2()
	//channelDemo3()
	//bufferChannel()
	channelClose()
}
