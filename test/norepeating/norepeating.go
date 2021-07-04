package main

import "fmt"

var lastOccurred = make([]int, 0xffff)

// 最长不重复子串
func MaxLenNoRepeatString(resource string) int {
	//lastOccurred := make(map[rune]int)

	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLen := 0
	for i, ch := range []rune(resource) {
		if lastIndex := lastOccurred[ch]; lastIndex != -1 && lastIndex >= start {
			start = lastIndex + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}

func main() {
	fmt.Println(
		MaxLenNoRepeatString("aaa"),
		MaxLenNoRepeatString("abcde"),
		MaxLenNoRepeatString(""),
		MaxLenNoRepeatString("a"),
		MaxLenNoRepeatString("abbcdeop"),
		MaxLenNoRepeatString("中文字符串"),
		MaxLenNoRepeatString("一二三二一"),
	)
}
