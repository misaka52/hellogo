package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func get(m map[string]string, key string) {
	if value, ok := m[key]; ok {
		fmt.Printf("get(%s)=%s\n", key, value)
	} else {
		fmt.Println("not existed key", key)
	}
}

func mapOp1() {
	m := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	}
	fmt.Println(m)

	// 遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
	get(m, "k1")
	get(m, "kk")

	delete(m, "k2")
	fmt.Println(m)
}

func charOp() {
	s := "go语言学习！"
	fmt.Println(len(s))
	for k, v := range s {
		fmt.Printf("(%d %X) ", k, v)
	}
	fmt.Println()
	for k, v := range []byte(s) {
		fmt.Printf("(%d %X) ", k, v)
	}
	fmt.Println()

	for k, v := range []rune(s) {
		fmt.Printf("(%d %c) ", k, v)
	}
	fmt.Println()

	fmt.Printf("(%s)字符长度=%d, 字节长度=%d\n", s, utf8.RuneCountInString(s), len(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
}

func stringOp() {
	s := " I'm fine, golang learning !!!  "
	// 按照空白字符串切分字符串
	ss := strings.Fields(s)
	for k, v := range ss {
		fmt.Println(k, v)
	}
	// 指定字符分割
	ss2 := strings.Split(s, ",")
	for k, v := range ss2 {
		fmt.Println(k, v)
	}
	// 字符串数组拼接，采用指定字符串拼接
	t := strings.Join(ss, "abc")
	fmt.Println(t)
	// 找到第一个目标字符的索引位置
	fmt.Println(strings.Index(s, "ll"))
	fmt.Println(strings.ToLower(s))
	// 去除首尾指定字符串
	fmt.Println(strings.Trim(s, " "))
}

func main() {
	//mapOp1()
	//charOp()
	//maxLenNoRepeatStringTest()
	stringOp()
}
