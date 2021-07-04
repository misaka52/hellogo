package main

import (
	"fmt"
)

func doWork(id int, w worker) {
	//for {
	//	n, ok := <-c
	// 	通过ok判定channel是否关闭
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("%d receiver %d\n", id, n)
	//}
	// range自动判定channel是否关闭
	for n := range w.in {
		fmt.Printf("%d receiver %c\n", id, n)
		w.done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	// 定义channel，返回类型int
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	// 调用监听方法
	go doWork(id, w)
	return w
}

// 手动创建chan int通知关闭
func close() {
	const n = 10
	var workers [n]worker
	for i := 0; i < n; i++ {
		workers[i] = createWorker(i + 1)
	}

	for i := 0; i < n; i++ {
		// 向channel发送数据
		workers[i].in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i := 0; i < n; i++ {
		workers[i].in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	close()
}
