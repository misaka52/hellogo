package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(c chan int, id int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("%d receiver %d\n", id, n)
	}
}

// 返回仅发送数据类型的channel
func createWorker(id int) chan<- int {
	// 定义channel，返回类型int
	c := make(chan int)
	// 调用监听方法
	go worker(c, id)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	n := 0
	//over := time.After(8 * time.Second)
	//tick := time.Tick(time.Second)

	var values []int

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
			//case <-time.After(500 * time.Millisecond):
			//	fmt.Println("timeout")
			//case <-tick:
			//	fmt.Println("value len = ", len(values))
			//case <-over:
			//	fmt.Println("bye")
			//	return
		}
	}
}
