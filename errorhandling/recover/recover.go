package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	// 创建匿名函数并调用
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't konw what to handle it: %v", r))
		}
	}()
	// 1. 创建异常
	panic(errors.New("this is an error"))
	// 2. 其他类型
	//panic(123)
}

func main() {
	tryRecover()
}
