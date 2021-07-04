package main

import (
	"fmt"
	"sync"
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
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	// 定义channel，返回类型int
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	// 调用监听方法
	go doWork(id, w)
	return w
}

// 手动创建chan int通知关闭
func close() {
	const n = 10
	var workers [n]worker

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		workers[i] = createWorker(i+1, &wg)
	}

	for i := 0; i < n; i++ {
		// 向channel发送数据
		workers[i].in <- 'a' + i
		wg.Add(1)
	}

	for i := 0; i < n; i++ {
		workers[i].in <- 'A' + i
		wg.Add(1)
	}
	wg.Wait()

}

func main() {
	close()
}
