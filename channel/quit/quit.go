package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chan struct{} 可以直接传递空数据
func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
				c <- fmt.Sprintf("%s send %d", name, i)
				i++
			// 等待结束信号
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleaning done")
				done <- struct{}{}
				return
			}
		}
	}()
	return c
}

func nonBlockingWait(c chan string) (string, bool) {
	select {
	case n := <-c:
		return n, true
	default:
		return "", false
	}
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case n := <-c:
		return n, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	done := make(chan struct{})
	c := msgGen("service1", done)
	for i := 0; i < 5; i++ {
		if v, ok := timeoutWait(c, time.Duration(500)*time.Millisecond); ok {
			fmt.Println(v)
		} else {
			fmt.Println("timeout")
		}
	}
	// 通知channel数据发送完成
	done <- struct{}{}
	// 等待channel成功关闭
	<-done
}
