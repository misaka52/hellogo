package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello")
	fmt.Println(runtime.GOARCH)
}
