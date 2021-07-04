package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func operator(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupportted operation %s", op)
	}
}

// 函数式编程
func apply(op func(int, int) int, a, b int) int {
	// 反射获取函数名
	p := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("call fun %s with args(%d,%d)\n", name, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum(values ...int) int {
	res := 0
	// num为下标，从0开始
	for num := range values {
		res += values[num]
	}
	return res
}

// 值传递
func swap(a, b int) {
	a, b = b, a
}

// 引用传递
func swap2(a, b *int) {
	*a, *b = *b, *a
}

func swap3(a, b int) (int, int) {
	return a, b
}

func main() {
	// _ 忽略结果
	q, _ := div(13, 4)
	fmt.Println(q)

	fmt.Println(apply(pow, 2, 5))

	// 匿名函数
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 2, 5))

	fmt.Println(sum(1, 2, 3, 4))

	a, b := 3, 4
	swap2(&a, &b)
	fmt.Println(a, b)
}
