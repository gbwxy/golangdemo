package basic

import "math"

func Triangle(a int, b int) int {
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func GetNonRepeatLen(s string) int {
	start := 0
	maxLength := 0
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
