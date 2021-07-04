package main

import (
	"fmt"
	"time"
)

func main() {
	const conc = 20
	var a [conc]int
	for i := 0; i < conc; i++ {
		go func(i int) {
			for {
				a[i]++
				// 主动释放cpu资源
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
