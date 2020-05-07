package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum() int {

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printfile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	//fmt.Println(sum())
	//fmt.Println(
	//	convertToBin(3),
	//	convertToBin(13),
	//	convertToBin(11122111))

	printfile("abc.txt")
	forever()
}
