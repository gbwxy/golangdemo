package container

import (
	"fmt"
	"strings"
)

/*
	寻找最长不含有重复字符的字串
	1.lastOccurred[x]不存在，或者  <start  -->无需操作
	2.lastOccurred[x] >= start  -->更新start
	3.更新lastOccurred[x]，更新maxLangth
*/
func lengthNonRepeat(s string) int {

	start := 0
	maxLength := 0
	//英文版
	//lastOccurred := make(map[byte]int)
	//for i, ch := range []byte(s) {
	//	if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
	//		start = lastI + 1
	//	}
	//	if i-start+1 > maxLength {
	//		maxLength = i - start + 1
	//	}
	//	lastOccurred[ch] = i
	//}

	//国际版
	lastOccurred := make(map[rune]int)
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	//var s string = "abced";
	//fmt.Println(lengthNonRepeat("abcabcbb"))
	//fmt.Println(lengthNonRepeat("bbbbb"))
	//fmt.Println(lengthNonRepeat("pwwkew"))
	//fmt.Println(lengthNonRepeat("abcedfgh"))
	//fmt.Println(lengthNonRepeat("黑化肥挥发发灰会花飞灰化肥挥发发黑回飞花"))

	var strs []string
	strs = append(strs, "abc")
	strs = append(strs, "123")
	strs = append(strs, "abc")
	strs = append(strs, "123")
	str := strings.Join(strs, ",")

	fmt.Println(str)

	split := strings.Split(str, ",")
	fmt.Println(split)

	str = strings.Join(strs, " ")
	fields := strings.Fields(str)
	fmt.Println(fields)

	if contains := strings.Contains(str, "abc"); contains {
		fmt.Println("contains abc ...")
	}

}
