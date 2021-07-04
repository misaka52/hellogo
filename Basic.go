package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"strconv"
)

func variablesZero() {
	var a int
	var b string
	fmt.Println(a, b)
	// %q 携带引号打印字符串
	fmt.Printf("%d %q\n", a, b)
}

func variablesInit() {
	// 变量初始化
	var a int = 1
	var b string = "ddd"
	fmt.Println(a, b)
}

func variablesInit2() {
	// 多变量初始化
	var a, b, c = 1, true, "str"
	fmt.Println(a, b, c)
}

func variablesInit3() {
	// :=多变量初始化
	a, b, c := 1, true, "str"
	fmt.Println(a, b, c)
}

var a = 1

// var b := 1，报错，:=不能用于全局变量
var (
	aa = 3
	bb = true
	cc = "fa"
)

func globalVariables() {
	fmt.Println(a, aa, bb, cc)
}

func euler() {
	a := 3 + 4i
	fmt.Println(cmplx.Abs(a))
	// 欧拉公式e^iπ + 1 = 0
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

	var d float64 = 4.99
	// 直接抹除小数位，靠近0取整
	c = int(d)
	fmt.Println(int(c))
}

func consts() {
	const filename = "a.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enum() {
	// iota自增值，从0开始
	const (
		c = iota
		java
		_
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(c, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb)
}

func ifCondition() {
	const filename = "a.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func switchCondition(score int) string {
	grade := ""
	switch {
	case score < 0 || score > 100:
		// 抛出异常
		panic(fmt.Sprintf("invalid score, %d", score))
	case score < 40:
		grade = "D"
		// 忽略等级D
		fallthrough
	case score < 60:
		grade = "C"
	case score < 80:
		grade = "B"
	case score < 100:
		grade = "A"
	}
	return grade
}

func intToBin(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for ; num > 0; num /= 2 {
		tmp := num % 2
		result = strconv.Itoa(tmp) + result
	}
	return result
}

func main() {
	fmt.Println("hello world")
	variablesZero()
	variablesInit()
	variablesInit2()
	variablesInit3()
	globalVariables()
	euler()
	triangle()
	enum()

	ifCondition()
	fmt.Println(
		//switchCondition(-1),
		switchCondition(59),
		switchCondition(79),
		switchCondition(99),
	)

	fmt.Println(
		intToBin(3),
		intToBin(0),
		intToBin(321231987))
}
