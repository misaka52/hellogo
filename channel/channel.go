package main

import (
	"fmt"
	"time"
)

func worker(c chan int, id int) {
	//for {
	//	n, ok := <-c
	// 	通过ok判定channel是否关闭
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("%d receiver %d\n", id, n)
	//}
	// range自动判定channel是否关闭
	for n := range c {
		fmt.Printf("%d receiver %d\n", id, n)
	}
}

func createWorker(id int) chan int {
	// 定义channel，返回类型int
	c := make(chan int)
	// 调用监听方法
	go worker(c, id)
	return c
}

func createWorker2(id int) chan<- int {
	// 定义channel，返回类型int
	c := make(chan int)
	// 调用监听方法
	go worker(c, id)
	return c
}

func chanDemo() {
	const n = 10
	var channels [n]chan int
	for i := 0; i < n; i++ {
		channels[i] = createWorker(i + 1)
	}

	for i := 0; i < n; i++ {
		// 向channel发送数据
		channels[i] <- 'a' + i
	}

	for i := 0; i < n; i++ {
		channels[i] <- 'A' + i
	}
	// 睡眠一定时间，保证go runtine方法打印出接收数据
	time.Sleep(time.Millisecond)
}

func chanDemo2() {
	var c chan<- int
	c = createWorker2(0)
	c <- 1
	c <- 2
	c <- 3

	fmt.Println()
	time.Sleep(time.Millisecond)
}

func bufferChan() {
	// channel缓存区大小3，发送消息超过3仍未被接收则发生deadlock
	c := make(chan int, 3)
	go worker(c, 10)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	time.Sleep(time.Millisecond)
}

func chanClose() {
	c := make(chan int)
	go worker(c, 11)
	c <- 1
	c <- 2
	c <- 3

	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//bufferChan()
	//chanClose()
	chanDemo2()
}
