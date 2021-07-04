package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			c <- fmt.Sprintf("%s send %d", name, i)
			i++
		}
	}()
	return c
}

// 不确定channel的个数
func fallIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, channel := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(channel)
	}
	return c
}

// 确定channel的个数，使用select
func fallInSelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

// 并行发送消息，通过统一channel接收多个channel的数据，打印
func main() {
	c1 := msgGen("service1")
	c2 := msgGen("service2")
	c3 := msgGen("service3")
	c := fallIn(c1, c2, c3)
	//c := fallInSelect(c1, c2)
	for {
		fmt.Println(<-c)
	}
}
