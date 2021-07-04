package main

import (
	"fmt"
)

func deferHand() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 20 {
			panic("number too big")
		}
	}
}

func trace(s string) string {
	fmt.Println("enter ", s)
	return s
}

func leave(s string) {
	fmt.Println("leaving ", s)
}

func a() {
	defer leave(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer leave(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	//deferHand()
	b()
}
